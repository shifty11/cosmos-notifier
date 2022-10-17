/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/services/chain_crawler"

	"github.com/spf13/cobra"
)

// startChainCrawlerCmd represents the startChainCrawler command
var startChainCrawlerCmd = &cobra.Command{
	Use:   "start-chain-crawler",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		assetsPath := common.GetEnvX("ASSETS_PATH")
		chain_crawler.NewChainCrawler(database.NewDefaultDbManagers(), assetsPath).AddOrUpdateChains()
	},
}

func init() {
	rootCmd.AddCommand(startChainCrawlerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startChainCrawlerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startChainCrawlerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
