package database

import (
	"context"
	"fmt"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"golang.org/x/exp/slices"
	"testing"
)

func newTestValidatorManager(t *testing.T) *ValidatorManager {
	manager := NewValidatorManager(testClient(t), context.Background(), newTestAddressTrackerManager(t))
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func createValidBech32Address(bech32Prefix string, address string) string {
	_, valAddr, err := bech32.DecodeAndConvert(address)
	if err != nil {
		panic(err)
	}
	accAddr, err := cosmossdk.Bech32ifyAddressBytes(bech32Prefix, valAddr)
	if err != nil {
		panic(err)
	}
	return accAddr
}

func addValidators(m *ValidatorManager, chains []*ent.Chain) []*ent.Validator {
	var validators []*ent.Validator
	for _, chainEnt := range chains {
		val, err := m.AddValidator(
			chainEnt,
			createValidBech32Address(chainEnt.Bech32Prefix+"valoper", "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"),
			fmt.Sprintf("validator %s", chainEnt.Name),
			true,
		)
		if err != nil {
			panic(err)
		}
		validators = append(validators, val)
	}
	return validators
}

func TestValidatorManager_AddValidator(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	m := newTestValidatorManager(t)

	_, err := m.AddValidator(chains[0], "", "validator 1", true)
	if err == nil {
		t.Error("expected error")
	}

	val, err := m.AddValidator(chains[0], "cosmosvaloper156gqf9837u7d4c4678yt3rl4ls9c5vuursrrzf", "validator 1", true)
	if err != nil {
		t.Error(err)
	}
	if val.OperatorAddress != "cosmosvaloper156gqf9837u7d4c4678yt3rl4ls9c5vuursrrzf" {
		t.Error("expected address to be cosmosvaloper156gqf9837u7d4c4678yt3rl4ls9c5vuursrrzf")
	}
	if val.Address != "cosmos156gqf9837u7d4c4678yt3rl4ls9c5vuuxyhkw6" {
		t.Error("expected address to be cosmos156gqf9837u7d4c4678yt3rl4ls9c5vuuxyhkw6")
	}
	if val.Moniker != "validator 1" {
		t.Error("expected moniker to be validator 1")
	}
	if val.QueryChain().AllX(m.ctx)[0].ID != chains[0].ID {
		t.Error("expected chain id to be 1")
	}
	if val.FirstInactiveTime != nil {
		t.Error("expected first inactive time to be nil")
	}

	val, err = m.AddValidator(chains[0], "cosmosvaloper1vvwtk805lxehwle9l4yudmq6mn0g32px9xtkhc", "other val", false)
	if err != nil {
		t.Error(err)
	}
	if val.FirstInactiveTime == nil {
		t.Error("expected first inactive time to be set")
	}

	// check constraints and validation
	_, err = m.AddValidator(chains[0], "cosmosvaloper196ax4vc0lwpxndu9dyhvca7jhxp70rmcvrj90c", "new val", true)
	if err != nil {
		t.Error("did not expect error")
	}
	_, err = m.AddValidator(chains[0], "cosmosvaloper196ax4vc0lwpxndu9dyhvca7jhxp70rmcvrj90c", "new val", true)
	if err == nil {
		t.Error("expected error")
	}
	_, err = m.AddValidator(chains[0], "cosmosvaloper196ax4vc0lwpxndu9dyhvca7jhxp70rmcvrj90c", "other moniker", true)
	if err == nil {
		t.Error("expected error")
	}
	_, err = m.AddValidator(chains[0], "cosmosvaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u2lcnj0", "new val", true)
	if err != nil {
		t.Error("did not expect error")
	}
	_, err = m.AddValidator(chains[0], "invalid address", "new val", true)
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidatorManager_UpdateValidator(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	m := newTestValidatorManager(t)
	vals := addValidators(m, chains[:1])

	err := m.UpdateValidator(vals[0], "validator 2", false)
	if err != nil {
		t.Error(err)
	}
	if m.client.Validator.GetX(m.ctx, vals[0].ID).FirstInactiveTime == nil {
		t.Error("expected first inactive time to be set")
	}
}

func TestValidatorManager_DeleteValidator(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	m := newTestValidatorManager(t)
	vals := addValidators(m, chains[:1])

	cnt := m.client.Validator.Query().CountX(m.ctx)
	if cnt != 1 {
		t.Error("expected 1 validator")
	}
	err := m.DeleteValidator(vals[0])
	if err != nil {
		t.Error(err)
	}
	cnt = m.client.Validator.Query().CountX(m.ctx)
	if cnt != 0 {
		t.Error("expected 0 validator")
	}
}

func TestValidatorManager_GetActive(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	m := newTestValidatorManager(t)
	addValidators(m, chains)

	vals := m.GetActive()
	if len(vals) != len(chains) {
		t.Error("expected all validators to be active")
	}
	if vals[0].Edges.Chain.ID != chains[0].ID {
		t.Error("expected chain id to be 1")
	}
}

func TestValidatorManager_GetByMoniker(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	m := newTestValidatorManager(t)
	vals := addValidators(m, chains)
	for _, v := range vals[:2] {
		_, err := m.client.Validator.
			UpdateOne(v).
			SetMoniker("validator").
			Save(m.ctx)
		if err != nil {
			t.Error(err)
		}
	}

	byMoniker := m.GetByMoniker("validator")
	if len(byMoniker) != 2 {
		t.Error("expected 2 validators")
	}
	for _, v := range byMoniker {
		if v.Moniker != "validator" {
			t.Error("expected moniker to be validator")
		}
		if !slices.ContainsFunc(chains[:2], func(c *ent.Chain) bool {
			return c.ID == v.Edges.Chain.ID
		}) {
			t.Error("expected validator to be from one of the chains")
		}
	}
}

func TestValidatorManager_GetForUser_Discord(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 2, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users[:1])
	addressTrackers := addAddressTrackers(newTestAddressTrackerManager(t), []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels, []*ent.TelegramChat{})

	m := newTestValidatorManager(t)
	vals := addValidators(m, chains)

	forUser, err := m.GetForUser(users[0])
	if err != nil {
		t.Error(err)
	}
	if len(forUser) != 0 {
		t.Errorf("expected 0 validators, got %d", len(forUser))
	}
	channels[0].Update().AddValidators(vals[0]).ExecX(m.ctx)
	forUser, err = m.GetForUser(users[0])
	if err != nil {
		t.Error(err)
	}
	if len(forUser) != 1 {
		t.Error("expected 1 validator")
	}
	if len(forUser[0].Edges.AddressTrackers) != 0 {
		t.Error("expected 0 address trackers")
	}

	forUser, err = m.GetForUser(users[1])
	if err != nil {
		t.Error(err)
	}
	if len(forUser) != 0 {
		t.Error("expected 0 validators")
	}

	vals[0].Update().AddAddressTrackers(addressTrackers...).ExecX(m.ctx)

	forUser, err = m.GetForUser(users[0])
	if err != nil {
		t.Error(err)
	}
	if len(forUser[0].Edges.AddressTrackers) != len(addressTrackers) {
		t.Errorf("expected %d address trackers", len(addressTrackers))
	}
}

func TestValidatorManager_GetForUser_Telegram(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 2, user.TypeTelegram)
	telegramChats := addTelegramChats(newTestTelegramChatManager(t), users[:1])
	addressTrackers := addAddressTrackers(newTestAddressTrackerManager(t), []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, []*ent.DiscordChannel{}, telegramChats)

	m := newTestValidatorManager(t)
	vals := addValidators(m, chains)

	forUser, err := m.GetForUser(users[0])
	if err != nil {
		t.Error(err)
	}
	if len(forUser) != 0 {
		t.Error("expected 0 validators")
	}
	telegramChats[0].Update().AddValidators(vals[0]).ExecX(m.ctx)
	forUser, err = m.GetForUser(users[0])
	if err != nil {
		t.Error(err)
	}
	if len(forUser) != 1 {
		t.Errorf("expected 1 validator, got %d", len(forUser))
	}
	if len(forUser[0].Edges.AddressTrackers) != 0 {
		t.Error("expected 0 address trackers")
	}

	forUser, err = m.GetForUser(users[1])
	if err != nil {
		t.Error(err)
	}
	if len(forUser) != 0 {
		t.Error("expected 0 validators")
	}

	vals[0].Update().AddAddressTrackers(addressTrackers...).ExecX(m.ctx)

	forUser, err = m.GetForUser(users[0])
	if err != nil {
		t.Error(err)
	}
	if len(forUser[0].Edges.AddressTrackers) != len(addressTrackers) {
		t.Errorf("expected %d address trackers", len(addressTrackers))
	}
}

