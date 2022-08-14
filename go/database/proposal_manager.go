package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/ent/proposal"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/types"
)

type ProposalManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewProposalManager(client *ent.Client, ctx context.Context) *ProposalManager {
	return &ProposalManager{client: client, ctx: ctx}
}

type ProposalStatus string

const (
	ProposalCreated       ProposalStatus = "created"
	ProposalUpdated       ProposalStatus = "updated"
	ProposalStatusChanged ProposalStatus = "status_changed"
	ProposalNoChanges     ProposalStatus = "no_changes"
)

// CreateOrUpdate creates a new proposal or updates an existing one
// Status:
// - created: proposal created
// - updated: proposal updated
// - status_changed: proposal went from 'open' to another state
// - no_changes: proposal not changed
//
// returns (proposal, status)
func (manager *ProposalManager) CreateOrUpdate(c *ent.Contract, propData *types.Proposal) (*ent.Proposal, ProposalStatus) {
	prop, err := manager.client.Proposal.
		Query().
		Where(
			proposal.And(
				proposal.HasContractWith(contract.IDEQ(c.ID)),
				proposal.ProposalIDEQ(propData.Id),
			)).
		First(manager.ctx)
	if err != nil && ent.IsNotFound(err) {
		prop, err = manager.client.Proposal.
			Create().
			SetContract(c).
			SetProposalID(propData.Id).
			SetTitle(propData.Title).
			SetDescription(propData.Description).
			SetStatus(proposal.Status(propData.Status)).
			SetExpiresAt(propData.Expires.AtTime).
			Save(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating proposal: %v", err)
		}
		return prop, ProposalCreated
	} else if err != nil {
		log.Sugar.Errorf("Error while querying proposal: %v", err)
	} else {
		if prop.Title != propData.Title || prop.Description != propData.Description || prop.Status != proposal.Status(propData.Status) {
			updatedProp, err := manager.client.Proposal.
				UpdateOne(prop).
				SetTitle(propData.Title).
				SetDescription(propData.Description).
				SetStatus(proposal.Status(propData.Status)).
				Save(manager.ctx)
			if err != nil {
				log.Sugar.Panicf("Error while updating proposal: %v", err)
			}
			if prop.Status == proposal.StatusOpen && updatedProp.Status != prop.Status {
				return updatedProp, ProposalStatusChanged
			}
			return updatedProp, ProposalUpdated
		}
	}
	return prop, ProposalNoChanges
}
