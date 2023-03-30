package database

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/enttest"
	"testing"
)

func testClient(t *testing.T) *ent.Client {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	filename := fmt.Sprintf("file:ent%v?mode=memory&cache=shared&_fk=1", t.Name())
	client := enttest.Open(t, "sqlite3", filename, opts...)
	return client
}

func closeTestClient(client *ent.Client) {
	//goland:noinspection GoUnhandledErrorResult
	defer client.Close()
}

func newTestDbManagers(t *testing.T) *DbManagers {
	chainManager := newTestChainManager(t)
	contractManager := newTestContractManager(t)
	userManager := newTestUserManager(t)
	discordChannelManager := newTestDiscordChannelManager(t)
	telegramChatManager := newTestTelegramChatManager(t)
	subscriptionManager := newTestSubscriptionManager(t)
	chainProposalManager := newTestChainProposalManager(t)
	addressTrackerManager := newTestAddressTrackerManager(t)
	validatorManager := newTestValidatorManager(t)
	return &DbManagers{
		ContractManager:       contractManager,
		ProposalManager:       nil,
		UserManager:           userManager,
		DiscordChannelManager: discordChannelManager,
		TelegramChatManager:   telegramChatManager,
		SubscriptionManager:   subscriptionManager,
		ChainManager:          chainManager,
		ChainProposalManager:  chainProposalManager,
		StatsManager:          nil,
		AddressTrackerManager: addressTrackerManager,
		ValidatorManager:      validatorManager,
	}
}
