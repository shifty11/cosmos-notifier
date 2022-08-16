package database

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
	"os"
)

var dbClient *ent.Client

var (
	dbType     = "postgres"
	dbHost     = "localhost"
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "daodao-notifier-db"
	dbPort     = "5432"
	dbSSLMode  = "disable"
	dbTimezone = "Europe/Zurich"
)

func DbCon() string {
	return fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=%v&TimeZone=%v", dbType, dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode, dbTimezone)
}

func connect() (*ent.Client, context.Context) {
	if dbClient == nil {
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			dsn = DbCon()
		}
		newClient, err := ent.Open("postgres", dsn)
		if err != nil {
			log.Sugar.Panic("failed to connect to server ", err)
		}
		dbClient = newClient
	}
	return dbClient, context.Background()
}

func Close() {
	if dbClient != nil {
		err := dbClient.Close()
		if err != nil {
			log.Sugar.Error(err)
		}
	}
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			//goland:noinspection GoUnhandledErrorResult
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

type DbManagers struct {
	ContractManager *ContractManager
	ProposalManager *ProposalManager
	UserManager     *UserManager
}

func NewDefaultDbManagers() *DbManagers {
	client, ctx := connect()
	return NewCustomDbManagers(client, ctx)
}

func NewCustomDbManagers(client *ent.Client, ctx context.Context) *DbManagers {
	contractManager := NewContractManager(client, ctx)
	proposalManager := NewProposalManager(client, ctx)
	userManager := NewUserManager(client, ctx)
	return &DbManagers{
		ContractManager: contractManager,
		ProposalManager: proposalManager,
		UserManager:     userManager,
	}
}
