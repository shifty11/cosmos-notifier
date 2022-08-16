package cmd

import (
	"github.com/shifty11/dao-dao-notifier/service_crawler"

	"github.com/spf13/cobra"
)

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Run the crawler script",
	Run: func(cmd *cobra.Command, args []string) {
		service_crawler.Crawl()
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)
}
