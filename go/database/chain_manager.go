package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/types"
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

func (manager *ChainManager) Enable(chainId int, isEnabled bool) error {
	return manager.client.Chain.
		UpdateOneID(chainId).
		SetIsEnabled(isEnabled).
		Exec(manager.ctx)
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

func (manager *ChainManager) Get(id int) (*ent.Chain, error) {
	return manager.client.Chain.
		Query().
		Where(chain.IDEQ(id)).
		Only(manager.ctx)
}

func (manager *ChainManager) GetByName(name string) (*ent.Chain, error) {
	return manager.client.Chain.
		Query().
		Where(chain.NameEQ(name)).
		Only(manager.ctx)
}

func (manager *ChainManager) Create(chainData *types.Chain, thumbnailUrl string) *ent.Chain {
	log.Sugar.Infof("Create new chain: %v (%v)", chainData.PrettyName, chainData.ChainId)
	c, err := manager.client.Chain.
		Create().
		SetChainID(chainData.ChainId).
		SetName(chainData.Name).
		SetPrettyName(chainData.PrettyName).
		SetImageURL(chainData.Image).
		SetThumbnailURL(thumbnailUrl).
		SetIsEnabled(true).
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
		SetChainID(chainData.ChainId).
		SetName(chainData.Name).
		SetPrettyName(chainData.PrettyName).
		SetPath(chainData.Path).
		SetDisplay(chainData.Display).
		SetImageURL(chainData.Image).
		SetThumbnailURL(thumbnailUrl).
		Save(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while updating chain: %v", err)
	}
	return c
}
