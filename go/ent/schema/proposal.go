package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var ProposalStatus = map[int32]string{
	0: "PROPOSAL_STATUS_UNSPECIFIED",
	1: "PROPOSAL_STATUS_DEPOSIT_PERIOD",
	2: "PROPOSAL_STATUS_VOTING_PERIOD",
	3: "PROPOSAL_STATUS_PASSED",
	4: "PROPOSAL_STATUS_REJECTED",
	5: "PROPOSAL_STATUS_FAILED",
}

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
	for _, status := range ProposalStatus {
		statusValues = append(statusValues, status)
	}
	return []ent.Field{
		field.String("proposal_id"),
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
