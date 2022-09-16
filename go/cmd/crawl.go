package cmd

import (
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	notifier "github.com/shifty11/dao-dao-notifier/notifier"
	"github.com/shifty11/dao-dao-notifier/service_crawler"
	"time"

	"github.com/spf13/cobra"
)

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Run the crawler script",
	Run: func(cmd *cobra.Command, args []string) {
		nodejsUrl := common.GetEnvX("NODEJS_URL")
		telegramBotToken := common.GetEnvX("TELEGRAM_BOT_TOKEN")
		useTestApi := common.GetEnvAsBoolX("TELEGRAM_USE_TEST_API")
		assetsPath := common.GetEnvX("ASSETS_PATH")
		discordBotToken := common.GetEnvX("DISCORD_BOT_TOKEN")

		apiEndpoint := ""
		if useTestApi {
			apiEndpoint = "https://api.telegram.org/bot%s/test/%s"
		}

		managers := database.NewDefaultDbManagers()
		notifier := notifier.NewNotifier(managers, telegramBotToken, apiEndpoint, discordBotToken)
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
