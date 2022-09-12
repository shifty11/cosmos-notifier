package crawler

import (
	"bufio"
	"github.com/robfig/cron/v3"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/notifier"
	"os"
)

type Crawler struct {
	contractManager *database.ContractManager
	proposalManager *database.ProposalManager
	notifier        *notifier.Notifier
	apiUrl          string
	assetsPath      string
}

func NewCrawler(managers *database.DbManagers, notifier *notifier.Notifier, apiUrl string, assetsPath string) *Crawler {
	return &Crawler{
		contractManager: managers.ContractManager,
		proposalManager: managers.ProposalManager,
		notifier:        notifier,
		apiUrl:          apiUrl,
		assetsPath:      assetsPath,
	}
}

func (c *Crawler) contracts() []string {
	file, err := os.Open("contracts.txt")
	if err != nil {
		file, err = os.Open("../contracts.txt")
		if err != nil {
			log.Sugar.Error(err)
		}
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	sc := bufio.NewScanner(file)
	contracts := make([]string, 0)

	// Read through 'contracts' until an EOF is encountered.
	for sc.Scan() {
		contracts = append(contracts, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Sugar.Error(err)
	}
	return contracts
}

func (c *Crawler) UpdateContracts() {
	log.Sugar.Info("updating contracts")
	for _, contractAddr := range c.contracts() {
		client := NewContractClient(c.apiUrl, contractAddr)
		config, err := client.config()
		if err != nil {
			log.Sugar.Errorf("while getting config for contract %v: %v", contractAddr, err)
			continue
		}
		proposals, err := client.proposals()
		if err != nil {
			log.Sugar.Errorf("while getting proposals for contract %v: %v", contractAddr, err)
			continue
		}

		contract, contractStatus := c.contractManager.CreateOrUpdate(contractAddr, config)
		for _, proposal := range proposals.Proposals {
			dbProp, proposalStatus := c.proposalManager.CreateOrUpdate(contract, &proposal)
			if proposalStatus == database.ProposalStatusChanged {
				log.Sugar.Infof("Proposal %v changed status to %v", dbProp.ID, dbProp.Status)
			} else if proposalStatus == database.ProposalCreated {
				c.notifier.Notify(dbProp)
			}
		}

		im := NewImageManager(contractAddr, c.assetsPath, "images/contracts/", 100, 100)
		if contractStatus == database.ContractCreated || contractStatus == database.ContractImageChanged || !im.DoesExist() {
			if contract.ImageURL != "" {
				err := im.downloadAndCreateThumbnail(contract.ImageURL)
				if err != nil {
					log.Sugar.Errorf("while downloading image for contract %v: %v", contractAddr, err)
				} else {
					c.contractManager.SaveThumbnailUrl(contract, im.ThumbnailUrl)
				}
			} else if contractStatus == database.ContractImageChanged {
				c.contractManager.SaveThumbnailUrl(contract, "")
				e := os.Remove(im.ThumbnailPath)
				if e != nil {
					log.Sugar.Errorf("while removing image for contract %v: %v", contractAddr, e)
				}
			}
		}

		log.Sugar.Infof("processed contract %v (%v): %v", config.Name, contractAddr, contractStatus)
	}
}

func (c *Crawler) ScheduleCrawl() {
	log.Sugar.Info("Scheduling crawl")
	cr := cron.New()
	_, err := cr.AddFunc("@every 1h", func() { c.UpdateContracts() })
	if err != nil {
		log.Sugar.Errorf("while executing 'updateContracts' via cron: %v", err)
	}
	cr.Start()
}
