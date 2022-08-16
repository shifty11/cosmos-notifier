package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/user"
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

func (m *UserManager) SetName(entUser *ent.User, name string) (*ent.User, error) {
	return entUser.Update().
		SetName(name).
		Save(m.ctx)
}
