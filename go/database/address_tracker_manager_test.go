package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/chainproposal"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"testing"
	"time"
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
			userEnt := discordChannel.QueryUsers().FirstX(m.ctx)
			addressTracker, err := m.Create(m.ctx, userEnt, address, discordChannel.ID, 0, 10000)
			addressTrackers = append(addressTrackers, addressTracker)
			if err != nil {
				panic(err)
			}
		}
		for _, telegramChat := range telegramChats {
			userEnt := telegramChat.QueryUsers().FirstX(m.ctx)
			addressTracker, err := m.Create(m.ctx, userEnt, address, 0, telegramChat.ID, 10000)
			addressTrackers = append(addressTrackers, addressTracker)
			if err != nil {
				panic(err)
			}
		}
	}
	return addressTrackers
}

func TestAddressTrackerManager_QueryByUser(t *testing.T) {
	addChains(newTestChainManager(t))
	discordUsers := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	telegramUsers := addUsers(newTestUserManager(t), 1, user.TypeTelegram)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), discordUsers[:1])
	telegramChats := addTelegramChats(newTestTelegramChatManager(t), telegramUsers[:1])

	m := newTestAddressTrackerManager(t)
	addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels, []*ent.TelegramChat{})
	addAddressTrackers(m, []string{"osmo166y8reslaeuedyc6gd83m8r5p0pmdnvq0dggsq", "comdex1cx82d7pm4dgffy7a93rl6ul5g84vjgxkqfyp2m"}, []*ent.DiscordChannel{}, telegramChats)

	trackers, err := m.QueryByUser(discordUsers[0])
	if err != nil {
		t.Fatal(err)
	}
	if len(trackers) != 2 {
		t.Errorf("Expected 2 tracker, got %d", len(trackers))
	}
	if trackers[0].Address != "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02" {
		t.Errorf("Expected address %s, got %s", "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", trackers[0].Address)
	}
	if trackers[0].Edges.DiscordChannel.ID != channels[0].ID {
		t.Errorf("Expected discord channel %d, got %d", channels[0].ID, trackers[0].Edges.DiscordChannel.ID)
	}
	if trackers[0].Edges.TelegramChat != nil {
		t.Error("Telegram chat is not nil")
	}
	if trackers[1].Address != "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02" {
		t.Errorf("Expected address %s, got %s", "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", trackers[1].Address)
	}
	if trackers[1].Edges.DiscordChannel.ID != channels[1].ID {
		t.Errorf("Expected discord channel %d, got %d", channels[1].ID, trackers[1].Edges.DiscordChannel.ID)
	}
	if trackers[1].Edges.TelegramChat != nil {
		t.Error("Telegram chat is not nil")
	}

	trackers, err = m.QueryByUser(telegramUsers[0])
	if err != nil {
		t.Fatal(err)
	}
	if len(trackers) != 4 {
		t.Errorf("Expected 4 tracker, got %d", len(trackers))
	}
	for _, tracker := range trackers {
		if tracker.Address != "osmo166y8reslaeuedyc6gd83m8r5p0pmdnvq0dggsq" && tracker.Address != "comdex1cx82d7pm4dgffy7a93rl6ul5g84vjgxkqfyp2m" {
			t.Errorf("Expected address %s or %s, got %s", "osmo166y8reslaeuedyc6gd83m8r5p0pmdnvq0dggsq", "comdex1cx82d7pm4dgffy7a93rl6ul5g84vjgxkqfyp2m", tracker.Address)
		}
		if tracker.Edges.DiscordChannel != nil {
			t.Error("Discord channel is not nil")
		}
		if tracker.Edges.TelegramChat.ID != telegramChats[0].ID && tracker.Edges.TelegramChat.ID != telegramChats[1].ID {
			t.Errorf("Expected telegram chat %d or %d, got %d", telegramChats[0].ID, telegramChats[1].ID, tracker.Edges.TelegramChat.ID)
		}
	}
}

