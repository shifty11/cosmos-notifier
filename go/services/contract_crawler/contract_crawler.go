package contract_crawler

import (
	"context"
	"github.com/hasura/go-graphql-client"
	"github.com/robfig/cron/v3"
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/notifier"
	"github.com/shifty11/dao-dao-notifier/types"
	"os"
)

type ContractCrawler struct {
	contractManager database.IContractManager
	chainManager    *database.ChainManager
	proposalManager *database.ContractProposalManager
	notifier        *notifier.ContractNotifier
	apiUrl          string
	assetsPath      string
}

func NewContractCrawler(managers *database.DbManagers, notifier *notifier.ContractNotifier, apiUrl string, assetsPath string) *ContractCrawler {
	return &ContractCrawler{
		contractManager: managers.ContractManager,
		chainManager:    managers.ChainManager,
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
	for _, contract := range contracts {
		client := NewContractClient(c.apiUrl, contract.Address, contract.RPCEndpoint)
		config, err := client.config(types.ContractVersion(contract.ConfigVersion))
		if err != nil {
			log.Sugar.Debugf("error while getting config for contract %v (%v): %v", contract.Name, contract.Address, err)
			continue
		}
		proposals, err := client.proposals("")
		if err != nil {
			log.Sugar.Debugf("error while getting proposals for contract %v (%v): %v", contract.Name, contract.Address, err)
			continue
		}

		oldImageUrl := contract.ImageURL
		updatedContract := c.contractManager.Update(contract, config)
		for _, proposal := range proposals.Proposals {
			if proposal.Status == types.StatusOpen {
				dbProp, proposalStatus := c.proposalManager.CreateOrUpdate(updatedContract, &proposal)
				if proposalStatus == database.ProposalStatusChangedFromOpen {
					log.Sugar.Infof("Proposal %v changed status to %v", dbProp.ID, dbProp.Status)
				} else if proposalStatus == database.ProposalCreated {
					c.notifier.Notify(updatedContract, dbProp)
				}
			}
		}

		im := common.NewImageManager(
			updatedContract.Address,
			updatedContract.Name,
			updatedContract.ImageURL,
			c.assetsPath,
			"shared/contracts/",
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

func (c *ContractCrawler) AddContracts() {
	log.Sugar.Info("Add contracts")

	var query struct {
		Daos struct {
			Nodes []struct {
				Address     string `graphql:"id"`
				Name        string
				Description string
				ImageUrl    string
			}
		}
	}

	client := graphql.NewClient("https://index.daodao.zone/daos", nil)

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		log.Sugar.Errorf("while querying daos: %v", err)
	}

	junoChain, err := c.chainManager.GetByName("juno")
	if err != nil {
		log.Sugar.Panicf("chain Juno not found")
	}

	contracts := c.contractManager.All()
	for _, dao := range query.Daos.Nodes {
		found := false
		for _, contract := range contracts {
			if contract.Address == dao.Address {
				found = true
				break
			}
		}
		if !found {
			_, err := c.AddContract(junoChain, dao.Address, "")
			if err != nil {
				log.Sugar.Errorf("while adding contract %v: %v", dao.Address, err)
			}
		}
	}
}

func (c *ContractCrawler) AddContract(chain *ent.Chain, contractAddr string, proposalQuery string) (*ent.Contract, error) {
	rpcName := chain.Path
	if rpcName == "" {
		rpcName = chain.Name
	}
	client := NewContractClient(c.apiUrl, contractAddr, "https://rpc.cosmos.directory/"+rpcName)
	config, err := client.config(types.ContractVersionUnknown)
	if err != nil {
		return nil, err
	}
	proposals, err := client.proposals(proposalQuery)
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
		"shared/contracts/",
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
		log.Sugar.Errorf("while executing 'UpdateContracts' via cron: %v", err)
	}
	_, err = cr.AddFunc("0 10 * * *", func() { c.AddContracts() }) // every day at 10:00
	if err != nil {
		log.Sugar.Errorf("while executing 'AddContracts' via cron: %v", err)
	}
	cr.Start()
}
