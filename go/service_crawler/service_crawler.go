package service_crawler

import "github.com/shifty11/dao-dao-notifier/database"

func Crawl() {
	managers := database.NewDefaultDbManagers()
	updateContracts(managers.ContractManager, managers.ProposalManager)
}