func TestAddressTrackerManager_QueryIsValid(t *testing.T) {
	m := newTestAddressTrackerManager(t)

	addChains(newTestChainManager(t))

	if isValid, _ := m.QueryIsValid(""); isValid {
		t.Error("Empty address is valid")
	}
	if isValid, _ := m.QueryIsValid("juno1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"); isValid {
		t.Error("Address of unknown chain is valid")
	}
	if isValid, _ := m.QueryIsValid("cosmos1"); isValid {
		t.Error("Invalid address is valid")
	}

	addresses := [][]string{
		{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", "Cosmos"},
		{"osmo166y8reslaeuedyc6gd83m8r5p0pmdnvq0dggsq", "Osmosis"},
		{"comdex1cx82d7pm4dgffy7a93rl6ul5g84vjgxkqfyp2m", "Comdex"},
	}
	for _, address := range addresses {
		isValid, chain := m.QueryIsValid(address[0])
		if !isValid || chain.Name != address[1] {
			t.Errorf("Address %s of chain %s is invalid", address[0], address[1])
		}
	}
}

func TestAddressTrackerManager_Create_DiscordChannel(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)

	users := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users)

	m := newTestAddressTrackerManager(t)

	if _, err := m.Create(m.ctx, users[0], "", 0, 0, 10000); err == nil {
		t.Error("Empty address is valid")
	}
	if _, err := m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 0, 0, 10000); err == nil {
		t.Error("Empty discordChannelId and telegramChatId is valid")
	}
	if _, err := m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 1, 1, 10000); err == nil {
		t.Error("Both discordChannelId and telegramChatId are valid")
	}

	tracker, err := m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", channels[0].ID, 0, 10000)
	if err != nil {
		t.Error(err)
	}
	if tracker.Address != "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02" {
		t.Error("Address is not saved")
	}
	if tracker.NotificationInterval != 10000 {
		t.Error("Notification interval is not saved")
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

	if tracker.Edges.Chain == nil {
		t.Error("Chain is nil")
	}
	if tracker.Edges.DiscordChannel == nil {
		t.Error("Discord channel is nil")
	}
	if tracker.Edges.TelegramChat != nil {
		t.Error("Telegram chat is not nil")
	}
}

func TestAddressTrackerManager_Create_TelegramChat(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)

	users := addUsers(newTestUserManager(t), 1, user.TypeTelegram)
	chats := addTelegramChats(newTestTelegramChatManager(t), users)

	m := newTestAddressTrackerManager(t)

	tracker, err := m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 0, chats[0].ID, 10000)
	if err != nil {
		t.Error(err)
	}
	if tracker.Address != "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02" {
		t.Error("Address is not saved")
	}
	if tracker.NotificationInterval != 10000 {
		t.Error("Notification interval is not saved")
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

	if tracker.Edges.Chain == nil {
		t.Error("Chain is nil")
	}
	if tracker.Edges.DiscordChannel != nil {
		t.Error("Discord channel is not nil")
	}
	if tracker.Edges.TelegramChat == nil {
		t.Error("Telegram chat is nil")
	}
}

func TestAddressTrackerManager_Create_DiscordChannel_Twice(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)

	users := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users)

	m := newTestAddressTrackerManager(t)

	_, err := m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", channels[0].ID, 0, 10000)
	if err != nil {
		t.Error(err)
	}
	_, err = m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", channels[0].ID, 0, 10000)
	if err != nil {
		t.Error(err)
	}
	if m.client.AddressTracker.Query().CountX(m.ctx) != 2 {
		t.Error("Address tracker is not added twice")
	}
}

func TestAddressTrackerManager_Create_TelegramChat_Twice(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)

	users := addUsers(newTestUserManager(t), 1, user.TypeTelegram)
	chats := addTelegramChats(newTestTelegramChatManager(t), users)

	m := newTestAddressTrackerManager(t)

	_, err := m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 0, chats[0].ID, 10000)
	if err != nil {
		t.Error(err)
	}
	_, err = m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 0, chats[0].ID, 10000)
	if err != nil {
		t.Error(err)
	}
	if m.client.AddressTracker.Query().CountX(m.ctx) != 2 {
		t.Error("Address tracker is not added twice")
	}
}

