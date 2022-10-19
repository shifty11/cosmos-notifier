package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/ent/contractproposal"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/types"
)

type ContractProposalManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewContractProposalManager(client *ent.Client, ctx context.Context) *ContractProposalManager {
	return &ContractProposalManager{client: client, ctx: ctx}
}

type ProposalStatus string

const (
	ProposalCreated               ProposalStatus = "created"
	ProposalUpdated               ProposalStatus = "updated"
	ProposalStatusChangedFromOpen ProposalStatus = "status_changed_from_open"
	ProposalNoChanges             ProposalStatus = "no_changes"
)

// CreateOrUpdate creates a new proposal or updates an existing one
// Status:
// - created: proposal created
// - updated: proposal updated
// - status_changed_from_open: proposal went from 'open' to another state
// - no_changes: proposal not changed
//
// returns (proposal, status)
func (m *ContractProposalManager) CreateOrUpdate(c *ent.Contract, propData *types.Proposal) (*ent.ContractProposal, ProposalStatus) {
	log.Sugar.Debugf("CreateOrUpdate proposal %v of contract %v", propData.Id, c.Name)
	prop, err := m.client.ContractProposal.
		Query().
		Where(
			contractproposal.And(
				contractproposal.HasContractWith(contract.IDEQ(c.ID)),
				contractproposal.ProposalIDEQ(propData.Id),
			)).
		First(m.ctx)
	if err != nil && ent.IsNotFound(err) {
		prop, err = m.client.ContractProposal.
			Create().
			SetContract(c).
			SetProposalID(propData.Id).
			SetTitle(propData.Title).
			SetDescription(propData.Description).
			SetStatus(contractproposal.Status(propData.Status)).
			SetExpiresAt(propData.Expires.AtTime).
			Save(m.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating ContractProposal: %v", err)
		}
		return prop, ProposalCreated
	} else if err != nil {
		log.Sugar.Errorf("Error while querying ContractProposal: %v", err)
	} else {
		if prop.Title != propData.Title || prop.Description != propData.Description || prop.Status != contractproposal.Status(propData.Status) {
			updatedProp, err := m.client.ContractProposal.
				UpdateOne(prop).
				SetTitle(propData.Title).
				SetDescription(propData.Description).
				SetStatus(contractproposal.Status(propData.Status)).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Error while updating ContractProposal: %v", err)
			}
			if prop.Status == contractproposal.StatusOpen && updatedProp.Status != prop.Status {
				return updatedProp, ProposalStatusChangedFromOpen
			}
			return updatedProp, ProposalUpdated
		}
	}
	return prop, ProposalNoChanges
}
