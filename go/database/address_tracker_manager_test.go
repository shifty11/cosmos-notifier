package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"testing"
)

func newTestAddressTrackerManager(t *testing.T) *AddressTrackerManager {
	manager := NewAddressTrackerManager(testClient(t), context.Background())
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func addAddressTrackers(m *AddressTrackerManager, addresses []string, discordChannels []*ent.DiscordChannel, telegramChats []*ent.TelegramChat) []*ent.AddressTracker {
	var addressTrackers []*ent.AddressTracker
	for _, address := range addresses {
		for _, discordChannel := range discordChannels {
			addressTracker, err := m.AddTracker(address, discordChannel.ID, 0)
			addressTrackers = append(addressTrackers, addressTracker)
			if err != nil {
				panic(err)
			}
		}
		for _, telegramChat := range telegramChats {
			addressTracker, err := m.AddTracker(address, 0, telegramChat.ID)
			addressTrackers = append(addressTrackers, addressTracker)
			if err != nil {
				panic(err)
			}
		}
	}
	return addressTrackers
}

func TestAddressTrackerManager_IsValid(t *testing.T) {
	m := newTestAddressTrackerManager(t)

	addChains(newTestChainManager(t))

	if isValid, _ := m.IsValid(""); isValid {
		t.Error("Empty address is valid")
	}
	if isValid, _ := m.IsValid("juno1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"); isValid {
		t.Error("Address of unknown chain is valid")
	}
	if isValid, _ := m.IsValid("cosmos1"); isValid {
		t.Error("Invalid address is valid")
	}

	addresses := [][]string{
		{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", "Cosmos"},
		{"osmo166y8reslaeuedyc6gd83m8r5p0pmdnvq0dggsq", "Osmosis"},
		{"comdex1cx82d7pm4dgffy7a93rl6ul5g84vjgxkqfyp2m", "Comdex"},
	}
	for _, address := range addresses {
		isValid, chain := m.IsValid(address[0])
		if !isValid || chain.Name != address[1] {
			t.Errorf("Address %s of chain %s is invalid", address[0], address[1])
		}
	}
}

func TestAddressTrackerManager_AddTrackerForDiscordChannel(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)

	users := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users)

	m := newTestAddressTrackerManager(t)

	if _, err := m.AddTracker("", 0, 0); err == nil {
		t.Error("Empty address is valid")
	}
	if _, err := m.AddTracker("cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 0, 0); err == nil {
		t.Error("Empty discordChannelId and telegramChatId is valid")
	}
	if _, err := m.AddTracker("cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 1, 1); err == nil {
		t.Error("Both discordChannelId and telegramChatId are valid")
	}

	tracker, err := m.AddTracker("cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", channels[0].ID, 0)
	if err != nil {
		t.Error(err)
	}
	if tracker.Address != "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02" {
		t.Error("Address is not saved")
	}
	if tracker.QueryChain().FirstX(m.ctx).Name != "Cosmos" {
		t.Error("Chain is not saved")
	}
	if tracker.QueryChainProposals().CountX(m.ctx) != 0 {
		t.Error("Chain proposals are saved")
	}
	if tracker.QueryDiscordChannel().FirstX(m.ctx).ID != channels[0].ID {
		t.Error("Discord channel is not saved")
	}
	if tracker.QueryTelegramChat().CountX(m.ctx) != 0 {
		t.Error("Telegram chat is saved")
	}
}

func TestAddressTrackerManager_AddTrackerForTelegramChat(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)

	users := addUsers(newTestUserManager(t), 1, user.TypeTelegram)
	chats := addTelegramChats(newTestTelegramChatManager(t), users)

	m := newTestAddressTrackerManager(t)

	tracker, err := m.AddTracker("cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 0, chats[0].ID)
	if err != nil {
		t.Error(err)
	}
	if tracker.Address != "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02" {
		t.Error("Address is not saved")
	}
	if tracker.QueryChain().FirstX(m.ctx).Name != "Cosmos" {
		t.Error("Chain is not saved")
	}
	if tracker.QueryChainProposals().CountX(m.ctx) != 0 {
		t.Error("Chain proposals are saved")
	}
	if tracker.QueryDiscordChannel().CountX(m.ctx) != 0 {
		t.Error("Discord channel is saved")
	}
	if tracker.QueryTelegramChat().FirstX(m.ctx).ID != chats[0].ID {
		t.Error("Telegram chat is not saved")
	}
}

func TestAddressTracker_NoUser(t *testing.T) {
	addChains(newTestChainManager(t))
	manager := newTestAddressTrackerManager(t)
	_, err := manager.AddTracker("cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 1, 0)
	if err == nil {
		t.Error("Address is added without user")
	}

	users := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users)
	_, err = manager.AddTracker("cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", channels[0].ID, 0)
	if err != nil {
		t.Error(err)
	}
}

func TestAddressTracker_CascadeDeleteForChain(t *testing.T) {
	addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users)
	m := newTestAddressTrackerManager(t)

	addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels, []*ent.TelegramChat{})

	m.client.Chain.
		Delete().
		ExecX(m.ctx)

	if m.client.AddressTracker.Query().CountX(m.ctx) != 0 {
		t.Error("AddressTracker is not deleted")
	}
}

func TestAddressTracker_CascadeDeleteForDiscordChannel(t *testing.T) {
	addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users)
	m := newTestAddressTrackerManager(t)

	addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels, []*ent.TelegramChat{})

	m.client.DiscordChannel.
		Delete().
		ExecX(m.ctx)

	if m.client.AddressTracker.Query().CountX(m.ctx) != 0 {
		t.Error("AddressTracker is not deleted")
	}
}

func TestAddressTracker_CascadeDeleteForTelegramChat(t *testing.T) {
	addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 1, user.TypeTelegram)
	chats := addTelegramChats(newTestTelegramChatManager(t), users)
	m := newTestAddressTrackerManager(t)

	addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, []*ent.DiscordChannel{}, chats)

	m.client.TelegramChat.
		Delete().
		ExecX(m.ctx)

	if m.client.AddressTracker.Query().CountX(m.ctx) != 0 {
		t.Error("AddressTracker is not deleted")
	}
}
