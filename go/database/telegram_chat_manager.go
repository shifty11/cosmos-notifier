package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/ent/telegramchat"
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