func TestAddressTracker_NoDiscordChannel(t *testing.T) {
	addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 1, user.TypeDiscord)

	m := newTestAddressTrackerManager(t)
	_, err := m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 1, 0, 10000)
	if err == nil {
		t.Error("Address is added without user")
	}

	channels := addDiscordChannels(newTestDiscordChannelManager(t), users)
	_, err = m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", channels[0].ID, 0, 10000)
	if err != nil {
		t.Error(err)
	}
}

func TestAddressTracker_NoTelegramChat(t *testing.T) {
	addChains(newTestChainManager(t))
	users := addUsers(newTestUserManager(t), 1, user.TypeTelegram)

	m := newTestAddressTrackerManager(t)
	_, err := m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 0, 1, 10000)
	if err == nil {
		t.Error("Address is added without user")
	}

	chats := addTelegramChats(newTestTelegramChatManager(t), users)
	_, err = m.Create(m.ctx, users[0], "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", 0, chats[0].ID, 10000)
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

func TestAddressTracker_Update_Discord(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)

	users := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users)
	m := newTestAddressTrackerManager(t)

	trackers := addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels, []*ent.TelegramChat{})

	tracker, err := m.Update(users[0], trackers[0].ID, 2, 0, 999)
	if err != nil {
		t.Error(err)
	}
	if tracker.Address != "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02" {
		t.Errorf("Wrong address: %s", tracker.Address)
	}
	if tracker.QueryDiscordChannel().FirstX(m.ctx).ID != 2 {
		t.Errorf("Wrong discord channel id: %d", tracker.QueryDiscordChannel().FirstX(m.ctx).ID)
	}
	if tracker.QueryTelegramChat().ExistX(m.ctx) {
		t.Errorf("Telegram chat should be empty")
	}
	if tracker.NotificationInterval != 999 {
		t.Errorf("Wrong notification interval: %d", tracker.NotificationInterval)
	}
}

func TestAddressTracker_Update_Telegram(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)

	users := addUsers(newTestUserManager(t), 1, user.TypeTelegram)
	tgChats := addTelegramChats(newTestTelegramChatManager(t), users)
	m := newTestAddressTrackerManager(t)

	trackers := addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, []*ent.DiscordChannel{}, tgChats)

	tracker, err := m.Update(users[0], trackers[0].ID, 0, 2, 999)
	if err != nil {
		t.Error(err)
	}
	if tracker.Address != "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02" {
		t.Errorf("Wrong address: %s", tracker.Address)
	}
	if tracker.QueryDiscordChannel().ExistX(m.ctx) {
		t.Errorf("Discord channel should be empty")
	}
	if tracker.QueryTelegramChat().FirstX(m.ctx).ID != 2 {
		t.Errorf("Wrong telegram chat id: %d", tracker.QueryTelegramChat().FirstX(m.ctx).ID)
	}
	if tracker.NotificationInterval != 999 {
		t.Errorf("Wrong notification interval: %d", tracker.NotificationInterval)
	}
}

func TestAddressTracker_Update_Failure(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)

	users := addUsers(newTestUserManager(t), 1, user.TypeTelegram)
	tgChats := addTelegramChats(newTestTelegramChatManager(t), users)
	m := newTestAddressTrackerManager(t)

	trackers := addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, []*ent.DiscordChannel{}, tgChats)

	_, err := m.Update(users[0], 999, 0, 2, 999)
	if err == nil {
		t.Error("Should fail")
	}

	_, err = m.Update(users[0], trackers[0].ID, 0, 0, 999)
	if err == nil {
		t.Error("Should fail")
	}

	_, err = m.Update(users[0], trackers[0].ID, 999, 2, 999)
	if err == nil {
		t.Error("Should fail")
	}

	_, err = m.Update(users[0], trackers[0].ID, 0, 999, 999)
	if err == nil {
		t.Error("Should fail")
	}

	_, err = m.Update(users[0], trackers[0].ID, 0, 2, -1)
	if err == nil {
		t.Error("Should fail")
	}
}

