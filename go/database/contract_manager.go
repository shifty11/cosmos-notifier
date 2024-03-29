package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/contract"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/types"
)

type IContractManager interface {
	Create(data *types.ContractData) (*ent.Contract, error)
	Update(c *ent.Contract, data *types.ContractData) *ent.Contract
	QueryAll() []*ent.Contract
	QueryById(id int) (*ent.Contract, error)
	UpdateSetThumbnailUrl(entContract *ent.Contract, url string) *ent.Contract
	QueryByAddress(contractAddress string) (*ent.Contract, error)
	Delete(id int) error
}

type ContractManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewContractManager(client *ent.Client, ctx context.Context) *ContractManager {
	return &ContractManager{client: client, ctx: ctx}
}

type ContractStatus string

func (m *ContractManager) Create(data *types.ContractData) (*ent.Contract, error) {
	c, err := m.client.Contract.
		Create().
		SetAddress(data.Address).
		SetName(data.Name).
		SetDescription(data.Description).
		SetImageURL(data.ImageUrl).
		SetRPCEndpoint(data.RpcEndpoint).
		SetConfigVersion(contract.ConfigVersion(data.ContractVersion)).
		Save(m.ctx)
	return c, err
}

// Update creates a new contract or updates an existing one
// returns (contract, created)
func (m *ContractManager) Update(c *ent.Contract, data *types.ContractData) *ent.Contract {
	if c.Name != data.Name ||
		c.Description != data.Description ||
		c.ImageURL != data.ImageUrl ||
		c.RPCEndpoint != data.RpcEndpoint ||
		c.ConfigVersion != contract.ConfigVersion(data.ContractVersion) {
		updated, err := m.client.Contract.
			UpdateOne(c).
			SetName(data.Name).
			SetDescription(data.Description).
			SetImageURL(data.ImageUrl).
			SetRPCEndpoint(data.RpcEndpoint).
			SetConfigVersion(contract.ConfigVersion(data.ContractVersion)).
			Save(m.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while updating contract: %v", err)
		}
		return updated
	}
	return c
}

func (m *ContractManager) QueryAll() []*ent.Contract {
	all, err := m.client.Contract.
		Query().
		Order(ent.Asc(contract.FieldName)).
		All(m.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying contracts: %v", err)
	}
	return all
}

func (m *ContractManager) QueryById(id int) (*ent.Contract, error) {
	return m.client.Contract.
		Query().
		Where(contract.ID(id)).
		Only(m.ctx)
}

func (m *ContractManager) UpdateSetThumbnailUrl(entContract *ent.Contract, url string) *ent.Contract {
	updated, err := m.client.Contract.
		UpdateOne(entContract).
		SetThumbnailURL(url).
		Save(m.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while updating contract: %v", err)
	}
	return updated
}

func (m *ContractManager) QueryByAddress(contractAddress string) (*ent.Contract, error) {
	return m.client.Contract.
		Query().
		Where(contract.AddressEQ(contractAddress)).
		Only(m.ctx)
}

func (m *ContractManager) Delete(id int) error {
	return m.client.Contract.
		DeleteOneID(id).
		Exec(m.ctx)
}
