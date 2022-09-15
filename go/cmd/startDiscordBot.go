package cmd

import (
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	discord "github.com/shifty11/dao-dao-notifier/service_discord"
	"github.com/spf13/cobra"
)

// startDiscordBotCmd represents the startDiscordBot command
var startDiscordBotCmd = &cobra.Command{
	Use:   "start-discord-bot",
	Short: "Starts the discord bot",
	Run: func(cmd *cobra.Command, args []string) {
		webAppUrl := common.GetEnvX("WEBAPP_URL")
		discordToken := common.GetEnvX("DISCORD_BOT_TOKEN")

		dbManagers := database.NewDefaultDbManagers()

		discordClient := discord.NewDiscordClient(dbManagers, discordToken, webAppUrl)
		discordClient.Start()
	},
}

func init() {
	rootCmd.AddCommand(startDiscordBotCmd)
}