func TestValidatorManager_TrackValidator(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 2, user.TypeDiscord)
	m := newTestValidatorManager(t)
	vals := addValidators(m, chains)

	_, err := m.TrackValidator(users[0], vals[0], 0, 0, 0)
	if err == nil {
		t.Error("expected error")
	}
	_, err = m.TrackValidator(users[0], vals[0], 1, 1, 0)
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidatorManager_TrackValidator_Discord(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 2, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users[:2])
	m := newTestValidatorManager(t)
	vals := addValidators(m, chains)

	tracker, err := m.TrackValidator(users[0], vals[0], channels[0].ID, 0, 100)
	if err != nil {
		t.Error(err)
	}
	if tracker == nil {
		t.Error("expected tracker")
	}
	if tracker.Address != vals[0].Address {
		t.Error("expected address to match")
	}
	if tracker.Edges.DiscordChannel == nil && tracker.Edges.DiscordChannel.ID != channels[0].ID {
		t.Error("expected discord channel to match")
	}
	if tracker.Edges.TelegramChat != nil {
		t.Error("expected telegram chat to be nil")
	}
	if tracker.QueryValidator().FirstX(m.ctx).ID != vals[0].ID {
		t.Error("expected validator to match")
	}
	if tracker.Edges.Chain == nil && tracker.Edges.Chain.ID != chains[0].ID {
		t.Error("expected chain to match")
	}
	if m.client.AddressTracker.Query().CountX(m.ctx) != 1 {
		t.Errorf("expected 1 address tracker, got %d", m.client.AddressTracker.Query().CountX(m.ctx))
	}

	addAddressTrackers(newTestAddressTrackerManager(t), []string{vals[0].Address}, channels[1:2], []*ent.TelegramChat{})
	tracker, err = m.TrackValidator(users[0], vals[0], channels[1].ID, 0, 100)
	if err != nil {
		t.Error(err)
	}
	if tracker == nil {
		t.Error("expected tracker")
	}
	if tracker.Address != vals[0].Address {
		t.Error("expected address to match")
	}
	if tracker.Edges.DiscordChannel == nil && tracker.Edges.DiscordChannel.ID != channels[1].ID {
		t.Error("expected discord channel to match")
	}
	if tracker.Edges.TelegramChat != nil {
		t.Error("expected telegram chat to be nil")
	}
	if tracker.QueryValidator().FirstX(m.ctx).ID != vals[0].ID {
		t.Error("expected validator to match")
	}
	if tracker.Edges.Chain == nil && tracker.Edges.Chain.ID != chains[0].ID {
		t.Error("expected chain to match")
	}
	if m.client.AddressTracker.Query().CountX(m.ctx) != 2 {
		t.Errorf("expected 2 address trackers, got %d", m.client.AddressTracker.Query().CountX(m.ctx))
	}
}

