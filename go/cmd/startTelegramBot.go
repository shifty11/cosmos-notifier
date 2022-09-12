package cmd

import (
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
	telegram "github.com/shifty11/dao-dao-notifier/service_telegram"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// startTelegramBotCmd represents the startTelegramBot command
var startTelegramBotCmd = &cobra.Command{
	Use:   "start-telegram-bot",
	Short: "Starts the telegram bot",
	Run: func(cmd *cobra.Command, args []string) {
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

		webAppUrl := os.Getenv("TELEGRAM_WEBAPP_URL")
		if webAppUrl == "" {
			log.Sugar.Panic("TELEGRAM_WEBAPP_URL must be set")
		}

		dbManagers := database.NewDefaultDbManagers()

		tgClient := telegram.NewTelegramClient(dbManagers, telegramToken, apiEndpoint, webAppUrl)
		tgClient.Start()
	},
}

func init() {
	rootCmd.AddCommand(startTelegramBotCmd)
}
