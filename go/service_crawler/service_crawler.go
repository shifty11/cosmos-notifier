package crawler

import (
	"github.com/robfig/cron/v3"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
)

func Crawl(url string) {
	managers := database.NewDefaultDbManagers()
	updateContracts(url, managers.ContractManager, managers.ProposalManager)
}

func ScheduleCrawl(url string) {
	log.Sugar.Info("Scheduling crawl")
	managers := database.NewDefaultDbManagers()
	c := cron.New()
	_, err := c.AddFunc("@every 1h", func() { updateContracts(url, managers.ContractManager, managers.ProposalManager) })
	if err != nil {
		log.Sugar.Errorf("while executing 'updateContracts' via cron: %v", err)
	}
	c.Start()
}
