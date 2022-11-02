package contract_crawler

import (
	"github.com/robfig/cron/v3"
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/notifier"
	"os"
)

type ContractCrawler struct {
	contractManager database.IContractManager
	proposalManager *database.ContractProposalManager
	notifier        *notifier.ContractNotifier
	apiUrl          string
	assetsPath      string
}

func NewContractCrawler(managers *database.DbManagers, notifier *notifier.ContractNotifier, apiUrl string, assetsPath string) *ContractCrawler {
	return &ContractCrawler{
		contractManager: managers.ContractManager,
		proposalManager: managers.ProposalManager,
		notifier:        notifier,
		apiUrl:          apiUrl,
		assetsPath:      assetsPath,
	}
}

func (c *ContractCrawler) UpdateContracts() {
	log.Sugar.Info("Updating contracts")

	contracts := c.contractManager.All()
	var cntSuccess, cntFails = 0, len(contracts)
	for _, oldContract := range contracts {
		client := NewContractClient(c.apiUrl, oldContract.Address)
		config, err := client.config()
		if err != nil {
			log.Sugar.Debugf("error while getting config for contract %v (%v): %v", oldContract.Name, oldContract.Address, err)
			continue
		}
		proposals, err := client.proposals()
		if err != nil {
			log.Sugar.Debugf("error while getting proposals for contract %v (%v): %v", oldContract.Name, oldContract.Address, err)
			continue
		}

		oldImageUrl := oldContract.ImageURL
		updatedContract := c.contractManager.Update(oldContract, config)
		for _, proposal := range proposals.Proposals {
			dbProp, proposalStatus := c.proposalManager.CreateOrUpdate(updatedContract, &proposal)
			if proposalStatus == database.ProposalStatusChangedFromOpen {
				log.Sugar.Infof("Proposal %v changed status to %v", dbProp.ID, dbProp.Status)
			} else if proposalStatus == database.ProposalCreated {
				c.notifier.Notify(updatedContract, dbProp)
			}
		}

		im := common.NewImageManager(
			updatedContract.Address,
			updatedContract.Name,
			updatedContract.ImageURL,
			c.assetsPath,
			"images/contracts/",
			100,
			100,
		)
		if oldImageUrl != updatedContract.ImageURL || !im.DoesExist() {
			if updatedContract.ImageURL != "" {
				err := im.DownloadAndCreateThumbnail()
				if err != nil {
					log.Sugar.Infof("while downloading image for contract %v (%v): %v", updatedContract.Name, updatedContract.Address, err)
				} else {
					c.contractManager.SaveThumbnailUrl(updatedContract, im.ThumbnailUrl)
				}
			} else if oldImageUrl != updatedContract.ImageURL {
				c.contractManager.SaveThumbnailUrl(updatedContract, "")
				e := os.RemoveAll(im.ThumbnailPath)
				if e != nil {
					log.Sugar.Errorf("while removing image for contract %v (%v): %v", updatedContract.Name, updatedContract.Address, e)
				}
			}
		}
		cntSuccess++
		cntFails--
		log.Sugar.Infof("processed contract %v (%v)", config.Name, updatedContract.Address)
	}

	log.Sugar.Infof("processed %v contracts, success: %v failed: %v", len(contracts), cntSuccess, cntFails)
}

func (c *ContractCrawler) AddContract(contractAddr string) (*ent.Contract, error) {
	client := NewContractClient(c.apiUrl, contractAddr)
	config, err := client.config()
	if err != nil {
		return nil, err
	}
	proposals, err := client.proposals()
	if err != nil {
		return nil, err
	}

	contract, err := c.contractManager.Create(config)
	if err != nil {
		return nil, err
	}

	for _, proposal := range proposals.Proposals {
		c.proposalManager.CreateOrUpdate(contract, &proposal)
	}

	im := common.NewImageManager(
		contractAddr,
		contract.Name,
		contract.ImageURL,
		c.assetsPath,
		"images/contracts/",
		100,
		100,
	)
	err = im.DownloadAndCreateThumbnail()
	if err != nil {
		log.Sugar.Infof("while downloading image for contract %v: %v", contractAddr, err)
	} else {
		c.contractManager.SaveThumbnailUrl(contract, im.ThumbnailUrl)
	}

	return contract, nil
}

func (c *ContractCrawler) ByAddress(contractAddr string) (*ent.Contract, error) {
	return c.contractManager.ByAddress(contractAddr)
}

func (c *ContractCrawler) ScheduleCrawl() {
	log.Sugar.Info("Scheduling crawl")
	cr := cron.New()
	_, err := cr.AddFunc("@every 1h", func() { c.UpdateContracts() })
	if err != nil {
		log.Sugar.Errorf("while executing 'updateContracts' via cron: %v", err)
	}
	cr.Start()
}
