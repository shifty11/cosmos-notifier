package database

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/shifty11/cosmos-notifier/common"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
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

func ConnectDuringDevelopment() (*ent.Client, context.Context) {
	if os.Getenv("DEV") != "true" {
		log.Sugar.Panicf("ConnectDuringDevelopment should only be called in development mode")
	}
	return connect()
}

func Close() {
	if dbClient != nil {
		err := dbClient.Close()
		if err != nil {
			log.Sugar.Error(err)
		}
	}
}

type TxContextValue struct {
	Tx           *ent.Tx
	IsCommited   bool
	IsRolledBack bool
}

func getClient(ctx context.Context, client *ent.Client) *ent.Client {
	if ctx.Value(common.ContextKeyTx) != nil {
		return ctx.Value(common.ContextKeyTx).(TxContextValue).Tx.Client()
	}
	return client
}

func startTx(ctx context.Context, client *ent.Client) (context.Context, error) {
	if ctx.Value(common.ContextKeyTx) != nil {
		return nil, errors.New("transaction already started")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	val := TxContextValue{
		Tx:         tx,
		IsCommited: false,
	}
	ctx = context.WithValue(ctx, common.ContextKeyTx, val)
	return ctx, nil
}

func RollbackTxIfUncommitted(ctx context.Context) (context.Context, error) {
	if ctx.Value(common.ContextKeyTx) == nil {
		return ctx, errors.New("transaction not started")
	}
	val := ctx.Value(common.ContextKeyTx).(TxContextValue)
	if val.IsCommited {
		return ctx, nil
	}
	if val.IsRolledBack {
		return ctx, nil
	}
	err := val.Tx.Rollback()
	if err != nil {
		log.Sugar.Error(err)
	}
	val.IsRolledBack = true
	ctx = context.WithValue(ctx, common.ContextKeyTx, val)
	return ctx, nil
}

func CommitTx(ctx context.Context) (context.Context, error) {
	if ctx.Value(common.ContextKeyTx) == nil {
		return ctx, errors.New("transaction not started")
	}
	val := ctx.Value(common.ContextKeyTx).(TxContextValue)
	if val.IsCommited {
		return ctx, nil
	}
	err := val.Tx.Commit()
	if err != nil {
		return ctx, errors.Wrap(err, "committing transaction")
	}
	val.IsCommited = true
	ctx = context.WithValue(ctx, common.ContextKeyTx, val)
	return ctx, nil
}

func withTx(client *ent.Client, ctx context.Context, fn func(tx *ent.Tx) error) error {
	_, err := withTxResult(client, ctx, func(tx *ent.Tx) (*interface{}, error) {
		return nil, fn(tx)
	})
	return err
}

func withTxResult[T any](client *ent.Client, ctx context.Context, fn func(tx *ent.Tx) (*T, error)) (*T, error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if v := recover(); v != nil {
			//goland:noinspection GoUnhandledErrorResult
			tx.Rollback()
			panic(v)
		}
	}()
	result, err := fn(tx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, errors.Wrapf(err, "committing transaction: %v", err)
	}
	return result, nil
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
	StatsManager          *StatsManager
	AddressTrackerManager *AddressTrackerManager
	ValidatorManager      *ValidatorManager
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
	statsManager := NewStatsManager(client, ctx)
	addressTrackerManager := NewAddressTrackerManager(client, ctx)
	validatorManager := NewValidatorManager(client, ctx, addressTrackerManager)
	return &DbManagers{
		ContractManager:       contractManager,
		ProposalManager:       proposalManager,
		UserManager:           userManager,
		DiscordChannelManager: discordChannelManager,
		TelegramChatManager:   telegramChatManager,
		SubscriptionManager:   subscriptionManager,
		ChainManager:          chainManager,
		ChainProposalManager:  chainProposalManager,
		StatsManager:          statsManager,
		AddressTrackerManager: addressTrackerManager,
		ValidatorManager:      validatorManager,
	}
}
