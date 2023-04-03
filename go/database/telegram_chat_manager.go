package database

import (
	"context"
	"errors"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/contract"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/types"
)

type ITelegramChatManager interface {
	UpdateAddOrRemoveChain(tgChatId int64, chainId int) (hasContract bool, err error)
	UpdateAddOrRemoveContract(tgChatId int64, contractId int) (hasContract bool, err error)
	CreateOrUpdate(userId int64, userName string, tgChatId int64, name string, isGroup bool) (tc *ent.TelegramChat, created bool)
	QuerySubscribedIds(query *ent.TelegramChatQuery) []types.TgChatQueryResult
	Delete(userId int64, chatId int64) error
	DeleteMultiple(chatIds []int64)
	QuerySubscriptionsCount(chatId int64) int
	QueryUsers(chatId int64) []*ent.User
	QueryAllIds() []types.TgChatQueryResult
}

type TelegramChatManager struct {
	client          *ent.Client
	ctx             context.Context
	contractManager IContractManager
	chainManager    *ChainManager
	userManager     *UserManager
}

func NewTelegramChatManager(
	client *ent.Client,
	ctx context.Context,
	chainManager *ChainManager,
	contractManager IContractManager,
	userManager *UserManager,
) *TelegramChatManager {
	return &TelegramChatManager{
		client:          client,
		ctx:             ctx,
		chainManager:    chainManager,
		contractManager: contractManager,
		userManager:     userManager,
	}
}

// UpdateAddOrRemoveChain adds or removes a chain from a telegram chat
// Returns true if the contract is now added to the chat
func (m *TelegramChatManager) UpdateAddOrRemoveChain(tgChatId int64, chainId int) (hasContract bool, err error) {
	log.Sugar.Debugf("Adding or removing chain %d from telegram chat %d", chainId, tgChatId)
	tgChat, err := m.client.TelegramChat.
		Query().
		Where(telegramchat.ChatIDEQ(tgChatId)).
		First(m.ctx)
	if err != nil {
		return false, err
	}

	entChain, err := m.chainManager.QueryById(chainId)
	if err != nil {
		return false, err
	}

	exists, err := tgChat.
		QueryChains().
		Where(chain.IDEQ(entChain.ID)).
		Exist(m.ctx)
	if err != nil {
		return false, err
	}
	if exists {
		_, err := tgChat.
			Update().
			RemoveChainIDs(entChain.ID).
			Save(m.ctx)
		if err != nil {
			return false, err
		}
	} else {
		_, err := tgChat.
			Update().
			AddChainIDs(entChain.ID).
			Save(m.ctx)
		if err != nil {
			return false, err
		}
	}
	return !exists, nil
}

// UpdateAddOrRemoveContract adds or removes a contract from a telegram chat
// Returns true if the contract is now added to the chat
func (m *TelegramChatManager) UpdateAddOrRemoveContract(tgChatId int64, contractId int) (hasContract bool, err error) {
	log.Sugar.Debugf("Adding or removing contract %d from telegram chat %d", contractId, tgChatId)
	tgChat, err := m.client.TelegramChat.
		Query().
		Where(telegramchat.ChatIDEQ(tgChatId)).
		First(m.ctx)
	if err != nil {
		return false, err
	}

	dbContract, err := m.contractManager.QueryById(contractId)
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

func (m *TelegramChatManager) CreateOrUpdate(userId int64, userName string, tgChatId int64, name string, isGroup bool) (tc *ent.TelegramChat, created bool) {
	log.Sugar.Debugf("Create or update Telegram chat %s (%d) for user %s (%d)", name, tgChatId, userName, userId)
	entUser := m.userManager.createOrUpdate(userId, userName, user.TypeTelegram)
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

func (m *TelegramChatManager) QuerySubscribedIds(query *ent.TelegramChatQuery) []types.TgChatQueryResult {
	var v []types.TgChatQueryResult
	err := query.
		Select(telegramchat.FieldChatID, telegramchat.FieldName).
		Scan(m.ctx, &v)
	if err != nil {
		log.Sugar.Panicf("Error while querying Telegram chatIds: %v", err)
	}
	return v
}

func (m *TelegramChatManager) QueryAllIds() []types.TgChatQueryResult {
	return m.QuerySubscribedIds(m.client.TelegramChat.Query())
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

func (m *TelegramChatManager) QuerySubscriptionsCount(chatId int64) int {
	countChains, err := m.client.TelegramChat.
		Query().
		Where(telegramchat.ChatIDEQ(chatId)).
		QueryChains().
		Count(m.ctx)
	if err != nil {
		log.Sugar.Errorf("Could not count chains subscriptions for telegram chat: %v", err)
	}
	countContracts, err := m.client.TelegramChat.
		Query().
		Where(telegramchat.ChatIDEQ(chatId)).
		QueryContracts().
		Count(m.ctx)
	if err != nil {
		log.Sugar.Errorf("Could not count contract subscriptions for telegram chat: %v", err)
	}
	return countChains + countContracts
}

func (m *TelegramChatManager) QueryUsers(chatId int64) []*ent.User {
	users, err := m.client.TelegramChat.
		Query().
		Where(telegramchat.ChatIDEQ(chatId)).
		QueryUsers().
		All(m.ctx)
	if err != nil {
		log.Sugar.Errorf("Could not get users for telegram chat: %v", err)
	}
	return users
}
