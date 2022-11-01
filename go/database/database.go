package database

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	dbName     = "cosmos-notifier-db"
	dbPort     = "5432"
	dbSSLMode  = "disable"
	dbTimezone = "Europe/Zurich"
)

func DbCon() string {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=%v&TimeZone=%v", dbType, dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode, dbTimezone)
	}
	return dsn
}

func connect() (*ent.Client, context.Context) {
	if dbClient == nil {
		newClient, err := ent.Open("postgres", DbCon())
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

func MigrateDb() error {
	m, err := migrate.New("file://database/migrations/", DbCon())
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Sugar.Info("no migration needed")
			return nil
		}
	}
	return err
}

type DbManagers struct {
	ContractManager       IContractManager
	ProposalManager       *ContractProposalManager
	UserManager           *UserManager
	DiscordChannelManager IDiscordChannelManager
	TelegramChatManager   ITelegramChatManager
	SubscriptionManager   *SubscriptionManager
	ChainManager          *ChainManager
	ChainProposalManager  *ChainProposalManager
}

func NewDefaultDbManagers() *DbManagers {
	client, ctx := connect()
	return NewCustomDbManagers(client, ctx)
}

func NewCustomDbManagers(client *ent.Client, ctx context.Context) *DbManagers {
	chainManager := NewChainManager(client, ctx)
	contractManager := NewContractManager(client, ctx)
	proposalManager := NewContractProposalManager(client, ctx)
	userManager := NewUserManager(client, ctx)
	discordChannelManager := NewDiscordChannelManager(client, ctx, chainManager, contractManager, userManager)
	telegramChatManager := NewTelegramChatManager(client, ctx, chainManager, contractManager, userManager)
	subscriptionManager := NewSubscriptionManager(client, ctx, userManager, chainManager, contractManager, telegramChatManager, discordChannelManager)
	chainProposalManager := NewChainProposalManager(client, ctx)
	return &DbManagers{
		ContractManager:       contractManager,
		ProposalManager:       proposalManager,
		UserManager:           userManager,
		DiscordChannelManager: discordChannelManager,
		TelegramChatManager:   telegramChatManager,
		SubscriptionManager:   subscriptionManager,
		ChainManager:          chainManager,
		ChainProposalManager:  chainProposalManager,
	}
}
