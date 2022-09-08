package cmd

import (
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Run: func(cmd *cobra.Command, args []string) {
		err := database.MigrateDb()
		if err != nil {
			log.Sugar.Panicf("failed to migrate database: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
