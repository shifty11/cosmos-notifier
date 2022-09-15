package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"testing"
)

func newTestUserManager(t *testing.T) *UserManager {
	return NewUserManager(testClient(t), context.Background())
}

func TestUserManager_CreateOrUpdateUser(t *testing.T) {
	m := newTestUserManager(t)
	u := m.createOrUpdateUser(1, "username", user.TypeDiscord)
	if u.UserID != 1 {
		t.Fatalf("Expected 1, got %d", u.UserID)
	}
	if u.Name != "username" {
		t.Fatalf("Expected username, got %s", u.Name)
	}
	if u.Type != user.TypeDiscord {
		t.Fatalf("Expected discord, got %s", u.Type)
	}

	u = m.createOrUpdateUser(1, "updated", user.TypeDiscord)
	if u.UserID != 1 {
		t.Fatalf("Expected 1, got %d", u.UserID)
	}
	if u.Name != "updated" {
		t.Fatalf("Expected updated, got %s", u.Name)
	}
}

func TestUserManager_DeleteIfUnused(t *testing.T) {
	m := newTestUserManager(t)
	u := m.createOrUpdateUser(1, "username", user.TypeDiscord)

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
