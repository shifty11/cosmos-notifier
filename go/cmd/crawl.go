package cmd

import (
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
	notifier "github.com/shifty11/dao-dao-notifier/notifier"
	"github.com/shifty11/dao-dao-notifier/service_crawler"
	"os"
	"strconv"
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
		telegramToken := os.Getenv("TELEGRAM_TOKEN")
		if telegramToken == "" {
			log.Sugar.Panic("TELEGRAM_TOKEN must be set")
		}
		useTestApi, err := strconv.ParseBool(os.Getenv("TELEGRAM_USE_TEST_API"))
		if err != nil {
			log.Sugar.Panic("TELEGRAM_USE_TEST_API must be set")
		}
		apiEndpoint := ""
		if useTestApi {
			apiEndpoint = "https://api.telegram.org/bot%s/test/%s"
		}

		assetsPath := os.Getenv("ASSETS_PATH")
		if err != nil {
			log.Sugar.Panic("ASSETS_PATH must be set")
		}

		managers := database.NewDefaultDbManagers()
		notifier := notifier.NewNotifier(managers, telegramToken, apiEndpoint)
		c := crawler.NewCrawler(managers, notifier, nodejsUrl, assetsPath)
		c.UpdateContracts()
		if cmd.Flag("repeat").Value.String() == "true" {
			c.ScheduleCrawl()

			time.Sleep(time.Duration(1<<63 - 1))
		}
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)
	crawlCmd.PersistentFlags().Bool("repeat", false, "Repeat crawling every hour")
}