func TestValidatorManager_TrackValidator_Telegram(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 2, user.TypeTelegram)
	telegramChats := addTelegramChats(newTestTelegramChatManager(t), users[:2])
	m := newTestValidatorManager(t)
	vals := addValidators(m, chains)

	tracker, err := m.TrackValidator(users[0], vals[0], 0, telegramChats[0].ID, 100)
	if err != nil {
		t.Error(err)
	}
	if tracker == nil {
		t.Error("expected tracker")
	}
	if tracker.Address != vals[0].Address {
		t.Error("expected address to match")
	}
	if tracker.Edges.DiscordChannel != nil {
		t.Error("expected discord channel to be nil")
	}
	if tracker.Edges.TelegramChat == nil && tracker.Edges.TelegramChat.ID != telegramChats[0].ID {
		t.Error("expected telegram chat to match")
	}
	if tracker.QueryValidator().FirstX(m.ctx).ID != vals[0].ID {
		t.Error("expected validator to match")
	}
	if tracker.Edges.Chain == nil && tracker.Edges.Chain.ID != chains[0].ID {
		t.Error("expected chain to match")
	}
	if m.client.AddressTracker.Query().CountX(m.ctx) != 1 {
		t.Errorf("expected 1 address tracker, got %d", m.client.AddressTracker.Query().CountX(m.ctx))
	}

	addAddressTrackers(newTestAddressTrackerManager(t), []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, []*ent.DiscordChannel{}, telegramChats[1:2])
	tracker, err = m.TrackValidator(users[0], vals[0], 0, telegramChats[1].ID, 100)
	if err != nil {
		t.Error(err)
	}
	if tracker == nil {
		t.Error("expected tracker")
	}
	if tracker.Address != vals[0].Address {
		t.Error("expected address to match")
	}
	if tracker.Edges.DiscordChannel != nil {
		t.Error("expected discord channel to be nil")
	}
	if tracker.Edges.TelegramChat == nil && tracker.Edges.TelegramChat.ID != telegramChats[1].ID {
		t.Error("expected telegram chat to match")
	}
	if tracker.QueryValidator().FirstX(m.ctx).ID != vals[0].ID {
		t.Error("expected validator to match")
	}
	if tracker.Edges.Chain == nil && tracker.Edges.Chain.ID != chains[0].ID {
		t.Error("expected chain to match")
	}
	if m.client.AddressTracker.Query().CountX(m.ctx) != 2 {
		t.Errorf("expected 2 address trackers, got %d", m.client.AddressTracker.Query().CountX(m.ctx))
	}
}

