package cmd

import (
	"github.com/ravener/discord-oauth2"
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/notifier"
	crawler "github.com/shifty11/dao-dao-notifier/service_crawler"
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
		telegramBotToken := common.GetEnvX("TELEGRAM_BOT_TOKEN")
		telegramUseTestApi := common.GetEnvAsBoolX("TELEGRAM_USE_TEST_API")
		telegramApiEndpoint := ""
		if telegramUseTestApi {
			telegramApiEndpoint = "https://api.telegram.org/bot%s/test/%s"
		}
		webAppUrl := common.GetEnvX("WEBAPP_URL")
		discordBotToken := common.GetEnvX("DISCORD_BOT_TOKEN")
		discordClientId := common.GetEnvX("DISCORD_CLIENT_ID")
		discordClientSecret := common.GetEnvX("DISCORD_CLIENT_SECRET")
		nodejsUrl := common.GetEnvX("NODEJS_URL")
		assetsPath := common.GetEnvX("ASSETS_PATH")
		var config = &grpc.Config{
			Port:                 ":50051",
			AccessTokenDuration:  time.Minute * 15,
			RefreshTokenDuration: time.Hour * 24 * 7,
			JwtSecretKey:         jwtSecretKey,
			TelegramToken:        telegramBotToken,
			DiscordOAuth2Config: &oauth2.Config{
				RedirectURL:  webAppUrl,
				ClientID:     discordClientId,
				ClientSecret: discordClientSecret,
				Scopes:       []string{discord.ScopeIdentify},
				Endpoint:     discord.Endpoint,
			},
		}
		dbManagers := database.NewDefaultDbManagers()
		notifier := notifier.NewNotifier(dbManagers, telegramBotToken, telegramApiEndpoint, discordBotToken)
		crawlerClient := crawler.NewCrawler(dbManagers, notifier, nodejsUrl, assetsPath)
		server := grpc.NewGRPCServer(dbManagers, config, crawlerClient)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
