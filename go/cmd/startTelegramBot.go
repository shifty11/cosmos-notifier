package cmd

import (
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/services/telegram"
	"github.com/spf13/cobra"
)

// startTelegramBotCmd represents the startTelegramBot command
var startTelegramBotCmd = &cobra.Command{
	Use:   "start-telegram-bot",
	Short: "Starts the telegram bot",
	Run: func(cmd *cobra.Command, args []string) {
		webAppUrl := common.GetEnvX("WEBAPP_URL")
		telegramToken := common.GetEnvX("TELEGRAM_BOT_TOKEN")
		useTestApi := common.GetEnvAsBoolX("TELEGRAM_USE_TEST_API")
		apiEndpoint := ""
		if useTestApi {
			apiEndpoint = "https://api.telegram.org/bot%s/test/%s"
		}

		dbManagers := database.NewDefaultDbManagers()

		tgClient := telegram.NewTelegramClient(dbManagers, telegramToken, apiEndpoint, webAppUrl)
		tgClient.Start()
	},
}

func init() {
	rootCmd.AddCommand(startTelegramBotCmd)
}
