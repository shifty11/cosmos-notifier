package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/chainproposal"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/types"
	"time"
)

type ChainProposalManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewChainProposalManager(client *ent.Client, ctx context.Context) *ChainProposalManager {
	return &ChainProposalManager{client: client, ctx: ctx}
}

// CreateOrUpdate creates a new proposal or updates an existing one
// Status:
// - created: proposal created
// - updated: proposal updated
// - status_changed_from_open: proposal went from 'open' to another state
// - no_changes: proposal not changed
//
// returns (proposal, status)
func (m *ChainProposalManager) CreateOrUpdate(c *ent.Chain, propData *types.ChainProposal) (*ent.ChainProposal, ProposalStatus) {
	log.Sugar.Debugf("CreateOrUpdate proposal %v of chain %v", propData.ProposalId, c.PrettyName)
	prop, err := m.client.ChainProposal.
		Query().
		Where(
			chainproposal.And(
				chainproposal.HasChainWith(chain.IDEQ(c.ID)),
				chainproposal.ProposalIDEQ(propData.ProposalId),
			)).
		First(m.ctx)
	if err != nil && ent.IsNotFound(err) {
		prop, err = m.client.ChainProposal.
			Create().
			SetChain(c).
			SetProposalID(propData.ProposalId).
			SetTitle(propData.Content.Title).
			SetDescription(propData.Content.Description).
			SetStatus(chainproposal.Status(propData.Status.ToString())).
			SetVotingStartTime(propData.VotingStartTime).
			SetVotingEndTime(propData.VotingEndTime).
			Save(m.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating proposal: %v", err)
		}
		return prop, ProposalCreated
	} else if err != nil {
		log.Sugar.Errorf("Error while querying proposal: %v", err)
	} else {
		if prop.Title != propData.Content.Title ||
			prop.Description != propData.Content.Description ||
			prop.Status != chainproposal.Status(propData.Status.ToString()) {
			updatedProp, err := m.client.ChainProposal.
				UpdateOne(prop).
				SetTitle(propData.Content.Title).
				SetDescription(propData.Content.Description).
				SetStatus(chainproposal.Status(propData.Status.ToString())).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Error while updating proposal: %v", err)
			}
			if prop.Status == chainproposal.StatusPROPOSAL_STATUS_VOTING_PERIOD && updatedProp.Status != prop.Status {
				return updatedProp, ProposalStatusChangedFromOpen
			}
			return updatedProp, ProposalUpdated
		}
	}
	return prop, ProposalNoChanges
}

func (m *ChainProposalManager) VotingPeriodExpired(c *ent.Chain) []*ent.ChainProposal {
	result, err := c.QueryChainProposals().
		Where(
			chainproposal.And(
				chainproposal.StatusEQ(chainproposal.StatusPROPOSAL_STATUS_VOTING_PERIOD),
				chainproposal.VotingEndTimeLTE(time.Now()),
			)).
		All(m.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying proposals: %v", err)
	}
	return result
}
