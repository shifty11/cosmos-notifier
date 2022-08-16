package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/types"
)

type ContractManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewContractManager(client *ent.Client, ctx context.Context) *ContractManager {
	return &ContractManager{client: client, ctx: ctx}
}

// CreateOrUpdate creates a new contract or updates an existing one
//
// returns (contract, created)
func (m *ContractManager) CreateOrUpdate(contractAddr string, config *types.ContractData) (*ent.Contract, bool) {
	c, err := m.client.Contract.
		Query().
		Where(contract.AddressEQ(contractAddr)).
		First(m.ctx)
	if err != nil && ent.IsNotFound(err) {
		c, err = m.client.Contract.
			Create().
			SetAddress(contractAddr).
			SetName(config.Name).
			SetDescription(config.Description).
			SetImageURL(config.ImageUrl).
			Save(m.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating contract: %v", err)
		}
		return c, false
	} else if err != nil {
		log.Sugar.Panicf("Error while querying contract: %v", err)
	} else {
		if c.Name != config.Name || c.Description != config.Description || c.ImageURL != config.ImageUrl {
			c, err = m.client.Contract.
				UpdateOne(c).
				SetName(config.Name).
				SetDescription(config.Description).
				SetImageURL(config.ImageUrl).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Error while updating contract: %v", err)
			}
			return c, true
		}
	}
	return c, false
}
