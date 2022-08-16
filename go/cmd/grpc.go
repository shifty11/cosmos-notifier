/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
	grpc "github.com/shifty11/dao-dao-notifier/service_grpc"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Run gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
		if jwtSecretKey == "" {
			log.Sugar.Panic("JWT_SECRET_KEY must be set")
		}
		telegramToken := os.Getenv("TELEGRAM_TOKEN")
		if telegramToken == "" {
			log.Sugar.Panic("TELEGRAM_TOKEN must be set")
		}
		var config = &grpc.Config{
			Port:                 ":50051",
			AccessTokenDuration:  time.Minute * 15,
			RefreshTokenDuration: time.Hour * 24,
			JwtSecretKey:         jwtSecretKey,
			TelegramToken:        telegramToken,
		}
		dbManagers := database.NewDefaultDbManagers()
		server := grpc.NewGRPCServer(dbManagers, config)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
