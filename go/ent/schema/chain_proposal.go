package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// ChainProposal holds the schema definition for the ChainProposal entity.
type ChainProposal struct {
	ent.Schema
}

func (ChainProposal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the ChainProposal.
func (ChainProposal) Fields() []ent.Field {
	var statusValues []string
	for _, status := range v1beta1.ProposalStatus_name {
		statusValues = append(statusValues, status)
	}
	return []ent.Field{
		field.Int("proposal_id"),
		field.String("title"),
		field.String("description"),
		field.Time("voting_start_time"),
		field.Time("voting_end_time"),
		field.Enum("status").
			Values(statusValues...),
	}
}

// Edges of the ChainProposal.
func (ChainProposal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chain", Chain.Type).
			Ref("chain_proposals").
			Unique(),
		edge.From("address_tracker", AddressTracker.Type).
			Ref("chain_proposals"),
	}
}
