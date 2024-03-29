package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/types"
	"testing"
)

func newTestTelegramChatManager(t *testing.T) *TelegramChatManager {
	manager := NewTelegramChatManager(testClient(t), context.Background(), newTestChainManager(t), newTestContractManager(t), newTestUserManager(t))
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func addTelegramChats(m *TelegramChatManager, users []*ent.User) []*ent.TelegramChat {
	var chats []*ent.TelegramChat
	var nextChatID = int64(m.client.TelegramChat.Query().CountX(m.ctx) + 1)
	for _, userDto := range users {
		if userDto.Type == user.TypeTelegram {
			c := m.client.TelegramChat.
				Create().
				SetChatID(nextChatID).
				SetName("channel-1").
				SetIsGroup(false).
				AddUsers(userDto).
				SaveX(m.ctx)
			chats = append(chats, c)
			nextChatID++
			c = m.client.TelegramChat.
				Create().
				SetChatID(nextChatID).
				SetName("channel-2").
				SetIsGroup(true).
				AddUsers(userDto).
				SaveX(m.ctx)
			chats = append(chats, c)
			nextChatID++
		}
	}
	return chats
}

func TestTelegramChatManager_UpdateAddOrRemoveChain(t *testing.T) {
	m := newTestTelegramChatManager(t)
	_, err := m.UpdateAddOrRemoveChain(1, 1)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	u := m.userManager.createOrUpdate(1, "username", user.TypeDiscord)

	dc := m.client.TelegramChat.
		Create().
		AddUsers(u).
		SetChatID(1).
		SetName("test").
		SetIsGroup(false).
		SaveX(m.ctx)

	_, err = m.UpdateAddOrRemoveChain(1, 1)
	if !ent.IsNotFound(err) {
		t.Fatalf("Expected not found error, got %s", err)
	}

	data := &types.Chain{
		ChainId:     "chainid-1",
		Name:        "chain-1",
		PrettyName:  "Chain 1",
		NetworkType: "mainnet",
		Image:       "https://image.com",
	}

	c, _ := m.chainManager.Create(data, data.Image)
	hasChain, err := m.UpdateAddOrRemoveChain(dc.ChatID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if !hasChain {
		t.Fatal("Expected true, got false")
	}

	hasChain, err = m.UpdateAddOrRemoveChain(dc.ChatID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if hasChain {
		t.Fatal("Expected false, got true")
	}
}

func TestTelegramChatManager_UpdateAddOrRemoveContract(t *testing.T) {
	m := newTestTelegramChatManager(t)
	_, err := m.UpdateAddOrRemoveContract(1, 1)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	u := m.userManager.createOrUpdate(1, "username", user.TypeDiscord)

	dc := m.client.TelegramChat.
		Create().
		AddUsers(u).
		SetChatID(1).
		SetName("test").
		SetIsGroup(false).
		SaveX(m.ctx)

	_, err = m.UpdateAddOrRemoveContract(1, 1)
	if !ent.IsNotFound(err) {
		t.Fatalf("Expected not found error, got %s", err)
	}

	data := &types.ContractData{
		Address:         "0x123",
		Name:            "test",
		Description:     "description",
		ImageUrl:        "https://image.com",
		ContractVersion: types.ContractVersionUnknown,
	}

	c, _ := m.contractManager.Create(data)
	hasContract, err := m.UpdateAddOrRemoveContract(dc.ChatID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if !hasContract {
		t.Fatal("Expected true, got false")
	}

	hasContract, err = m.UpdateAddOrRemoveContract(dc.ChatID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if hasContract {
		t.Fatal("Expected false, got true")
	}
}

func TestTelegramChatManager_CreateOrUpdate(t *testing.T) {
	m := newTestTelegramChatManager(t)

	dc, created := m.CreateOrUpdate(1, "username", 1, "channelname", true)
	if !created {
		t.Fatal("Expected true, got false")
	}
	if dc.ChatID != 1 {
		t.Fatalf("Expected 1, got %d", dc.ChatID)
	}
	if dc.Name != "channelname" {
		t.Fatalf("Expected channelname, got %s", dc.Name)
	}
	if dc.IsGroup != true {
		t.Fatal("Expected true, got false")
	}
	_, err := dc.QueryUsers().Only(m.ctx)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}

	dc, created = m.CreateOrUpdate(1, "updated", 1, "updated", true)
	u, err := dc.QueryUsers().Only(m.ctx)
	if u.Name != "updated" {
		t.Fatalf("Expected updated, got %s", dc.Edges.Users[0].Name)
	}
	if created {
		t.Fatal("Expected false, got true")
	}
	if dc.Name != "updated" {
		t.Fatalf("Expected updated, got %s", dc.Name)
	}

	dc, created = m.CreateOrUpdate(2, "newuser", 1, "updated", true)
	users := dc.QueryUsers().AllX(m.ctx)
	if len(users) != 2 {
		t.Fatalf("Expected 2, got %d", len(users))
	}
	if created {
		t.Fatal("Expected false, got true")
	}

}

func TestTelegramChatManager_QuerySubscribedIds(t *testing.T) {
	m := newTestTelegramChatManager(t)

	data := &types.ContractData{
		Address:         "0x123",
		Name:            "test",
		Description:     "description",
		ImageUrl:        "https://image.com",
		ContractVersion: types.ContractVersionUnknown,
	}

	c1, _ := m.contractManager.Create(data)

	u := m.userManager.createOrUpdate(1, "username", user.TypeTelegram)

	tg := m.client.TelegramChat.
		Create().
		SetChatID(1).
		SetName("test").
		SetIsGroup(false).
		AddUsers(u).
		SaveX(m.ctx)

	ids := m.QuerySubscribedIds(c1.QueryTelegramChats())
	if len(ids) != 0 {
		t.Fatalf("Expected 0, got %d", len(ids))
	}

	_, _ = m.UpdateAddOrRemoveContract(tg.ChatID, c1.ID)
	ids = m.QuerySubscribedIds(c1.QueryTelegramChats())
	if len(ids) != 1 {
		t.Fatalf("Expected 1, got %d", len(ids))
	}
	if ids[0].ChatId != 1 {
		t.Fatalf("Expected %d, got %d", 1, ids[0].ChatId)
	}
	if ids[0].Name != "test" {
		t.Fatalf("Expected test, got %s", ids[0].Name)
	}

	tg = m.client.TelegramChat.
		Create().
		SetChatID(2).
		SetName("test2").
		SetIsGroup(false).
		AddUsers(u).
		SaveX(m.ctx)
	_, _ = m.UpdateAddOrRemoveContract(tg.ChatID, c1.ID)
	ids = m.QuerySubscribedIds(c1.QueryTelegramChats())
	if len(ids) != 2 {
		t.Fatalf("Expected 2, got %d", len(ids))
	}
	if ids[0].ChatId != 1 {
		t.Fatalf("Expected %d, got %d", 1, ids[0].ChatId)
	}
	if ids[0].Name != "test" {
		t.Fatalf("Expected test, got %s", ids[0].Name)
	}
	if ids[1].ChatId != 2 {
		t.Fatalf("Expected %d, got %d", 2, ids[1].ChatId)
	}
	if ids[1].Name != "test2" {
		t.Fatalf("Expected test2, got %s", ids[1].Name)
	}

}

func TestTelegramChatManager_Delete(t *testing.T) {
	m := newTestTelegramChatManager(t)
	u1 := m.userManager.createOrUpdate(1, "username", user.TypeTelegram)
	u2 := m.userManager.createOrUpdate(2, "username", user.TypeTelegram)

	err := m.Delete(1, 100)
	if !ent.IsNotFound(err) {
		t.Fatal("Expected not found error, got nil")
	}

	dc := m.client.TelegramChat.
		Create().
		SetChatID(1).
		SetName("test").
		SetIsGroup(false).
		AddUsers(u1, u2).
		SaveX(m.ctx)

	err = m.Delete(100, dc.ChatID)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	err = m.Delete(u1.UserID, dc.ChatID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	cnt := u2.QueryTelegramChats().CountX(m.ctx)
	if cnt != 1 {
		t.Fatalf("Expected 1, got %d", cnt)
	}
	cnt = m.client.TelegramChat.Query().CountX(m.ctx)
	if cnt != 1 {
		t.Fatalf("Expected 1, got %d", cnt)
	}
	_, err = m.client.User.Get(m.ctx, u1.ID)
	if !ent.IsNotFound(err) {
		t.Fatalf("Expected not found error, got %s", err)
	}

	err = m.Delete(u2.UserID, dc.ChatID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	cnt = m.client.TelegramChat.Query().CountX(m.ctx)
	if cnt != 0 {
		t.Fatalf("Expected 0, got %d", cnt)
	}
	cnt = m.client.User.Query().CountX(m.ctx)
	if cnt != 0 {
		t.Fatalf("Expected 0, got %d", cnt)
	}

}

func TestTelegramChatManager_DeleteMultiple(t *testing.T) {
	m := newTestTelegramChatManager(t)
	u1 := m.userManager.createOrUpdate(1, "username", user.TypeTelegram)
	u2 := m.userManager.createOrUpdate(2, "username", user.TypeTelegram)

	dc := m.client.TelegramChat.
		Create().
		SetChatID(1).
		SetName("test").
		SetIsGroup(false).
		AddUsers(u1, u2).
		SaveX(m.ctx)

	m.DeleteMultiple([]int64{dc.ChatID})

	cnt := m.client.TelegramChat.Query().CountX(m.ctx)
	if cnt != 0 {
		t.Fatalf("Expected 0, got %d", cnt)
	}
	cnt = m.client.User.Query().CountX(m.ctx)
	if cnt != 0 {
		t.Fatalf("Expected 0, got %d", cnt)
	}
}

func TestTelegramChatManager_DeleteMultiple_KeepOneUser(t *testing.T) {
	m := newTestTelegramChatManager(t)
	u1 := m.userManager.createOrUpdate(1, "username", user.TypeTelegram)
	u2 := m.userManager.createOrUpdate(2, "username", user.TypeTelegram)

	dc := m.client.TelegramChat.
		Create().
		SetChatID(1).
		SetName("test").
		SetIsGroup(false).
		AddUsers(u1, u2).
		SaveX(m.ctx)
	m.client.TelegramChat.
		Create().
		SetChatID(2).
		SetName("test").
		SetIsGroup(false).
		AddUsers(u2).
		SaveX(m.ctx)

	m.DeleteMultiple([]int64{dc.ChatID})

	cnt := m.client.TelegramChat.Query().CountX(m.ctx)
	if cnt != 1 {
		t.Fatalf("Expected 1, got %d", cnt)
	}
	cnt = m.client.User.Query().CountX(m.ctx)
	if cnt != 1 {
		t.Fatalf("Expected 1, got %d", cnt)
	}
	u := m.client.User.Query().OnlyX(m.ctx)
	if u.ID != u2.ID {
		t.Fatalf("Expected %d, got %d", u2.ID, u.ID)
	}
}
