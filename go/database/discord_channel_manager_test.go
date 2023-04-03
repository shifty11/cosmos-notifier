package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/types"
	"testing"
)

func newTestDiscordChannelManager(t *testing.T) *DiscordChannelManager {
	manager := NewDiscordChannelManager(testClient(t), context.Background(), newTestChainManager(t), newTestContractManager(t), newTestUserManager(t))
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func addDiscordChannels(m *DiscordChannelManager, users []*ent.User) []*ent.DiscordChannel {
	var channels []*ent.DiscordChannel
	var nextChannelID = int64(m.client.DiscordChannel.Query().CountX(m.ctx) + 1)
	for _, userDto := range users {
		if userDto.Type == user.TypeDiscord {
			c := m.client.DiscordChannel.
				Create().
				SetChannelID(nextChannelID).
				SetName("channel-1").
				SetIsGroup(false).
				AddUsers(userDto).
				SaveX(m.ctx)
			channels = append(channels, c)
			nextChannelID++
			c = m.client.DiscordChannel.
				Create().
				SetChannelID(nextChannelID).
				SetName("channel-2").
				SetIsGroup(true).
				AddUsers(userDto).
				SaveX(m.ctx)
			channels = append(channels, c)
			nextChannelID++
		}
	}
	return channels
}

func TestDiscordChannelManager_UpdateAddOrRemoveChain(t *testing.T) {
	m := newTestDiscordChannelManager(t)
	_, err := m.UpdateAddOrRemoveChain(1, 1)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	u := m.userManager.createOrUpdate(1, "username", user.TypeDiscord)

	dc := m.client.DiscordChannel.
		Create().
		AddUsers(u).
		SetChannelID(1).
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

	c := m.chainManager.Create(data, data.Image)
	added, err := m.UpdateAddOrRemoveChain(dc.ChannelID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if !added {
		t.Fatal("Expected true, got false")
	}

	added, err = m.UpdateAddOrRemoveChain(dc.ChannelID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if added {
		t.Fatal("Expected false, got true")
	}
}

func TestDiscordChannelManager_UpdateAddOrRemoveContract(t *testing.T) {
	m := newTestDiscordChannelManager(t)
	_, err := m.UpdateAddOrRemoveContract(1, 1)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	u := m.userManager.createOrUpdate(1, "username", user.TypeDiscord)

	dc := m.client.DiscordChannel.
		Create().
		AddUsers(u).
		SetChannelID(1).
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
	added, err := m.UpdateAddOrRemoveContract(dc.ChannelID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if !added {
		t.Fatal("Expected true, got false")
	}

	added, err = m.UpdateAddOrRemoveContract(dc.ChannelID, c.ID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if added {
		t.Fatal("Expected false, got true")
	}
}

func TestDiscordChannelManager_CreateOrUpdate(t *testing.T) {
	m := newTestDiscordChannelManager(t)

	dc, created := m.CreateOrUpdate(1, "username", 1, "channelname", true)
	if !created {
		t.Fatal("Expected true, got false")
	}
	if dc.ChannelID != 1 {
		t.Fatalf("Expected 1, got %d", dc.ChannelID)
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

func TestDiscordChannelManager_QuerySubscribedIds(t *testing.T) {
	m := newTestDiscordChannelManager(t)

	data := &types.ContractData{
		Address:         "0x123",
		Name:            "test",
		Description:     "description",
		ImageUrl:        "https://image.com",
		ContractVersion: types.ContractVersionUnknown,
	}

	c1, _ := m.contractManager.Create(data)

	u := m.userManager.createOrUpdate(1, "username", user.TypeDiscord)

	dc := m.client.DiscordChannel.
		Create().
		SetChannelID(1).
		SetName("test").
		SetIsGroup(false).
		AddUsers(u).
		SaveX(m.ctx)

	ids := m.QuerySubscribedIds(c1.QueryDiscordChannels())
	if len(ids) != 0 {
		t.Fatalf("Expected 0, got %d", len(ids))
	}

	_, _ = m.UpdateAddOrRemoveContract(dc.ChannelID, c1.ID)
	ids = m.QuerySubscribedIds(c1.QueryDiscordChannels())
	if len(ids) != 1 {
		t.Fatalf("Expected 1, got %d", len(ids))
	}
	if ids[0].ChannelId != 1 {
		t.Fatalf("Expected %d, got %d", 1, ids[0].ChannelId)
	}
	if ids[0].Name != "test" {
		t.Fatalf("Expected test, got %s", ids[0].Name)
	}

	dc = m.client.DiscordChannel.
		Create().
		SetChannelID(2).
		SetName("test2").
		SetIsGroup(false).
		AddUsers(u).
		SaveX(m.ctx)
	_, _ = m.UpdateAddOrRemoveContract(dc.ChannelID, c1.ID)
	ids = m.QuerySubscribedIds(c1.QueryDiscordChannels())
	if len(ids) != 2 {
		t.Fatalf("Expected 2, got %d", len(ids))
	}
	if ids[0].ChannelId != 1 {
		t.Fatalf("Expected %d, got %d", 1, ids[0].ChannelId)
	}
	if ids[0].Name != "test" {
		t.Fatalf("Expected test, got %s", ids[0].Name)
	}
	if ids[1].ChannelId != 2 {
		t.Fatalf("Expected %d, got %d", 2, ids[1].ChannelId)
	}
	if ids[1].Name != "test2" {
		t.Fatalf("Expected test2, got %s", ids[1].Name)
	}

}

func TestDiscordChannelManager_Delete(t *testing.T) {
	m := newTestDiscordChannelManager(t)
	u1 := m.userManager.createOrUpdate(1, "username", user.TypeDiscord)
	u2 := m.userManager.createOrUpdate(2, "username", user.TypeDiscord)

	err := m.Delete(1, 100)
	if !ent.IsNotFound(err) {
		t.Fatal("Expected not found error, got nil")
	}

	dc := m.client.DiscordChannel.
		Create().
		SetChannelID(1).
		SetName("test").
		SetIsGroup(false).
		AddUsers(u1, u2).
		SaveX(m.ctx)

	err = m.Delete(100, dc.ChannelID)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	err = m.Delete(u1.UserID, dc.ChannelID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	cnt := u2.QueryDiscordChannels().CountX(m.ctx)
	if cnt != 1 {
		t.Fatalf("Expected 1, got %d", cnt)
	}
	cnt = m.client.DiscordChannel.Query().CountX(m.ctx)
	if cnt != 1 {
		t.Fatalf("Expected 1, got %d", cnt)
	}
	_, err = m.client.User.Get(m.ctx, u1.ID)
	if !ent.IsNotFound(err) {
		t.Fatalf("Expected not found error, got %s", err)
	}

	err = m.Delete(u2.UserID, dc.ChannelID)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	cnt = m.client.DiscordChannel.Query().CountX(m.ctx)
	if cnt != 0 {
		t.Fatalf("Expected 0, got %d", cnt)
	}
	cnt = m.client.User.Query().CountX(m.ctx)
	if cnt != 0 {
		t.Fatalf("Expected 0, got %d", cnt)
	}

}

func TestDiscordChannelManager_DeleteMultiple(t *testing.T) {
	m := newTestDiscordChannelManager(t)
	u1 := m.userManager.createOrUpdate(1, "username", user.TypeDiscord)
	u2 := m.userManager.createOrUpdate(2, "username", user.TypeDiscord)

	dc := m.client.DiscordChannel.
		Create().
		SetChannelID(1).
		SetName("test").
		SetIsGroup(false).
		AddUsers(u1, u2).
		SaveX(m.ctx)

	m.DeleteMultiple([]int64{dc.ChannelID})

	cnt := m.client.DiscordChannel.Query().CountX(m.ctx)
	if cnt != 0 {
		t.Fatalf("Expected 0, got %d", cnt)
	}
	cnt = m.client.User.Query().CountX(m.ctx)
	if cnt != 0 {
		t.Fatalf("Expected 0, got %d", cnt)
	}

}

func TestDiscordChannelManager_DeleteMultiple_KeepOneUser(t *testing.T) {
	m := newTestDiscordChannelManager(t)
	u1 := m.userManager.createOrUpdate(1, "username", user.TypeDiscord)
	u2 := m.userManager.createOrUpdate(2, "username", user.TypeDiscord)

	dc := m.client.DiscordChannel.
		Create().
		SetChannelID(1).
		SetName("test").
		SetIsGroup(false).
		AddUsers(u1, u2).
		SaveX(m.ctx)
	m.client.DiscordChannel.
		Create().
		SetChannelID(2).
		SetName("test").
		SetIsGroup(false).
		AddUsers(u2).
		SaveX(m.ctx)

	m.DeleteMultiple([]int64{dc.ChannelID})

	cnt := m.client.DiscordChannel.Query().CountX(m.ctx)
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
