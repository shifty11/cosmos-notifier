package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/shifty11/cosmos-notifier/types"
)

// ContractProposal holds the schema definition for the ContractProposal entity.
type ContractProposal struct {
	ent.Schema
}

func (ContractProposal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the ContractProposal.
func (ContractProposal) Fields() []ent.Field {
	var statusValues []string
	for _, status := range types.ProposalStatusValues {
		statusValues = append(statusValues, string(status))
	}
	return []ent.Field{
		field.Int("proposal_id"),
		field.String("title"),
		field.String("description"),
		field.Time("expires_at"),
		field.Enum("status").
			Values(statusValues...),
	}
}

// Edges of the ContractProposal.
func (ContractProposal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("contract", Contract.Type).
			Ref("proposals").
			Unique(),
	}
}
