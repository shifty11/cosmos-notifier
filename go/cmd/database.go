/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ariga.io/atlas/sql/sqltool"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent/migrate"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/notifier"
	pb "github.com/shifty11/dao-dao-notifier/services/grpc/protobuf/go/admin_service"
	"github.com/spf13/cobra"
	"strings"
)

// databaseCmd represents the database command
var databaseCmd = &cobra.Command{
	Use:     "database",
	Short:   "Database commands",
	Aliases: []string{"db"},
}

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

// createMigrationsCmd represents the createMigrations command
var createMigrationsCmd = &cobra.Command{
	Use:   "create-migrations",
	Short: "Create migrations based on ent/schema/*.go files",
	Long: `Create migrations based on ent/schema/*.go files

Example with custom db:
go run main.go createMigrations postgres://postgres:postgres@localhost:5432/daodao-notifier-db?sslmode=disable&TimeZone=Europe/Zurich
`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			createMigrations(args[0])
		} else {
			createMigrations(strings.Replace(database.DbCon(), "5432", "5433", 1))
		}
	},
}

func createMigrations(dbCon string) {
	ctx := context.Background()
	// Create a local migration directory able to understand golang-migrate migration files for replay.
	dir, err := sqltool.NewGolangMigrateDir("database/migrations")
	if err != nil {
		log.Sugar.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Write migration diff.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithDropIndex(true),                  // Drop index if exists
		schema.WithDropColumn(true),                 // Drop column if exists
	}

	err = migrate.NamedDiff(ctx, dbCon, "migration", opts...)
	if err != nil {
		log.Sugar.Fatalf("failed generating migration file: %v", err)
	}
}

// TODO: remove after migration
// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import the database",
	Run: func(cmd *cobra.Command, args []string) {
		database.NewDefaultDbManagers().SubscriptionManager.ImportDb()
	},
}

// TODO: remove after migration
// importCmd represents the import command
var sendMsg = &cobra.Command{
	Use:   "send-msg",
	Short: "send msg to zero id chats",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		telegramBotToken := common.GetEnvX("TELEGRAM_BOT_TOKEN")
		telegramUseTestApi := common.GetEnvAsBoolX("TELEGRAM_USE_TEST_API")
		telegramApiEndpoint := ""
		if telegramUseTestApi {
			telegramApiEndpoint = "https://api.telegram.org/bot%s/test/%s"
		}
		discordBotToken := common.GetEnvX("DISCORD_BOT_TOKEN")
		dbManagers := database.NewDefaultDbManagers()
		generalNot := notifier.NewGeneralNotifier(dbManagers, telegramBotToken, telegramApiEndpoint, discordBotToken)
		if len(args) == 1 {
			admins, err := dbManagers.UserManager.GetAdmins()
			if err != nil {
				log.Sugar.Panicf("failed to get admins: %v", err)
			}
			for _, admin := range admins {
				if cmd.Flag("telegram").Value.String() == "true" && admin.Type == user.TypeTelegram {
					if cmd.Flag("test").Value.String() == "true" {
						generalNot.BroadcastMessageToZeroIds(args[0], pb.BroadcastMessageRequest_TELEGRAM_TEST, admin)
					} else {
						generalNot.BroadcastMessageToZeroIds(args[0], pb.BroadcastMessageRequest_TELEGRAM, admin)
					}
				}
				if cmd.Flag("discord").Value.String() == "true" && admin.Type == user.TypeDiscord {
					if cmd.Flag("test").Value.String() == "true" {
						generalNot.BroadcastMessageToZeroIds(args[0], pb.BroadcastMessageRequest_DISCORD_TEST, admin)
					} else {
						generalNot.BroadcastMessageToZeroIds(args[0], pb.BroadcastMessageRequest_DISCORD, admin)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.AddCommand(migrateCmd)
	databaseCmd.AddCommand(createMigrationsCmd)
	databaseCmd.AddCommand(importCmd)
	databaseCmd.AddCommand(sendMsg)

	sendMsg.PersistentFlags().Bool("telegram", false, "Is Telegram")
	sendMsg.PersistentFlags().Bool("discord", false, "Is Discord")
	sendMsg.PersistentFlags().Bool("test", true, "Is test")
}
