package cmd

import (
	"github.com/ravener/discord-oauth2"
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	grpc "github.com/shifty11/dao-dao-notifier/service_grpc"
	"golang.org/x/oauth2"
	"time"

	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Run gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		jwtSecretKey := common.GetEnvX("JWT_SECRET_KEY")
		telegramToken := common.GetEnvX("TELEGRAM_TOKEN")
		webAppUrl := common.GetEnvX("WEBAPP_URL")
		discordClientId := common.GetEnvX("DISCORD_CLIENT_ID")
		discordClientSecret := common.GetEnvX("DISCORD_CLIENT_SECRET")
		var config = &grpc.Config{
			Port:                 ":50051",
			AccessTokenDuration:  time.Minute * 15,
			RefreshTokenDuration: time.Hour * 24,
			JwtSecretKey:         jwtSecretKey,
			TelegramToken:        telegramToken,
			DiscordOAuth2Config: &oauth2.Config{
				RedirectURL:  webAppUrl,
				ClientID:     discordClientId,
				ClientSecret: discordClientSecret,
				Scopes:       []string{discord.ScopeIdentify},
				Endpoint:     discord.Endpoint,
			},
		}
		dbManagers := database.NewDefaultDbManagers()
		server := grpc.NewGRPCServer(dbManagers, config)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
