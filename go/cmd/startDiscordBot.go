package cmd

import (
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/services/discord"
	"github.com/spf13/cobra"
)

// startDiscordBotCmd represents the startDiscordBot command
var startDiscordBotCmd = &cobra.Command{
	Use:   "start-discord-bot",
	Short: "Starts the discord bot",
	Run: func(cmd *cobra.Command, args []string) {
		webAppUrl := common.GetEnvX("WEBAPP_URL")
		discordBotToken := common.GetEnvX("DISCORD_BOT_TOKEN")
		discordClientId := common.GetEnvX("DISCORD_CLIENT_ID")

		dbManagers := database.NewDefaultDbManagers()

		discordClient := discord.NewDiscordClient(dbManagers, discordBotToken, discordClientId, webAppUrl)
		discordClient.Start()
	},
}

func init() {
	rootCmd.AddCommand(startDiscordBotCmd)
}
