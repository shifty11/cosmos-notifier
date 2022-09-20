package database

import (
	"context"
	"fmt"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/types"
	"testing"
)

func newTestSubscriptionManager(t *testing.T) *SubscriptionManager {
	manager := NewSubscriptionManager(
		testClient(t),
		context.Background(),
		newTestUserManager(t),
		newTestContractManager(t),
		newTestTelegramChatManager(t),
		newTestDiscordChannelManager(t),
	)
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

var compStringTg = ""
var compStringDiscord = ""

type mockTelegramChatManager struct{}

func (m *mockTelegramChatManager) AddOrRemoveContract(tgChatId int64, contractId int) (hasContract bool, err error) {
	compStringTg = fmt.Sprintf("%T %d %d", m, tgChatId, contractId)
	return true, nil
}

func (m *mockTelegramChatManager) CreateOrUpdateChat(userId int64, userName string, tgChatId int64, name string, isGroup bool) (tc *ent.TelegramChat, created bool) {
	panic("implement me")
}

func (m *mockTelegramChatManager) GetSubscribedIds(entContract *ent.Contract) []TgChatQueryResult {
	panic("implement me")
}

func (m *mockTelegramChatManager) Delete(userId int64, chatId int64) error {
	panic("implement me")
}

func (m *mockTelegramChatManager) DeleteMultiple(chatIds []int64) {
	panic("implement me")
}

type mockDiscordChannelManager struct{}

func (m *mockDiscordChannelManager) AddOrRemoveContract(discordChannelId int64, contractId int) (hasContract bool, err error) {
	compStringDiscord = fmt.Sprintf("%T %d %d", m, discordChannelId, contractId)
	return true, nil
}

func (m *mockDiscordChannelManager) CreateOrUpdateChannel(userId int64, userName string, channelId int64, name string, isGroup bool) (dc *ent.DiscordChannel, created bool) {
	panic("implement me")
}

func (m *mockDiscordChannelManager) Delete(userId int64, channelId int64) error {
	panic("implement me")
}

func (m *mockDiscordChannelManager) GetChannelUsers(channelId int64) []*ent.User {
	panic("implement me")
}

func (m *mockDiscordChannelManager) CountSubscriptions(channelId int64) int {
	panic("implement me")
}

func (m *mockDiscordChannelManager) GetSubscribedIds(entContract *ent.Contract) []DiscordChannelQueryResult {
	panic("implement me")
}

func (m *mockDiscordChannelManager) DeleteMultiple(channelIds []int64) {
	panic("implement me")
}

func TestSubscriptionManager_ToggleSubscription(t *testing.T) {
	m := newTestSubscriptionManager(t)
	m.telegramChatManager = &mockTelegramChatManager{}
	m.discordChannelManager = &mockDiscordChannelManager{}

	tgUser := m.userManager.createOrUpdateUser(1, "username", user.TypeTelegram)

	//goland:noinspection GoUnhandledErrorResult
	m.ToggleSubscription(tgUser, int64(1), 1)
	if compStringTg != fmt.Sprintf("%T %d %d", m.telegramChatManager, int64(1), 1) {
		t.Error("Wrong function called")
	}

	discordUser := m.userManager.createOrUpdateUser(1, "username", user.TypeDiscord)

	//goland:noinspection GoUnhandledErrorResult
	m.ToggleSubscription(discordUser, int64(1), 1)
	if compStringDiscord != fmt.Sprintf("%T %d %d", m.discordChannelManager, int64(1), 1) {
		t.Error("Wrong function called")
	}

}

func TestSubscriptionManager_getSubscriptions(t *testing.T) {
	m := newTestSubscriptionManager(t)

	data1 := &types.ContractData{
		Address:     "0x123",
		Name:        "contract1",
		Description: "desc1",
		ImageUrl:    "url1",
	}
	data2 := &types.ContractData{
		Address:     "0x456",
		Name:        "contract2",
		Description: "desc2",
		ImageUrl:    "url2",
	}
	c2, _ := m.contractManager.Create(data2)
	c1, _ := m.contractManager.Create(data1)

	subscriptions := m.getSubscriptions([]*ent.Contract{c1})
	if len(subscriptions) != 2 {
		t.Fatalf("Expected 2 subscriptions, got %d", len(subscriptions))
	}
	if subscriptions[0].ContractAddress != data1.Address {
		t.Errorf("Expected %s, got %s", data1.Address, subscriptions[0].ContractAddress)
	}
	if subscriptions[1].ContractAddress != data2.Address {
		t.Errorf("Expected %s, got %s", data2.Address, subscriptions[1].ContractAddress)
	}
	if subscriptions[0].Name != data1.Name {
		t.Errorf("Expected %s, got %s", data1.Name, subscriptions[0].Name)
	}
	if subscriptions[1].Name != data2.Name {
		t.Errorf("Expected %s, got %s", data2.Name, subscriptions[1].Name)
	}
	if subscriptions[0].Id != int64(c1.ID) {
		t.Errorf("Expected %d, got %d", c1.ID, subscriptions[0].Id)
	}
	if subscriptions[1].Id != int64(c2.ID) {
		t.Errorf("Expected %d, got %d", c2.ID, subscriptions[1].Id)
	}
	if !subscriptions[0].Notify {
		t.Error("Expected notify to be true")
	}
	if subscriptions[1].Notify {
		t.Error("Expected notify to be false")
	}

}

func TestSubscriptionManager_GetSubscriptions(t *testing.T) {
	m := newTestSubscriptionManager(t)

	data1 := &types.ContractData{
		Address:     "0x123",
		Name:        "contract1",
		Description: "desc1",
		ImageUrl:    "url1",
	}
	data2 := &types.ContractData{
		Address:     "0x456",
		Name:        "contract2",
		Description: "desc2",
		ImageUrl:    "url2",
	}
	c2, _ := m.contractManager.Create(data2)
	c1, _ := m.contractManager.Create(data1)

	m.telegramChatManager.CreateOrUpdateChat(1, "telegramuser", 10, "chat2", true)
	m.telegramChatManager.CreateOrUpdateChat(1, "telegramuser", 11, "chat1", false)
	m.discordChannelManager.CreateOrUpdateChannel(1, "discorduser", 12, "channel1", false)

	tgUser1, _ := m.userManager.Get(1, user.TypeTelegram)
	m.userManager.Get(2, user.TypeTelegram)
	discordUser, _ := m.userManager.Get(1, user.TypeDiscord)

	_, err := m.ToggleSubscription(tgUser1, int64(10), c1.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.ToggleSubscription(discordUser, int64(12), c2.ID)
	if err != nil {
		t.Fatal(err)
	}

	chatRooms := m.GetSubscriptions(tgUser1)
	if len(chatRooms) != 2 {
		t.Fatalf("Expected 2 chatRooms, got %d", len(chatRooms))
	}
	ch1, ch2 := chatRooms[0], chatRooms[1]
	if ch1.Id != int64(11) && ch2.Id != int64(10) {
		t.Errorf("Expected 11 and 10, got %d and %d", ch1.Id, ch2.Id)
	}
	if ch1.Name != "chat1" && ch2.Name != "chat2" {
		t.Errorf("Expected chat1 and chat2, got %s and %s", ch1.Name, ch2.Name)
	}

	if len(chatRooms[0].Subscriptions) != 2 {
		t.Errorf("Expected 2 subscriptions, got %d", len(chatRooms[0].Subscriptions))
	}
	c1s1, c1s2, c2s1, c2s2 := chatRooms[0].Subscriptions[0], chatRooms[0].Subscriptions[1], chatRooms[1].Subscriptions[0], chatRooms[1].Subscriptions[1]
	if c1s1.Name != data1.Name && c2s1.Name != data1.Name && c1s2.Name != data2.Name && c2s2.Name != data2.Name {
		t.Errorf("Expected %s or %s, got %s and %s", data1.Name, data2.Name, c1s1.Name, c1s2.Name)
	}
	if c1s1.Id != int64(c1.ID) && c2s1.Id != int64(c1.ID) && c1s2.Id != int64(c2.ID) && c2s2.Id != int64(c2.ID) {
		t.Errorf("Expected %d or %d, got %d and %d", c1.ID, c2.ID, c1s1.Id, c1s2.Id)
	}
	if c1s1.Notify != false && c2s1.Notify != false && c1s2.Notify != true && c2s2.Notify != false {
		t.Errorf("Expected %v or %v, got %v and %v", false, true, c1s1.Notify, c1s2.Notify)
	}

	chatRooms = m.GetSubscriptions(discordUser)
	if len(chatRooms) != 1 {
		t.Errorf("Expected 1 chatRooms, got %d", len(chatRooms))
	}
	ch1 = chatRooms[0]
	if ch1.Id != int64(12) {
		t.Errorf("Expected 12, got %d", ch1.Id)
	}
	if ch1.Name != "channel1" {
		t.Errorf("Expected channel1, got %s", ch1.Name)
	}

	if len(chatRooms[0].Subscriptions) != 2 {
		t.Errorf("Expected 1 subscriptions, got %d", len(chatRooms[0].Subscriptions))
	}
	c1s1, c1s2 = chatRooms[0].Subscriptions[0], chatRooms[0].Subscriptions[1]
	if c1s1.Name != data2.Name && c1s2.Name != data2.Name {
		t.Errorf("Expected %s, got %s and %s", data2.Name, c1s1.Name, c1s2.Name)
	}
	if c1s1.Id != int64(c2.ID) && c1s2.Id != int64(c2.ID) {
		t.Errorf("Expected %d, got %d and %d", c2.ID, c1s1.Id, c1s2.Id)
	}
	if c1s1.Notify != false && c1s2.Notify != true {
		t.Errorf("Expected %v, got %v and %v", false, c1s1.Notify, c1s2.Notify)
	}

}