func TestAddressTracker_QueryUnnotifiedTrackers(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)
	discordUsers := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	telegramUsers := addUsers(newTestUserManager(t), 1, user.TypeTelegram)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), discordUsers)
	tgChats := addTelegramChats(newTestTelegramChatManager(t), telegramUsers)
	m := newTestAddressTrackerManager(t)

	trackers := addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels, tgChats)
	for _, tracker := range trackers {
		twoWeeksDuration := time.Hour * 24 * 14
		tracker.Update().
			SetNotificationInterval(int64(twoWeeksDuration.Seconds())).
			ExecX(m.ctx)
	}

	unnotifiedTrackers := m.QueryUnnotifiedTrackers()
	if len(unnotifiedTrackers) != 4 {
		t.Error("Wrong number of unnotifiedTrackers")
	}
	for _, tracker := range unnotifiedTrackers {
		if tracker.AddressTracker.QueryChain().FirstX(m.ctx).Name != "Cosmos" {
			t.Error("Wrong chain")
		}
		if tracker.ChainProposal.Status != chainproposal.StatusPROPOSAL_STATUS_VOTING_PERIOD {
			t.Error("Wrong chain proposal status")
		}
	}
}

func TestAddressTracker_QueryUnnotifiedTrackers_CheckTime(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)
	users := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users)
	tgChats := addTelegramChats(newTestTelegramChatManager(t), users)
	m := newTestAddressTrackerManager(t)

	trackers := addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels, tgChats)

	unnotifiedTrackers := m.QueryUnnotifiedTrackers()
	if len(unnotifiedTrackers) != 0 {
		t.Error("Wrong number of unnotifiedTrackers")
	}

	twoWeeksDuration := time.Hour * 24 * 14
	trackers[0].Update().
		SetNotificationInterval(int64(twoWeeksDuration.Seconds())).
		ExecX(m.ctx)

	unnotifiedTrackers = m.QueryUnnotifiedTrackers()
	if len(unnotifiedTrackers) != 1 {
		t.Error("Wrong number of unnotifiedTrackers")
	}
}

func TestAddressTracker_QueryUnnotifiedTrackers_UpdateSetNotified(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)
	users := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), users)
	tgChats := addTelegramChats(newTestTelegramChatManager(t), users)
	m := newTestAddressTrackerManager(t)

	trackers := addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels, tgChats)

	twoWeeksDuration := time.Hour * 24 * 14
	trackers[0].Update().
		SetNotificationInterval(int64(twoWeeksDuration.Seconds())).
		ExecX(m.ctx)

	unnotifiedTrackers := m.QueryUnnotifiedTrackers()
	if len(unnotifiedTrackers) != 1 {
		t.Error("Wrong number of unnotifiedTrackers")
	}

	m.UpdateSetNotified(unnotifiedTrackers[0])
	unnotifiedTrackers = m.QueryUnnotifiedTrackers()
	if len(unnotifiedTrackers) != 0 {
		t.Error("Wrong number of unnotifiedTrackers")
	}
}

func TestAddressTracker_QueryChatRooms_Discord(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)
	dUsers := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	tgUsers := addUsers(newTestUserManager(t), 1, user.TypeTelegram)

	m := newTestAddressTrackerManager(t)

	discordChannels, telegramChats, err := m.QueryChatRooms(dUsers[0])
	if err == nil {
		t.Error(err)
	}

	addDiscordChannels(newTestDiscordChannelManager(t), dUsers[:1])
	addTelegramChats(newTestTelegramChatManager(t), tgUsers[:1])

	discordChannels, telegramChats, err = m.QueryChatRooms(dUsers[0])
	if err != nil {
		t.Error(err)
	}
	if len(discordChannels) != 2 {
		t.Error("Wrong number of discord channels")
	}
	if len(telegramChats) != 0 {
		t.Error("Wrong number of telegram chats")
	}

	discordChannels, telegramChats, err = m.QueryChatRooms(tgUsers[0])
	if err != nil {
		t.Error(err)
	}
	if len(discordChannels) != 0 {
		t.Error("Wrong number of discord channels")
	}
	if len(telegramChats) != 2 {
		t.Error("Wrong number of telegram chats")
	}
}

