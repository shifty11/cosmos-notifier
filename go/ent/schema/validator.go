package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Validator holds the schema definition for the Validator entity.
type Validator struct {
	ent.Schema
}

func (Validator) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Validator.
func (Validator) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").
			Immutable(),
		field.String("moniker"),
	}
}

// Edges of the Validator.
func (Validator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chain", Chain.Type).
			Ref("validators").
			Unique().
			Required(),
		edge.To("address_trackers", AddressTracker.Type),
	}
}

func (Validator) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address"),
		index.Fields("moniker"),
	}
}
