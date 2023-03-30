package validator_crawler

import (
	"context"
	"fmt"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/robfig/cron/v3"
	"github.com/shifty11/cosmos-notifier/common"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/types"
	"net/http"
	"time"
)

const urlValidators = "https://rest.cosmos.directory/%v/cosmos/staking/v1beta1/validators"

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

func (c *ValidatorCrawler) validatorNeedsUpdate(validatorEnt *ent.Validator, data *types.Validator) bool {
	return validatorEnt.Moniker != data.Description.Moniker
}

func (c *ValidatorCrawler) isValidatorValid(data *types.Validator) bool {
	return data.OperatorAddress != "" && data.Description.Moniker != ""
}

func (c *ValidatorCrawler) getAccountAddress(validator *types.Validator, chainEnt *ent.Chain) (string, error) {
	_, valAddr, err := bech32.DecodeAndConvert(validator.OperatorAddress)
	if err != nil {
		return "", err
	}
	accAddr, err := cosmossdk.Bech32ifyAddressBytes(chainEnt.Bech32Prefix, valAddr)
	if err != nil {
		return "", err
	}
	return accAddr, nil
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
		existingValidators, err := chainEnt.QueryValidators().All(context.Background())
		if err != nil {
			log.Sugar.Panicf("error getting validators for chain %v: %v", chainEnt.PrettyName, err)
		}

		for _, validator := range validatorsResponse.Validators {
			if !c.isValidatorValid(&validator) {
				continue
			}
			var found = false
			accAddr, err := c.getAccountAddress(&validator, chainEnt)
			if err != nil {
				log.Sugar.Errorf("error getting account address for validator %v: %v", validator.Description.Moniker, err)
				continue
			}
			// TODO: save operator address
			for _, existingValidator := range existingValidators {
				if existingValidator.Address == accAddr {
					found = true
					if c.validatorNeedsUpdate(existingValidator, &validator) {
						log.Sugar.Infof("Updating validator %v", validator.Description.Moniker)
						err := c.validatorManager.UpdateValidator(existingValidator, validator.Description.Moniker)
						if err != nil {
							log.Sugar.Errorf("error updating validator %v: %v", existingValidator.Address, err)
							break
						}
					}
					break
				}
			}
			if !found {
				log.Sugar.Infof("Creating validator %v", validator.Description.Moniker)
				_, err = c.validatorManager.AddValidator(chainEnt, accAddr, validator.Description.Moniker)
				if err != nil {
					log.Sugar.Errorf("error creating validator %v: %v", validator.OperatorAddress, err)
					break
				}
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
