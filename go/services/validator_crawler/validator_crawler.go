package validator_crawler

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/shifty11/cosmos-notifier/common"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/types"
	"net/http"
	"time"
)

const urlValidators = "https://rest.cosmos.directory/%v/cosmos/staking/v1beta1/validators?pagination.limit=1000"
const urlValidatorSet = "https://rest.cosmos.directory/%v/cosmos/base/tendermint/v1beta1/validatorsets/latest"

type ValidatorCrawler struct {
	httpClient            *http.Client
	chainManager          *database.ChainManager
	addressTrackerManager *database.AddressTrackerManager
	validatorManager      *database.ValidatorManager
	errorReporter         common.ErrorReporter
}

func NewValidatorCrawler(dbManagers *database.DbManagers) *ValidatorCrawler {
	var client = &http.Client{Timeout: 10 * time.Second}
	return &ValidatorCrawler{
		httpClient:            client,
		chainManager:          dbManagers.ChainManager,
		addressTrackerManager: dbManagers.AddressTrackerManager,
		validatorManager:      dbManagers.ValidatorManager,
		errorReporter:         common.NewErrorReporter(common.DefaultMaxErrorCntUntilReport),
	}
}

func validatorNeedsUpdate(validatorEnt *ent.Validator, data *types.Validator, isValidatorInActiveSet bool) bool {
	return validatorEnt.Moniker != data.Description.Moniker ||
		(validatorEnt.FirstInactiveTime != nil && isValidatorInActiveSet) ||
		(validatorEnt.FirstInactiveTime == nil && !isValidatorInActiveSet)
}

func isValidatorValid(data *types.Validator) bool {
	return data.OperatorAddress != ""
}

func getExistingValidator(validators []*ent.Validator, validator *types.Validator) *ent.Validator {
	for _, val := range validators {
		if val.OperatorAddress == validator.OperatorAddress {
			return val
		}
	}
	return nil
}

// AddOrUpdateValidators compares a pubkey of a validator with the pubkeys of the validators in the active set.
// We can not use the address from the active set because it is a `valcons` address which we would have to convert first.
func isValidatorInActiveSet(pubKey string, activeValidatorSet []types.ValidatorSetValidator) bool {
	for _, validator := range activeValidatorSet {
		if pubKey == validator.PubKey.Key {
			return true
		}
	}
	return false
}

func (c *ValidatorCrawler) AddOrUpdateValidators() {
	log.Sugar.Info("Getting all validators")
	for _, chainEnt := range c.chainManager.Enabled() {
		log.Sugar.Infof("Getting validators for chain %v", chainEnt.PrettyName)
		url := fmt.Sprintf(urlValidators, chainEnt.Path)
		var validatorsResponse types.ValidatorsResponse
		_, err := common.GetJson(c.httpClient, url, &validatorsResponse)
		if err != nil {
			log.Sugar.Errorf("error calling %v: %v", url, err)
			continue
		}
		if validatorsResponse.Pagination.Total != "0" {
			log.Sugar.Errorf("pagination is not implemented yet")
		}

		var validatorSetResponse types.ValidatorSetResponse
		url = fmt.Sprintf(urlValidatorSet, chainEnt.Path)
		_, err = common.GetJson(c.httpClient, url, &validatorSetResponse)
		if err != nil {
			log.Sugar.Errorf("error calling %v: %v", url, err)
			continue
		}

		existingValidators, err := chainEnt.QueryValidators().All(context.Background())
		if err != nil {
			log.Sugar.Panicf("error getting validators for chain %v: %v", chainEnt.PrettyName, err)
		}

		for _, validator := range validatorsResponse.Validators {
			if isValidatorValid(&validator) {
				var isInActiveSet = isValidatorInActiveSet(validator.ConsensusPubkey.Key, validatorSetResponse.Validators)
				var existingValidator = getExistingValidator(existingValidators, &validator)
				if existingValidator != nil {
					if validatorNeedsUpdate(existingValidator, &validator, isInActiveSet) {
						log.Sugar.Infof("Updating validator %v %v", validator.OperatorAddress, validator.Description.Moniker)
						err := c.validatorManager.UpdateValidator(existingValidator, validator.Description.Moniker, isInActiveSet)
						if err != nil {
							log.Sugar.Errorf("error updating validator %v: %v", existingValidator.Address, err)
							continue
						}
					}
				} else {
					log.Sugar.Infof("Creating validator %v %v", validator.OperatorAddress, validator.Description.Moniker)
					_, err = c.validatorManager.AddValidator(chainEnt, validator.OperatorAddress, validator.Description.Moniker, isInActiveSet)
					if err != nil {
						log.Sugar.Errorf("error creating validator %v: %v", validator.OperatorAddress, err)
						continue
					}
				}
			} else {
				log.Sugar.Debugf("Validator %v %v is invalid", validator.OperatorAddress, validator.Description.Moniker)
			}
		}
	}
}

func (c *ValidatorCrawler) ScheduleCrawl() {
	log.Sugar.Info("Scheduling validator crawl")
	cr := cron.New()
	_, err := cr.AddFunc("0 10 * * *", func() { c.AddOrUpdateValidators() }) // every day at 10:00
	if err != nil {
		log.Sugar.Errorf("while executing 'AddOrUpdateValidators' via cron: %v", err)
	}
	cr.Start()
}
