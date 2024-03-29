package database

import (
	"context"
	"github.com/golang/mock/gomock"
	mock_database "github.com/shifty11/cosmos-notifier/database/mock_types"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/types"
	"testing"
)

func newTestSubscriptionManager(t *testing.T) *SubscriptionManager {
	manager := NewSubscriptionManager(
		testClient(t),
		context.Background(),
		newTestUserManager(t),
		newTestChainManager(t),
		newTestContractManager(t),
		newTestTelegramChatManager(t),
		newTestDiscordChannelManager(t),
	)
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func newTestSubscriptionManagerWithMocks(t *testing.T) (*SubscriptionManager, *gomock.Controller, *mock_database.MockITelegramChatManager, *mock_database.MockIDiscordChannelManager) {
	ctrl := gomock.NewController(t)
	tgMock := mock_database.NewMockITelegramChatManager(ctrl)
	dMock := mock_database.NewMockIDiscordChannelManager(ctrl)
	manager := NewSubscriptionManager(
		testClient(t),
		context.Background(),
		newTestUserManager(t),
		newTestChainManager(t),
		newTestContractManager(t),
		tgMock,
		dMock,
	)
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager, ctrl, tgMock, dMock
}

func TestSubscriptionManager_UpdateToggleContractSubscription(t *testing.T) {
	m, ctrl, tgMock, dMock := newTestSubscriptionManagerWithMocks(t)
	defer ctrl.Finish()

	tgUser := m.userManager.createOrUpdate(1, "username", user.TypeTelegram)

	tgMock.EXPECT().AddOrRemoveContract(int64(1), 1).Return(true, nil)
	//goland:noinspection GoUnhandledErrorResult
	m.UpdateToggleContractSubscription(tgUser, int64(1), 1)

	discordUser := m.userManager.createOrUpdate(1, "username", user.TypeDiscord)

	dMock.EXPECT().AddOrRemoveContract(int64(1), 1).Return(true, nil)
	//goland:noinspection GoUnhandledErrorResult
	m.UpdateToggleContractSubscription(discordUser, int64(1), 1)
}

func TestSubscriptionManager_querySubscriptions(t *testing.T) {
	m := newTestSubscriptionManager(t)

	data1 := &types.ContractData{
		Address:         "0x123",
		Name:            "contract1",
		Description:     "desc1",
		ImageUrl:        "url1",
		ContractVersion: types.ContractVersionUnknown,
	}
	data2 := &types.ContractData{
		Address:         "0x456",
		Name:            "contract2",
		Description:     "desc2",
		ImageUrl:        "url2",
		ContractVersion: types.ContractVersionUnknown,
	}
	c2, _ := m.contractManager.Create(data2)
	c1, _ := m.contractManager.Create(data1)

	subscriptions := m.querySubscriptions([]int{c1.ID}, "", false)
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
	if !subscriptions[0].IsSubscribed {
		t.Error("Expected notify to be true")
	}
	if subscriptions[1].IsSubscribed {
		t.Error("Expected notify to be false")
	}

}

func TestSubscriptionManager_QuerySubscriptions_Contracts(t *testing.T) {
	m := newTestSubscriptionManager(t)

	data1 := &types.ContractData{
		Address:         "0x123",
		Name:            "contract1",
		Description:     "desc1",
		ImageUrl:        "url1",
		ContractVersion: types.ContractVersionUnknown,
	}
	data2 := &types.ContractData{
		Address:         "0x456",
		Name:            "contract2",
		Description:     "desc2",
		ImageUrl:        "url2",
		ContractVersion: types.ContractVersionUnknown,
	}
	c2, _ := m.contractManager.Create(data2)
	c1, _ := m.contractManager.Create(data1)

	m.telegramChatManager.CreateOrUpdate(1, "telegramuser", 11, "chat1", false)
	m.telegramChatManager.CreateOrUpdate(1, "telegramuser", 10, "chat2", true)
	m.discordChannelManager.CreateOrUpdate(1, "discorduser", 12, "channel1", false)

	tgUser1, _ := m.userManager.QueryById(1, user.TypeTelegram)
	m.userManager.QueryById(2, user.TypeTelegram)
	discordUser, _ := m.userManager.QueryById(1, user.TypeDiscord)

	_, err := m.UpdateToggleContractSubscription(tgUser1, int64(10), c1.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.UpdateToggleContractSubscription(discordUser, int64(12), c2.ID)
	if err != nil {
		t.Fatal(err)
	}

	response := m.QuerySubscriptions(tgUser1)
	chatRooms := response.ContractChatRooms
	if len(chatRooms) != 2 {
		t.Fatalf("Expected 2 chatRooms, got %d", len(response.ChainChatRooms))
	}
	ch1, ch2 := chatRooms[0], chatRooms[1]
	if ch1.Id != int64(11) || ch2.Id != int64(10) {
		t.Errorf("Expected 11 and 10, got %d and %d", ch1.Id, ch2.Id)
	}
	if ch1.Name != "chat1" || ch2.Name != "chat2" {
		t.Errorf("Expected chat1 and chat2, got %s and %s", ch1.Name, ch2.Name)
	}

	if len(chatRooms[0].Subscriptions) != 2 {
		t.Errorf("Expected 2 subscriptions, got %d", len(chatRooms[0].Subscriptions))
	}
	c1s1, c1s2, c2s1, c2s2 := chatRooms[0].Subscriptions[0], chatRooms[0].Subscriptions[1], chatRooms[1].Subscriptions[0], chatRooms[1].Subscriptions[1]
	if c1s1.Name != data1.Name || c2s1.Name != data1.Name || c1s2.Name != data2.Name || c2s2.Name != data2.Name {
		t.Errorf("Expected %s, %s, %s, %s, got %s, %s, %s, %s", data1.Name, data1.Name, data2.Name, data2.Name, c1s1.Name, c2s1.Name, c1s2.Name, c2s2.Name)
	}
	if c1s1.Id != int64(c1.ID) || c2s1.Id != int64(c1.ID) || c1s2.Id != int64(c2.ID) || c2s2.Id != int64(c2.ID) {
		t.Errorf("Expected %d, %d, %d, %d, got %d, %d, %d, %d", c1.ID, c1.ID, c2.ID, c2.ID, c1s1.Id, c2s1.Id, c1s2.Id, c2s2.Id)
	}
	if c1s1.IsSubscribed != false || c2s1.IsSubscribed != true || c1s2.IsSubscribed != false || c2s2.IsSubscribed != false {
		t.Errorf("Expected false, true, false, false, got %t, %t, %t, %t", c1s1.IsSubscribed, c2s1.IsSubscribed, c1s2.IsSubscribed, c2s2.IsSubscribed)
	}

	response = m.QuerySubscriptions(discordUser)
	chatRooms = response.ContractChatRooms
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
		t.Errorf("Expected 2 subscriptions, got %d", len(chatRooms[0].Subscriptions))
	}
	c1s1, c1s2 = chatRooms[0].Subscriptions[0], chatRooms[0].Subscriptions[1]
	if c1s1.Name != data1.Name || c1s2.Name != data2.Name {
		t.Errorf("Expected %s, %s, got %s, %s", data1.Name, data2.Name, c1s1.Name, c1s2.Name)
	}
	if c1s1.Id != int64(c1.ID) || c1s2.Id != int64(c2.ID) {
		t.Errorf("Expected %d, %d, got %d, %d", c1.ID, c2.ID, c1s1.Id, c1s2.Id)
	}
	if c1s1.IsSubscribed != false || c1s2.IsSubscribed != true {
		t.Errorf("Expected %v, got %v and %v", false, c1s1.IsSubscribed, c1s2.IsSubscribed)
	}
}

func TestSubscriptionManager_QuerySubscriptions_Chains(t *testing.T) {
	m := newTestSubscriptionManager(t)

	data1 := &types.Chain{
		ChainId:     "chain1",
		Name:        "chain1",
		PrettyName:  "Chain 1",
		NetworkType: "mainnet",
		Image:       "https://image1.png",
	}
	data2 := &types.Chain{
		ChainId:     "chain2",
		Name:        "chain2",
		PrettyName:  "Chain 2",
		NetworkType: "mainnet",
		Image:       "https://image2.png",
	}
	c2, _ := m.chainManager.Create(data2, data2.Image)
	c1, _ := m.chainManager.Create(data1, data1.Image)

	m.telegramChatManager.CreateOrUpdate(1, "telegramuser", 10, "chat2", true)
	m.telegramChatManager.CreateOrUpdate(1, "telegramuser", 11, "chat1", false)
	m.discordChannelManager.CreateOrUpdate(1, "discorduser", 12, "channel1", false)

	tgUser1, _ := m.userManager.QueryById(1, user.TypeTelegram)
	m.userManager.QueryById(2, user.TypeTelegram)
	discordUser, _ := m.userManager.QueryById(1, user.TypeDiscord)

	_, err := m.UpdateToggleChainSubscription(tgUser1, int64(10), c1.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.UpdateToggleChainSubscription(discordUser, int64(12), c2.ID)
	if err != nil {
		t.Fatal(err)
	}

	response := m.QuerySubscriptions(tgUser1)
	chatRooms := response.ChainChatRooms
	if len(chatRooms) != 2 {
		t.Fatalf("Expected 2 chatRooms, got %d", len(chatRooms))
	}
	ch1, ch2 := chatRooms[0], chatRooms[1]
	if ch1.Id != int64(11) || ch2.Id != int64(10) {
		t.Errorf("Expected 11 and 10, got %d and %d", ch1.Id, ch2.Id)
	}
	if ch1.Name != "chat1" || ch2.Name != "chat2" {
		t.Errorf("Expected chat1 and chat2, got %s and %s", ch1.Name, ch2.Name)
	}

	if len(chatRooms[0].Subscriptions) != 2 {
		t.Errorf("Expected 2 subscriptions, got %d", len(chatRooms[0].Subscriptions))
	}
	c1s1, c1s2, c2s1, c2s2 := chatRooms[0].Subscriptions[0], chatRooms[0].Subscriptions[1], chatRooms[1].Subscriptions[0], chatRooms[1].Subscriptions[1]
	if c1s1.Name != data1.PrettyName || c2s1.Name != data1.PrettyName || c1s2.Name != data2.PrettyName || c2s2.Name != data2.PrettyName {
		t.Errorf("Expected %s, %s, %s, %s, got %s, %s, %s, %s", data1.PrettyName, data1.PrettyName, data2.PrettyName, data2.PrettyName, c1s1.Name, c2s1.Name, c1s2.Name, c2s2.Name)
	}
	if c1s1.Id != int64(c1.ID) || c2s1.Id != int64(c1.ID) || c1s2.Id != int64(c2.ID) || c2s2.Id != int64(c2.ID) {
		t.Errorf("Expected %d or %d, got %d and %d", c1.ID, c2.ID, c1s1.Id, c1s2.Id)
	}
	if c1s1.IsSubscribed != false || c2s1.IsSubscribed != true || c1s2.IsSubscribed != false || c2s2.IsSubscribed != false {
		t.Errorf("Expected (false, true, false, false), got (%v, %v, %v, %v)", c1s1.IsSubscribed, c2s1.IsSubscribed, c1s2.IsSubscribed, c2s2.IsSubscribed)
	}

	response = m.QuerySubscriptions(discordUser)
	chatRooms = response.ChainChatRooms
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
		t.Errorf("Expected 2 subscriptions, got %d", len(chatRooms[0].Subscriptions))
	}
	c1s1, c1s2 = chatRooms[0].Subscriptions[0], chatRooms[0].Subscriptions[1]
	if c1s1.Name != data1.PrettyName || c1s2.Name != data2.PrettyName {
		t.Errorf("Expected (%s, %s), got (%s, %s)", data1.PrettyName, data2.PrettyName, c1s1.Name, c1s2.Name)
	}
	if c1s1.Id != int64(c1.ID) || c1s2.Id != int64(c2.ID) {
		t.Errorf("Expected (%d, %d), got (%d, %d)", c1.ID, c2.ID, c1s1.Id, c1s2.Id)
	}
	if c1s1.IsSubscribed != false || c1s2.IsSubscribed != true {
		t.Errorf("Expected (false, true), got (%v, %v)", c1s1.IsSubscribed, c1s2.IsSubscribed)
	}
}

func TestSubscriptionManager_QuerySubscriptions_ContractsAndChains(t *testing.T) {
	m := newTestSubscriptionManager(t)

	data1 := &types.Chain{
		ChainId:     "chain1",
		Name:        "chain1",
		PrettyName:  "Chain 1",
		NetworkType: "mainnet",
		Image:       "https://image1.png",
	}
	data2 := &types.Chain{
		ChainId:     "chain2",
		Name:        "chain2",
		PrettyName:  "Chain 2",
		NetworkType: "mainnet",
		Image:       "https://image2.png",
	}
	data3 := &types.ContractData{
		Address:         "0x123",
		Name:            "contract1",
		Description:     "desc1",
		ImageUrl:        "url1",
		ContractVersion: types.ContractVersionUnknown,
	}
	c2, _ := m.chainManager.Create(data2, data2.Image)
	c1, _ := m.chainManager.Create(data1, data1.Image)
	contract1, _ := m.contractManager.Create(data3)

	m.telegramChatManager.CreateOrUpdate(1, "telegramuser", 10, "chat2", true)
	m.telegramChatManager.CreateOrUpdate(1, "telegramuser", 11, "chat1", false)
	m.discordChannelManager.CreateOrUpdate(1, "discorduser", 12, "channel1", false)

	tgUser1, _ := m.userManager.QueryById(1, user.TypeTelegram)
	m.userManager.QueryById(2, user.TypeTelegram)
	discordUser, _ := m.userManager.QueryById(1, user.TypeDiscord)

	_, err := m.UpdateToggleChainSubscription(tgUser1, int64(10), c1.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.UpdateToggleChainSubscription(discordUser, int64(12), c2.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.UpdateToggleContractSubscription(discordUser, int64(12), contract1.ID)
	if err != nil {
		t.Fatal(err)
	}

	response := m.QuerySubscriptions(tgUser1)
	chainChatRooms := response.ChainChatRooms
	if len(chainChatRooms) != 2 {
		t.Fatalf("Expected 2 chainChatRooms, got %d", len(chainChatRooms))
	}
	contractChatRooms := response.ContractChatRooms
	if len(contractChatRooms) != 2 {
		t.Fatalf("Expected 2 contractChatRooms, got %d", len(contractChatRooms))
	}

	if len(chainChatRooms[0].Subscriptions) != 2 {
		t.Errorf("Expected 2 subscriptions, got %d", len(chainChatRooms[0].Subscriptions))
	}
	c1s1, c1s2, c2s1, c2s2 := chainChatRooms[0].Subscriptions[0], chainChatRooms[0].Subscriptions[1], chainChatRooms[1].Subscriptions[0], chainChatRooms[1].Subscriptions[1]
	if c1s1.Name != data1.PrettyName || c2s1.Name != data1.PrettyName || c1s2.Name != data2.PrettyName || c2s2.Name != data2.PrettyName {
		t.Errorf("Expected (%s, %s, %s, %s), got (%s, %s, %s, %s)", data1.PrettyName, data1.PrettyName, data2.PrettyName, data2.PrettyName, c1s1.Name, c2s1.Name, c1s2.Name, c2s2.Name)
	}
	if c1s1.Id != int64(c1.ID) || c2s1.Id != int64(c1.ID) || c1s2.Id != int64(c2.ID) || c2s2.Id != int64(c2.ID) {
		t.Errorf("Expected %d or %d, got %d and %d", c1.ID, c2.ID, c1s1.Id, c1s2.Id)
	}
	if c1s1.IsSubscribed != false || c2s1.IsSubscribed != true || c1s2.IsSubscribed != false || c2s2.IsSubscribed != false {
		t.Errorf("Expected (false, true, false, false), got (%v, %v, %v, %v)", c1s1.IsSubscribed, c2s1.IsSubscribed, c1s2.IsSubscribed, c2s2.IsSubscribed)
	}

	if len(contractChatRooms[0].Subscriptions) != 1 {
		t.Errorf("Expected 1 subscription, got %d", len(contractChatRooms[0].Subscriptions))
	}
	contr1s1 := contractChatRooms[0].Subscriptions[0]
	if contr1s1.Name != contract1.Name {
		t.Errorf("Expected %s, got %s", contract1.Name, contr1s1.Name)
	}
	if contr1s1.Id != int64(contract1.ID) {
		t.Errorf("Expected %d, got %d", contract1.ID, contr1s1.Id)
	}
	if contr1s1.IsSubscribed != false {
		t.Errorf("Expected %v, got %v", false, contr1s1.IsSubscribed)
	}

	response = m.QuerySubscriptions(discordUser)
	chainChatRooms = response.ChainChatRooms
	if len(chainChatRooms) != 1 {
		t.Fatalf("Expected 1 chainChatRooms, got %d", len(chainChatRooms))
	}
	contractChatRooms = response.ContractChatRooms
	if len(contractChatRooms) != 1 {
		t.Fatalf("Expected 1 contractChatRooms, got %d", len(contractChatRooms))
	}

	if len(chainChatRooms[0].Subscriptions) != 2 {
		t.Errorf("Expected 2 subscriptions, got %d", len(chainChatRooms[0].Subscriptions))
	}
	c1s1, c1s2 = chainChatRooms[0].Subscriptions[0], chainChatRooms[0].Subscriptions[1]
	if c1s1.Name != data1.PrettyName || c1s2.Name != data2.PrettyName {
		t.Errorf("Expected (%s, %s), got (%s, %s)", data1.PrettyName, data2.PrettyName, c1s1.Name, c1s2.Name)
	}
	if c1s1.Id != int64(c1.ID) || c1s2.Id != int64(c2.ID) {
		t.Errorf("Expected (%d, %d), got (%d, %d)", c1.ID, c2.ID, c1s1.Id, c1s2.Id)
	}
	if c1s1.IsSubscribed != false || c1s2.IsSubscribed != true {
		t.Errorf("Expected (false, true), got (%v, %v)", c1s1.IsSubscribed, c1s2.IsSubscribed)
	}

	_, err = m.UpdateToggleChainSubscription(tgUser1, int64(10), c1.ID)
	if err != nil {
		t.Fatal(err)
	}
	response = m.QuerySubscriptions(tgUser1)
	chainChatRooms = response.ChainChatRooms
	c1s1, c1s2, c2s1, c2s2 = response.ChainChatRooms[0].Subscriptions[0],
		response.ChainChatRooms[0].Subscriptions[1],
		response.ChainChatRooms[1].Subscriptions[0],
		response.ChainChatRooms[1].Subscriptions[1]
	if c1s1.IsSubscribed != false || c2s1.IsSubscribed != false || c1s2.IsSubscribed != false || c2s2.IsSubscribed != false {
		t.Errorf("Expected (false, false, false, false), got (%v, %v, %v, %v)", c1s1.IsSubscribed, c2s1.IsSubscribed, c1s2.IsSubscribed, c2s2.IsSubscribed)
	}

	_, err = m.UpdateToggleContractSubscription(discordUser, int64(12), c2.ID)
	if err != nil {
		t.Fatal(err)
	}
	response = m.QuerySubscriptions(discordUser)
	c1s1 = response.ContractChatRooms[0].Subscriptions[0]
	if c1s1.IsSubscribed != false {
		t.Errorf("Expected %v, got %v", false, c1s1.IsSubscribed)
	}
}

func TestSubscriptionManager_QuerySubscriptions_WithStats(t *testing.T) {
	m := newTestSubscriptionManager(t)

	data1 := &types.Chain{
		ChainId:     "chain1",
		Name:        "chain1",
		PrettyName:  "Chain 1",
		NetworkType: "mainnet",
		Image:       "https://image1.png",
	}
	data2 := &types.Chain{
		ChainId:     "chain2",
		Name:        "chain2",
		PrettyName:  "Chain 2",
		NetworkType: "mainnet",
		Image:       "https://image2.png",
	}
	data3 := &types.ContractData{
		Address:         "0x123",
		Name:            "contract1",
		Description:     "desc1",
		ImageUrl:        "url1",
		ContractVersion: types.ContractVersionUnknown,
	}
	chain2, _ := m.chainManager.Create(data2, data2.Image)
	chain1, _ := m.chainManager.Create(data1, data1.Image)
	contract1, _ := m.contractManager.Create(data3)

	m.telegramChatManager.CreateOrUpdate(1, "telegramuser", 10, "chat2", true)
	m.telegramChatManager.CreateOrUpdate(1, "telegramuser", 11, "chat1", false)
	m.discordChannelManager.CreateOrUpdate(1, "discorduser", 12, "channel1", false)
	m.userManager.UpdateRole("telegramuser", user.RoleAdmin)
	m.userManager.UpdateRole("discorduser", user.RoleAdmin)

	tgUser1, _ := m.userManager.QueryById(1, user.TypeTelegram)
	m.userManager.QueryById(2, user.TypeTelegram)
	discordUser, _ := m.userManager.QueryById(1, user.TypeDiscord)

	_, err := m.UpdateToggleChainSubscription(tgUser1, int64(10), chain1.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.UpdateToggleChainSubscription(discordUser, int64(12), chain2.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.UpdateToggleContractSubscription(tgUser1, int64(10), contract1.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.UpdateToggleContractSubscription(discordUser, int64(12), contract1.ID)
	if err != nil {
		t.Fatal(err)
	}

	for i, u := range []*ent.User{tgUser1, discordUser} {
		response := m.QuerySubscriptions(u)
		if i == 0 {
			if len(response.ChainChatRooms) != 2 {
				t.Fatalf("Expected 2 chainChatRooms, got %d", len(response.ChainChatRooms))
			}
			if len(response.ContractChatRooms) != 2 {
				t.Fatalf("Expected 2 contractChatRooms, got %d", len(response.ContractChatRooms))
			}
		} else {
			if len(response.ChainChatRooms) != 1 {
				t.Fatalf("Expected 1 chainChatRooms, got %d", len(response.ChainChatRooms))
			}
			if len(response.ContractChatRooms) != 1 {
				t.Fatalf("Expected 1 contractChatRooms, got %d", len(response.ContractChatRooms))
			}
		}
		for _, c := range response.ChainChatRooms {
			c1s1, c1s2 := c.Subscriptions[0].Stats, c.Subscriptions[1].Stats
			if c1s1.Total != 1 || c1s1.Telegram != 1 || c1s1.Discord != 0 {
				t.Errorf("Expected (1, 1, 0), got (%d, %d, %d)", c1s1.Total, c1s1.Telegram, c1s1.Discord)
			}
			if c1s2.Total != 1 || c1s2.Telegram != 0 || c1s2.Discord != 1 {
				t.Errorf("Expected (1, 0, 1), got (%d, %d, %d)", c1s2.Total, c1s2.Telegram, c1s2.Discord)
			}
		}
		for _, c := range response.ContractChatRooms {
			c1s1 := c.Subscriptions[0].Stats
			if c1s1.Total != 2 || c1s1.Telegram != 1 || c1s1.Discord != 1 {
				t.Errorf("Expected (2, 1, 1), got (%d, %d, %d)", c1s1.Total, c1s1.Telegram, c1s1.Discord)
			}
		}
	}
}

func TestSubscriptionManager_QuerySubscriptions_EnabledChains(t *testing.T) {
	m := newTestSubscriptionManager(t)

	data1 := &types.Chain{
		ChainId:     "chain1",
		Name:        "chain1",
		PrettyName:  "Chain 1",
		NetworkType: "mainnet",
		Image:       "https://image1.png",
	}
	data2 := &types.Chain{
		ChainId:     "chain2",
		Name:        "chain2",
		PrettyName:  "Chain 2",
		NetworkType: "mainnet",
		Image:       "https://image2.png",
	}
	m.chainManager.Create(data2, data2.Image)
	chain1, _ := m.chainManager.Create(data1, data1.Image)

	m.telegramChatManager.CreateOrUpdate(1, "telegramuser", 10, "chat2", true)
	m.discordChannelManager.CreateOrUpdate(1, "discorduser", 12, "channel1", false)

	tgUser1, _ := m.userManager.QueryById(1, user.TypeTelegram)
	discordUser, _ := m.userManager.QueryById(1, user.TypeDiscord)

	for _, u := range []*ent.User{tgUser1, discordUser} {
		response := m.QuerySubscriptions(u)
		if len(response.ChainChatRooms[0].Subscriptions) != 2 {
			t.Fatalf("Expected 2 subscriptions, got %d", len(response.ChainChatRooms[0].Subscriptions))
		}
		for _, s := range response.ChainChatRooms[0].Subscriptions {
			if !s.IsEnabled {
				t.Errorf("Expected enabled chain subscription, got disabled")
			}
		}
	}

	m.chainManager.UpdateSetEnabled(chain1.ID, false)

	for _, u := range []*ent.User{tgUser1, discordUser} {
		response := m.QuerySubscriptions(u)
		if len(response.ChainChatRooms[0].Subscriptions) != 1 {
			t.Fatalf("Expected 1 subscription, got %d", len(response.ChainChatRooms[0].Subscriptions))
		}
		for _, s := range response.ChainChatRooms[0].Subscriptions {
			if !s.IsEnabled {
				t.Errorf("Expected enabled chain subscription, got disabled")
			}
		}
	}

	tgUser1, _ = m.userManager.UpdateRole("telegramuser", user.RoleAdmin)
	discordUser, _ = m.userManager.UpdateRole("discorduser", user.RoleAdmin)

	for _, u := range []*ent.User{tgUser1, discordUser} {
		response := m.QuerySubscriptions(u)
		if len(response.ChainChatRooms[0].Subscriptions) != 2 {
			t.Fatalf("Expected 2 subscriptions, got %d", len(response.ChainChatRooms[0].Subscriptions))
		}
		if response.ChainChatRooms[0].Subscriptions[0].IsEnabled {
			t.Errorf("Expected disabled chain subscription, got enabled")
		}
		if !response.ChainChatRooms[0].Subscriptions[1].IsEnabled {
			t.Errorf("Expected enabled chain subscription, got disabled")
		}
	}
}
