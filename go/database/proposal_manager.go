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

// CreateOrUpdate creates a new proposal or updates an existing one
//
// returns (proposal, created)
func (manager *ProposalManager) CreateOrUpdate(c *ent.Contract, propData *types.Proposal) (*ent.Proposal, bool) {
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
		return prop, true
	} else if err != nil {
		log.Sugar.Errorf("Error while querying proposal: %v", err)
	} else {
		if prop.Title != propData.Title || prop.Description != propData.Description || prop.Status != proposal.Status(propData.Status) {
			prop, err = manager.client.Proposal.
				UpdateOne(prop).
				SetTitle(propData.Title).
				SetDescription(propData.Description).
				SetStatus(proposal.Status(propData.Status)).
				Save(manager.ctx)
			if err != nil {
				log.Sugar.Panicf("Error while updating proposal: %v", err)
			}
		}
	}
	return prop, false
}
