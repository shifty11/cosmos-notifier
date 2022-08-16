package cmd

import (
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/service_crawler"

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
	service_crawler.UpdateContracts(managers.ContractManager, managers.ProposalManager)
}
