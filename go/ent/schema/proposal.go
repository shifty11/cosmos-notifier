package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/shifty11/dao-dao-notifier/types"
)

// Proposal holds the schema definition for the Proposal entity.
type Proposal struct {
	ent.Schema
}

func (Proposal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Proposal.
func (Proposal) Fields() []ent.Field {
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

// Edges of the Proposal.
func (Proposal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("contract", Contract.Type).
			Ref("proposals").
			Unique(),
	}
}