func TestValidatorManager_UntrackValidator_Discord(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 2, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users[:2])
	m := newTestValidatorManager(t)
	vals := addValidators(m, chains)

	if vals[0].QueryDiscordChannels().CountX(m.ctx) != 0 {
		panic("expected no discord channels")
	}
	deletedIds, err := m.UntrackValidator(users[0], vals[0])
	if err != nil {
		t.Error(err)
	}
	if len(deletedIds) != 0 {
		t.Error("expected no deleted ids")
	}

	trackers := addAddressTrackers(newTestAddressTrackerManager(t), []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels[:1], []*ent.TelegramChat{})
	trackers[0].Update().SetValidator(vals[0]).SaveX(m.ctx)

	deletedIds, err = m.UntrackValidator(users[0], vals[0])
	if err != nil {
		t.Error(err)
	}
	if len(deletedIds) != 1 {
		t.Errorf("expected 1 deleted id, got %d", len(deletedIds))
	}
	if deletedIds[0] != trackers[0].ID {
		t.Error("expected deleted id to match")
	}
}

func TestValidatorManager_UntrackValidator_Telegram(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 2, user.TypeTelegram)
	telegramChats := addTelegramChats(newTestTelegramChatManager(t), users[:2])
	m := newTestValidatorManager(t)
	vals := addValidators(m, chains)

	if vals[0].QueryTelegramChats().CountX(m.ctx) != 0 {
		panic("expected no telegram chats")
	}
	deletedIds, err := m.UntrackValidator(users[0], vals[0])
	if err != nil {
		t.Error(err)
	}
	if len(deletedIds) != 0 {
		t.Error("expected no deleted ids")
	}

	trackers := addAddressTrackers(newTestAddressTrackerManager(t), []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, []*ent.DiscordChannel{}, telegramChats[:1])
	trackers[0].Update().SetValidator(vals[0]).SaveX(m.ctx)

	deletedIds, err = m.UntrackValidator(users[0], vals[0])
	if err != nil {
		t.Error(err)
	}
	if len(deletedIds) != 1 {
		t.Errorf("expected 1 deleted id, got %d", len(deletedIds))
	}
	if deletedIds[0] != trackers[0].ID {
		t.Error("expected deleted id to match")
	}
}

func TestValidatorManager_CascadeDelete(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 2, user.TypeTelegram)
	telegramChats := addTelegramChats(newTestTelegramChatManager(t), users[:2])
	m := newTestValidatorManager(t)
	vals := addValidators(m, chains)
	tracker, err := m.TrackValidator(users[0], vals[0], 0, telegramChats[0].ID, 100)
	if err != nil {
		panic(err)
	}
	if tracker == nil {
		panic(err)
	}
	m.client.Validator.
		DeleteOne(vals[0]).
		ExecX(m.ctx)
	if m.client.AddressTracker.Query().CountX(m.ctx) != 1 {
		t.Errorf("expected 1 address trackers, got %d", m.client.AddressTracker.Query().CountX(m.ctx))
	}

	m.client.Chain.DeleteOne(chains[0]).ExecX(m.ctx)
	if m.client.AddressTracker.Query().CountX(m.ctx) != 0 {
		t.Errorf("expected 0 address trackers, got %d", m.client.AddressTracker.Query().CountX(m.ctx))
	}
}
