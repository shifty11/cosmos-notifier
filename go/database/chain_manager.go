package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/chain"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/types"
)

type ChainManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewChainManager(client *ent.Client, ctx context.Context) *ChainManager {
	return &ChainManager{client: client, ctx: ctx}
}

func (manager *ChainManager) ByName(name string) (*ent.Chain, error) {
	return manager.client.Chain.
		Query().
		Where(chain.NameEQ(name)).
		Only(manager.ctx)
}

func (manager *ChainManager) Enabled() []*ent.Chain {
	query := manager.client.Chain.
		Query().
		Where(chain.IsEnabledEQ(true)).
		Order(ent.Asc(chain.FieldPrettyName))
	allChains, err := query.All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying enabled chains: %v", err)
	}
	return allChains
}

func (manager *ChainManager) All() []*ent.Chain {
	chains, err := manager.client.Chain.
		Query().
		Order(ent.Asc(chain.FieldPrettyName)).
		All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying chains: %v", err)
	}
	return chains
}

func (manager *ChainManager) Create(chainData *types.Chain, thumbnailUrl string) *ent.Chain {
	log.Sugar.Infof("Create new chain: %v", chainData.PrettyName)
	c, err := manager.client.Chain.
		Create().
		SetChainID(chainData.ChainId).
		SetName(chainData.Name).
		SetPrettyName(chainData.PrettyName).
		SetImageURL(chainData.Image).
		SetThumbnailURL(thumbnailUrl).
		SetIsEnabled(false).
		Save(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while creating chain: %v", err)
	}
	return c
}

func (manager *ChainManager) Update(entChain *ent.Chain, chainData *types.Chain, thumbnailUrl string) *ent.Chain {
	log.Sugar.Infof("Update chain: %v", chainData.PrettyName)
	c, err := entChain.
		Update().
		SetName(chainData.Name).
		SetPrettyName(chainData.PrettyName).
		SetImageURL(chainData.Image).
		SetThumbnailURL(thumbnailUrl).
		Save(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while updating chain: %v", err)
	}
	return c
}
