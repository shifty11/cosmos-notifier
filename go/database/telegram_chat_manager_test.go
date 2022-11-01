package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/types"
	"testing"
)

func newTestTelegramChatManager(t *testing.T) *TelegramChatManager {
	manager := NewTelegramChatManager(testClient(t), context.Background(), newTestChainManager(t), newTestContractManager(t), newTestUserManager(t))
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func TestTelegramChatManager_AddOrRemoveChain(t *testing.T) {
	m := newTestTelegramChatManager(t)
	_, err := m.AddOrRemoveChain(1, 1)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	u := m.userManager.createOrUpdateUser(1, "username", user.TypeDiscord)

	dc := m.client.TelegramChat.
		Create().
		AddUsers(u).
		SetChatID(1).
		SetName("test").
		SetIsGroup(false).
		SaveX(m.ctx)

	_, err = m.AddOrRemoveChain(1, 1)
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

	c := m.chainManager.Create(data, data.Image)
	hasChain, err := m.AddOrRemoveChain(dc.ChatID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if !hasChain {
		t.Fatal("Expected true, got false")
	}

	hasChain, err = m.AddOrRemoveChain(dc.ChatID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if hasChain {
		t.Fatal("Expected false, got true")
	}
}

func TestTelegramChatManager_AddOrRemoveContract(t *testing.T) {
	m := newTestTelegramChatManager(t)
	_, err := m.AddOrRemoveContract(1, 1)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	u := m.userManager.createOrUpdateUser(1, "username", user.TypeDiscord)

	dc := m.client.TelegramChat.
		Create().
		AddUsers(u).
		SetChatID(1).
		SetName("test").
		SetIsGroup(false).
		SaveX(m.ctx)

	_, err = m.AddOrRemoveContract(1, 1)
	if !ent.IsNotFound(err) {
		t.Fatalf("Expected not found error, got %s", err)
	}

	data := &types.ContractData{
		Address:     "0x123",
		Name:        "test",
		Description: "description",
		ImageUrl:    "https://image.com",
	}

	c, _ := m.contractManager.Create(data)
	hasContract, err := m.AddOrRemoveContract(dc.ChatID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if !hasContract {
		t.Fatal("Expected true, got false")
	}

	hasContract, err = m.AddOrRemoveContract(dc.ChatID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if hasContract {
		t.Fatal("Expected false, got true")
	}
}

func TestTelegramChatManager_CreateOrUpdateChat(t *testing.T) {
	m := newTestTelegramChatManager(t)

	dc, created := m.CreateOrUpdateChat(1, "username", 1, "channelname", true)
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

	dc, created = m.CreateOrUpdateChat(1, "updated", 1, "updated", true)
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

	dc, created = m.CreateOrUpdateChat(2, "newuser", 1, "updated", true)
	users := dc.QueryUsers().AllX(m.ctx)
	if len(users) != 2 {
		t.Fatalf("Expected 2, got %d", len(users))
	}
	if created {
		t.Fatal("Expected false, got true")
	}

}

func TestTelegramChatManager_GetSubscribedIds(t *testing.T) {
	m := newTestTelegramChatManager(t)

	data := &types.ContractData{
		Address:     "0x123",
		Name:        "test",
		Description: "description",
		ImageUrl:    "https://image.com",
	}

	c1, _ := m.contractManager.Create(data)

	u := m.userManager.createOrUpdateUser(1, "username", user.TypeTelegram)

	tg := m.client.TelegramChat.
		Create().
		SetChatID(1).
		SetName("test").
		SetIsGroup(false).
		AddUsers(u).
		SaveX(m.ctx)

	ids := m.GetSubscribedIds(c1.QueryTelegramChats())
	if len(ids) != 0 {
		t.Fatalf("Expected 0, got %d", len(ids))
	}

	_, _ = m.AddOrRemoveContract(tg.ChatID, c1.ID)
	ids = m.GetSubscribedIds(c1.QueryTelegramChats())
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
	_, _ = m.AddOrRemoveContract(tg.ChatID, c1.ID)
	ids = m.GetSubscribedIds(c1.QueryTelegramChats())
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
	u1 := m.userManager.createOrUpdateUser(1, "username", user.TypeTelegram)
	u2 := m.userManager.createOrUpdateUser(2, "username", user.TypeTelegram)

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
	u1 := m.userManager.createOrUpdateUser(1, "username", user.TypeTelegram)
	u2 := m.userManager.createOrUpdateUser(2, "username", user.TypeTelegram)

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
