package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/log"
)

type UserManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewUserManager(client *ent.Client, ctx context.Context) *UserManager {
	return &UserManager{client: client, ctx: ctx}
}

func (m *UserManager) Get(userId int64, userType user.Type) (*ent.User, error) {
	return m.client.User.
		Query().
		Where(user.And(
			user.UserIDEQ(userId),
			user.TypeEQ(userType),
		)).
		Only(m.ctx)
}

func (m *UserManager) GetAdmins() ([]*ent.User, error) {
	return m.client.User.
		Query().
		Where(user.RoleEQ(user.RoleAdmin)).
		All(m.ctx)
}

func (m *UserManager) SetName(entUser *ent.User, name string) (*ent.User, error) {
	return entUser.Update().
		SetName(name).
		Save(m.ctx)
}

func (m *UserManager) SetRole(name string, role user.Role) (*ent.User, error) {
	entUser, err := m.client.User.
		Query().
		Where(user.NameEQ(name)).
		Only(m.ctx)
	if err != nil {
		return nil, err
	}
	return entUser.
		Update().
		SetRole(role).
		Save(m.ctx)
}

func (m *UserManager) createOrUpdateUser(userId int64, userName string, userType user.Type) *ent.User {
	entUser, err := m.client.User.
		Query().
		Where(user.And(
			user.UserIDEQ(userId),
			user.TypeEQ(userType),
		)).
		Only(m.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			entUser, err = m.client.User.
				Create().
				SetUserID(userId).
				SetName(userName).
				SetType(userType).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Could not create user: %v", err)
			}
		} else {
			log.Sugar.Panicf("Could not find user: %v", err)
		}
	} else if entUser.Name != userName {
		entUser, err = m.client.User.
			UpdateOne(entUser).
			SetName(userName).
			Save(m.ctx)
		if err != nil {
			log.Sugar.Panicf("Could not update user: %v", err)
		}
	}
	return entUser
}

// deletes a user if they have no more telegram chats or discord channels
func (m *UserManager) deleteIfUnused(entUser *ent.User) {
	cnt := entUser.QueryTelegramChats().CountX(m.ctx)
	if cnt == 0 {
		cnt = entUser.QueryDiscordChannels().CountX(m.ctx)
		if cnt == 0 {
			log.Sugar.Debugf("Deleting user %s (%d)", entUser.Name, entUser.UserID)
			err := m.client.User.DeleteOne(entUser).Exec(m.ctx)
			if err != nil {
				log.Sugar.Errorf("Could not delete user: %v", err)
			}
		}
	}
}
