package main

import (
	_ "github.com/lib/pq"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/service_crawler/contract_client"
)

func main() {
	var cc = contract_client.NewContractClient("juno1hqr0t3scwkrmuu3554lnqzrccuevd8huxmu533gtm6s7vuzeztzqh6tuwq")
	var config = cc.Config()
	log.Sugar.Infof("config: %v", config)
	var proposals = cc.Proposals()
	for _, proposal := range proposals.Proposals {
		log.Sugar.Infof("%v", proposal.Title)
	}
	log.Sugar.Infof("proposal: %v", cc.Proposal(1))
}
