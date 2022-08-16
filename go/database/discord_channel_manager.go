package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/ent/discordchannel"
)

type DiscordChannelManager struct {
	client          *ent.Client
	ctx             context.Context
	contractManager *ContractManager
}

func NewDiscordChannelManager(client *ent.Client, ctx context.Context, contractManager *ContractManager) *DiscordChannelManager {
	return &DiscordChannelManager{client: client, ctx: ctx, contractManager: contractManager}
}

func (m *DiscordChannelManager) AddOrRemoveChain(dChannelId int64, contractId int) (bool, error) {
	dChannel, err := m.client.DiscordChannel.
		Query().
		Where(discordchannel.ChannelID(dChannelId)).
		First(m.ctx)
	if err != nil {
		return false, err
	}

	dbContract, err := m.contractManager.Get(contractId)
	if err != nil {
		return false, err
	}

	exists, err := dChannel.
		QueryContracts().
		Where(contract.IDEQ(dbContract.ID)).
		Exist(m.ctx)
	if err != nil {
		return false, err
	}
	if exists {
		_, err := dChannel.
			Update().
			RemoveContractIDs(dbContract.ID).
			Save(m.ctx)
		if err != nil {
			return false, err
		}
	} else {
		_, err := dChannel.
			Update().
			AddContractIDs(dbContract.ID).
			Save(m.ctx)
		if err != nil {
			return false, err
		}
	}
	return !exists, nil
}
