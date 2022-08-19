package cmd

import (
	"github.com/shifty11/dao-dao-notifier/database"
	telegram "github.com/shifty11/dao-dao-notifier/service_telegram"

	"github.com/spf13/cobra"
)

// startTelegramBotCmd represents the startTelegramBot command
var startTelegramBotCmd = &cobra.Command{
	Use:   "start-telegram-bot",
	Short: "Starts the telegram bot",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dbManagers := database.NewDefaultDbManagers()
		tgClient := telegram.NewTelegramClient(dbManagers, args[0], "https://api.telegram.org/bot%s/test/%s", "test.mydomain.com:40001")
		tgClient.Start()
	},
}

func init() {
	rootCmd.AddCommand(startTelegramBotCmd)
}
