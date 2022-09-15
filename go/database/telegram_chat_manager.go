package database

import (
	"context"
	"errors"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/ent/telegramchat"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/log"
)

type TelegramChatManager struct {
	client          *ent.Client
	ctx             context.Context
	contractManager IContractManager
	userManager     *UserManager
}

func NewTelegramChatManager(client *ent.Client, ctx context.Context, contractManager IContractManager, userManager *UserManager) *TelegramChatManager {
	return &TelegramChatManager{client: client, ctx: ctx, contractManager: contractManager, userManager: userManager}
}

func (m *TelegramChatManager) AddOrRemoveContract(tgChatId int64, contractId int) (bool, error) {
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

func (m *TelegramChatManager) CreateOrUpdateChat(userId int64, userName string, tgChatId int64, name string, isGroup bool) (tc *ent.TelegramChat, created bool) {
	log.Sugar.Debugf("Create or update Telegram chat %s (%d) for user %s (%d)", name, tgChatId, userName, userId)
	entUser := m.userManager.createOrUpdateUser(userId, userName, user.TypeTelegram)
	tgChat, err := m.client.TelegramChat.
		Query().
		Where(telegramchat.ChatIDEQ(tgChatId)).
		WithUsers().
		Only(m.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			tgChat, err = m.client.TelegramChat.
				Create().
				AddUsers(entUser).
				SetChatID(tgChatId).
				SetName(name).
				SetIsGroup(isGroup).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Could not create telegram chat: %v", err)
			}
			return tgChat, true
		} else {
			log.Sugar.Panicf("Could not find telegram chat: %v", err)
		}
	} else {
		hasUser := false
		for _, u := range tgChat.Edges.Users {
			if u.ID == entUser.ID {
				hasUser = true
				break
			}
		}
		if !hasUser {
			tgChat, err = tgChat.
				Update().
				AddUsers(entUser).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Could not add user to telegram chat: %v", err)
			}
		}
		if tgChat.Name != name {
			tgChat, err = tgChat.
				Update().
				SetName(name).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Could not update telegram chat: %v", err)
			}
		}
	}
	return tgChat, false
}

type TgChatQueryResult struct {
	ChatId int64  `json:"chat_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

func (m *TelegramChatManager) GetSubscribedIds(entContract *ent.Contract) []TgChatQueryResult {
	var v []TgChatQueryResult
	err := entContract.
		QueryTelegramChats().
		Select(telegramchat.FieldChatID, telegramchat.FieldName).
		Scan(m.ctx, &v)
	if err != nil {
		log.Sugar.Panicf("Error while querying Telegram chatIds for contract %v (%v): %v", entContract.Name, entContract.Address, err)
	}
	return v
}

// Delete deletes a telegram chat for a user
// If the user doesn't have any more chats, the user is deleted
func (m *TelegramChatManager) Delete(userId int64, chatId int64) error {
	log.Sugar.Debugf("Deleting telegram chat %d for user %d", chatId, userId)
	telegramChat, err := m.client.TelegramChat.
		Query().
		Where(telegramchat.ChatID(chatId)).
		WithUsers().
		Only(m.ctx)
	if err != nil {
		log.Sugar.Errorf("Could not find telegram chat: %v", err)
		return err
	}
	var entUser *ent.User
	for _, u := range telegramChat.Edges.Users {
		if u.UserID == userId {
			entUser = u
			break
		}
	}
	if entUser == nil {
		log.Sugar.Errorf("Could not find user: %v", err)
		return errors.New("could not find user")
	}
	if len(telegramChat.Edges.Users) == 1 {
		err := m.client.TelegramChat.
			DeleteOne(telegramChat).
			Exec(m.ctx)
		if err != nil {
			log.Sugar.Errorf("Could not delete telegram chat: %v", err)
		}
	} else {
		_, err = m.client.TelegramChat.
			UpdateOne(telegramChat).
			RemoveUsers(entUser).
			Save(m.ctx)
		if err != nil {
			log.Sugar.Errorf("Could not remove user from telegram chat: %v", err)
		}
	}
	m.userManager.deleteIfUnused(entUser)
	return err
}

func (m *TelegramChatManager) DeleteMultiple(chatIds []int64) {
	log.Sugar.Debugf("Delete %v Telegram chat's", len(chatIds))

	for _, chatId := range chatIds {
		tgChats, err := m.client.TelegramChat.
			Query().
			Where(telegramchat.ChatID(chatId)).
			WithUsers().
			All(m.ctx)
		if err != nil {
			log.Sugar.Errorf("Error while querying Telegram chats: %v", err)
		}
		for _, tgChat := range tgChats {
			for _, u := range tgChat.Edges.Users {
				err := m.Delete(u.UserID, chatId)
				if err != nil {
					log.Sugar.Errorf("Error while deleting Telegram chat: %v", err)
				}
			}
		}
	}
}