func TestAddressTracker_Delete(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)
	discordUsers := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	telegramUsers := addUsers(newTestUserManager(t), 1, user.TypeTelegram)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), discordUsers)
	tgChats := addTelegramChats(newTestTelegramChatManager(t), telegramUsers)
	m := newTestAddressTrackerManager(t)

	trackers := addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels, tgChats)
	if len(trackers) != 4 {
		t.Error("Wrong number of trackers")
	}

	err := m.Delete(discordUsers[0], trackers[0].ID)
	if err != nil {
		return
	}
	if m.client.AddressTracker.Query().CountX(m.ctx) != 3 {
		t.Error("Wrong number of trackers")
	}

	err = m.Delete(telegramUsers[0], trackers[0].ID)
	if err != nil {
		return
	}
	if m.client.AddressTracker.Query().CountX(m.ctx) != 3 {
		t.Error("Wrong number of trackers")
	}

	err = m.Delete(telegramUsers[0], trackers[2].ID)
	if err != nil {
		return
	}
	if m.client.AddressTracker.Query().CountX(m.ctx) != 2 {
		t.Error("Wrong number of trackers")
	}
}

func TestAddressTracker_AllByChatRoomsAndAddress(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)
	discordUsers := addUsers(newTestUserManager(t), 1, user.TypeDiscord)
	telegramUsers := addUsers(newTestUserManager(t), 1, user.TypeTelegram)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), discordUsers)
	tgChats := addTelegramChats(newTestTelegramChatManager(t), telegramUsers)
	m := newTestAddressTrackerManager(t)
	addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels, tgChats)

	trackers, err := m.QueryByChatRoomsAndAddress(channels[0].ID, 0, "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02")
	if err != nil {
		t.Error(err)
	}
	if len(trackers) != 1 {
		t.Error("Wrong number of trackers")
	}
	if trackers[0].Address != "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02" {
		t.Error("Wrong address")
	}
	if trackers[0].QueryDiscordChannel().FirstX(m.ctx).ID != channels[0].ID {
		t.Error("Wrong discord channel")
	}

	trackers, err = m.QueryByChatRoomsAndAddress(0, tgChats[0].ID, "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02")
	if err != nil {
		t.Error(err)
	}
	if len(trackers) != 1 {
		t.Error("Wrong number of trackers")
	}
	if trackers[0].Address != "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02" {
		t.Error("Wrong address")
	}
	if trackers[0].QueryTelegramChat().FirstX(m.ctx).ID != tgChats[0].ID {
		t.Error("Wrong telegram chat")
	}
}

func TestAddressTracker_QueryDoesExist(t *testing.T) {
	chains := addChains(newTestChainManager(t))
	addChainProposals(newTestChainProposalManager(t), chains)
	discordUsers := addUsers(newTestUserManager(t), 2, user.TypeDiscord)
	telegramUsers := addUsers(newTestUserManager(t), 2, user.TypeTelegram)
	channels := addDiscordChannels(newTestDiscordChannelManager(t), discordUsers)
	tgChats := addTelegramChats(newTestTelegramChatManager(t), telegramUsers)
	m := newTestAddressTrackerManager(t)
	addAddressTrackers(m, []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02"}, channels[:1], tgChats[:1])

	if !m.QueryDoesExist(channels[0].ID, 0, "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02") {
		t.Error("Tracker should exist")
	}
	if !m.QueryDoesExist(0, tgChats[0].ID, "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02") {
		t.Error("Tracker should exist")
	}
	if m.QueryDoesExist(channels[0].ID, 0, "osmo166y8reslaeuedyc6gd83m8r5p0pmdnvq0dggsq") {
		t.Error("Tracker should not exist")
	}
	if m.QueryDoesExist(0, tgChats[0].ID, "osmo166y8reslaeuedyc6gd83m8r5p0pmdnvq0dggsq") {
		t.Error("Tracker should not exist")
	}
	if m.QueryDoesExist(channels[1].ID, 0, "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02") {
		t.Error("Tracker should not exist")
	}
	if m.QueryDoesExist(0, tgChats[1].ID, "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02") {
		t.Error("Tracker should not exist")
	}
}
