/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ravener/discord-oauth2"
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/notifier"
	"github.com/shifty11/dao-dao-notifier/services/chain_crawler"
	"github.com/shifty11/dao-dao-notifier/services/contract_crawler"
	discordService "github.com/shifty11/dao-dao-notifier/services/discord"
	"github.com/shifty11/dao-dao-notifier/services/grpc"
	"github.com/shifty11/dao-dao-notifier/services/telegram"
	"golang.org/x/oauth2"
	"time"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Service commands",
}

// startTelegramBotCmd represents the startTelegramBot command
var startTelegramBotCmd = &cobra.Command{
	Use:   "telegram",
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

// startDiscordBotCmd represents the startDiscordBot command
var startDiscordBotCmd = &cobra.Command{
	Use:   "discord",
	Short: "Starts the discord bot",
	Run: func(cmd *cobra.Command, args []string) {
		webAppUrl := common.GetEnvX("WEBAPP_URL")
		discordBotToken := common.GetEnvX("DISCORD_BOT_TOKEN")
		discordClientId := common.GetEnvX("DISCORD_CLIENT_ID")

		dbManagers := database.NewDefaultDbManagers()

		discordClient := discordService.NewDiscordClient(dbManagers, discordBotToken, discordClientId, webAppUrl)
		discordClient.Start()
	},
}

// runChainCrawlerCmd represents the startChainCrawler command
var runChainCrawlerCmd = &cobra.Command{
	Use:   "chain-crawler",
	Short: "Start the chain crawler",
	Run: func(cmd *cobra.Command, args []string) {
		telegramBotToken := common.GetEnvX("TELEGRAM_BOT_TOKEN")
		useTestApi := common.GetEnvAsBoolX("TELEGRAM_USE_TEST_API")
		assetsPath := common.GetEnvX("ASSETS_PATH")
		discordBotToken := common.GetEnvX("DISCORD_BOT_TOKEN")

		apiEndpoint := ""
		if useTestApi {
			apiEndpoint = "https://api.telegram.org/bot%s/test/%s"
		}

		managers := database.NewDefaultDbManagers()
		notifier := notifier.NewChainNotifier(managers, telegramBotToken, apiEndpoint, discordBotToken)
		crawler := chain_crawler.NewChainCrawler(managers, notifier, assetsPath)
		crawler.AddOrUpdateChains()
		crawler.UpdateProposals()

		if cmd.Flag("repeat").Value.String() == "true" {
			crawler.ScheduleCrawl()

			time.Sleep(time.Duration(1<<63 - 1))
		}
	},
}

// runContractCrawlerCmd represents the crawl command
var runContractCrawlerCmd = &cobra.Command{
	Use:   "contract-crawler",
	Short: "Run the contract crawler",
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
		notifier := notifier.NewContractNotifier(managers, telegramBotToken, apiEndpoint, discordBotToken)
		c := contract_crawler.NewContractCrawler(managers, notifier, nodejsUrl, assetsPath)
		c.UpdateContracts()
		if cmd.Flag("repeat").Value.String() == "true" {
			c.ScheduleCrawl()

			time.Sleep(time.Duration(1<<63 - 1))
		}
	},
}

// startGrpcServerCmd represents the grpc command
var startGrpcServerCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start the gRPC server",
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
		notifier := notifier.NewContractNotifier(dbManagers, telegramBotToken, telegramApiEndpoint, discordBotToken)
		crawlerClient := contract_crawler.NewContractCrawler(dbManagers, notifier, nodejsUrl, assetsPath)
		server := grpc.NewGRPCServer(dbManagers, config, crawlerClient)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)

	serviceCmd.AddCommand(startTelegramBotCmd)
	serviceCmd.AddCommand(startDiscordBotCmd)

	serviceCmd.AddCommand(runChainCrawlerCmd)
	runChainCrawlerCmd.PersistentFlags().Bool("repeat", false, "Repeat crawling every hour")

	serviceCmd.AddCommand(runContractCrawlerCmd)
	runContractCrawlerCmd.PersistentFlags().Bool("repeat", false, "Repeat crawling every hour")
}
