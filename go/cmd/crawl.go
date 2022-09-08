package cmd

import (
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/service_crawler"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Run the crawler script",
	Run: func(cmd *cobra.Command, args []string) {
		nodejsUrl := os.Getenv("NODEJS_URL")
		if nodejsUrl == "" {
			log.Sugar.Panic("NODEJS_URL must be set")
		}
		crawler.Crawl(nodejsUrl)
		if cmd.Flag("repeat").Value.String() == "true" {
			crawler.ScheduleCrawl(nodejsUrl)

			time.Sleep(time.Duration(1<<63 - 1))
		}
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)
	crawlCmd.PersistentFlags().Bool("repeat", false, "Repeat crawling every hour")
}
