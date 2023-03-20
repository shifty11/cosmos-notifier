/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ariga.io/atlas/sql/sqltool"
	"bufio"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent/migrate"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
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

func getMigrationChecksum(migrationPath string) string {
	file, err := os.Open(filepath.Join(migrationPath, "atlas.sum"))
	if err != nil {
		return ""
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const maxCapacity int = 100 // required line length
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func createMigrations(dbCon string) {
	ctx := context.Background()
	// Create a local migration directory able to understand golang-migrate migration files for replay.
	migrationPath := "database/migrations"
	dir, err := sqltool.NewGolangMigrateDir(migrationPath)
	if err != nil {
		log.Sugar.Fatalf("failed creating atlas migration directory: %v", err)
	}

	checksum := getMigrationChecksum(migrationPath)

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
	if checksum == getMigrationChecksum(migrationPath) {
		log.Sugar.Info("no changes detected")
	} else {
		log.Sugar.Info("migrations created successfully")
	}
}

func init() {
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.AddCommand(migrateCmd)
	databaseCmd.AddCommand(createMigrationsCmd)
}
