package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/ent/telegramchat"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/log"
)

type TelegramChatManager struct {
	client          *ent.Client
	ctx             context.Context
	contractManager *ContractManager
}

func NewTelegramChatManager(client *ent.Client, ctx context.Context, contractManager *ContractManager) *TelegramChatManager {
	return &TelegramChatManager{client: client, ctx: ctx, contractManager: contractManager}
}

func (m *TelegramChatManager) AddOrRemoveChain(tgChatId int64, contractId int) (bool, error) {
	tgChat, err := m.client.TelegramChat.
		Query().
		Where(telegramchat.ChatIDEQ(tgChatId)).
		First(m.ctx)
	if err != nil {
		return false, err
	}

	dbContract, err := m.contractManager.Get(contractId)
	if err != nil {
		return false, err
	}

	exists, err := tgChat.
		QueryContracts().
		Where(contract.IDEQ(dbContract.ID)).
		Exist(m.ctx)
	if err != nil {
		return false, err
	}
	if exists {
		_, err := tgChat.
			Update().
			RemoveContractIDs(dbContract.ID).
			Save(m.ctx)
		if err != nil {
			return false, err
		}
	} else {
		_, err := tgChat.
			Update().
			AddContractIDs(dbContract.ID).
			Save(m.ctx)
		if err != nil {
			return false, err
		}
	}
	return !exists, nil
}

func (m *TelegramChatManager) updateOrCreateUser(userId int64, userName string) *ent.User {
	entUser, err := m.client.User.
		Query().
		Where(user.And(
			user.UserIDEQ(userId),
			user.TypeEQ(user.TypeTelegram),
		)).
		Only(m.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			entUser, err = m.client.User.
				Create().
				SetUserID(userId).
				SetName(userName).
				SetType(user.TypeTelegram).
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

func (m *TelegramChatManager) UpdateOrCreateChat(userId int64, userName string, tgChatId int64, name string, isGroup bool) *ent.TelegramChat {
	entUser := m.updateOrCreateUser(userId, userName)
	entTgChat, err := entUser.
		QueryTelegramChats().
		Where(telegramchat.ChatIDEQ(tgChatId)).
		Only(m.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			entTgChat, err = m.client.TelegramChat.
				Create().
				SetUser(entUser).
				SetChatID(tgChatId).
				SetName(name).
				SetIsGroup(isGroup).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Could not create telegram chat: %v", err)
			}
		} else {
			log.Sugar.Panicf("Could not find telegram chat: %v", err)
		}
	} else if entTgChat.Name != name {
		entTgChat, err = m.client.TelegramChat.
			UpdateOne(entTgChat).
			SetName(name).
			Save(m.ctx)
		if err != nil {
			log.Sugar.Panicf("Could not update telegram chat: %v", err)
		}
	}
	return entTgChat
}
