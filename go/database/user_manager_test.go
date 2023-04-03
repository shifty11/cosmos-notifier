package database

import (
	"context"
	"fmt"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"testing"
)

func newTestUserManager(t *testing.T) *UserManager {
	manager := NewUserManager(testClient(t), context.Background())
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func addUsers(m *UserManager, number int, userType user.Type) []*ent.User {
	lastUserId := int64(1)
	userDto, err := m.client.User.
		Query().
		Order(ent.Desc(user.FieldUserID)).
		First(m.ctx)
	if !ent.IsNotFound(err) {
		lastUserId = userDto.UserID
	}
	users := make([]*ent.User, number)
	for i := 0; i < number; i++ {
		lastUserId++
		u := m.createOrUpdate(lastUserId, fmt.Sprintf("userDto %d", lastUserId), userType)
		users[i] = u
	}
	return users
}

func TestUserManager_CreateOrUpdate(t *testing.T) {
	m := newTestUserManager(t)
	u := m.createOrUpdate(1, "username", user.TypeDiscord)
	if u.UserID != 1 {
		t.Fatalf("Expected 1, got %d", u.UserID)
	}
	if u.Name != "username" {
		t.Fatalf("Expected username, got %s", u.Name)
	}
	if u.Type != user.TypeDiscord {
		t.Fatalf("Expected discord, got %s", u.Type)
	}

	u = m.createOrUpdate(1, "updated", user.TypeDiscord)
	if u.UserID != 1 {
		t.Fatalf("Expected 1, got %d", u.UserID)
	}
	if u.Name != "updated" {
		t.Fatalf("Expected updated, got %s", u.Name)
	}
}

func TestUserManager_DeleteIfUnused(t *testing.T) {
	m := newTestUserManager(t)
	u := m.createOrUpdate(1, "username", user.TypeDiscord)

	dc := m.client.DiscordChannel.Create().AddUsers(u).SetChannelID(1).SetName("test").SetIsGroup(false).SaveX(m.ctx)
	tg := m.client.TelegramChat.Create().AddUsers(u).SetChatID(1).SetName("test").SetIsGroup(false).SaveX(m.ctx)
	m.deleteIfUnused(u)
	if m.client.User.Query().CountX(m.ctx) != 1 {
		t.Fatal("Expected 1, got 0")
	}

	m.client.DiscordChannel.DeleteOne(dc).ExecX(m.ctx)
	m.deleteIfUnused(u)
	if m.client.User.Query().CountX(m.ctx) != 1 {
		t.Fatal("Expected 1, got 0")
	}

	m.client.TelegramChat.DeleteOne(tg).ExecX(m.ctx)
	m.deleteIfUnused(u)
	if m.client.User.Query().CountX(m.ctx) != 0 {
		t.Fatal("Expected 0, got 1")
	}
}

func TestUserManager_UpdateRole(t *testing.T) {
	m := newTestUserManager(t)
	_, err := m.UpdateRole("non existent", user.RoleAdmin)
	if !ent.IsNotFound(err) {
		t.Fatalf("Expected not found error, got %s", err)
	}

	u := m.createOrUpdate(1, "username", user.TypeDiscord)
	if u.Role != user.RoleUser {
		t.Fatalf("Expected user, got %s", u.Role)
	}

	u, err = m.UpdateRole(u.Name, user.RoleAdmin)
	if err != nil {
		t.Fatalf("Expected nil, got %s", err)
	}
	if u.Role != user.RoleAdmin {
		t.Fatalf("Expected admin, got %s", u.Role)
	}
}
