package cmd

import (
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/service_crawler/contract_client"

	"github.com/spf13/cobra"
)

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Run the crawler script",
	Run: func(cmd *cobra.Command, args []string) {
		crawl()
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)
}

func crawl() {
	managers := database.NewDefaultDbManagers()
	contract_client.UpdateContracts(managers.ContractManager, managers.ProposalManager)
}
