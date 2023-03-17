package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TelegramChat holds the schema definition for the TelegramChat entity.
type TelegramChat struct {
	ent.Schema
}

func (TelegramChat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the TelegramChat.
func (TelegramChat) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("chat_id").
			Unique(),
		field.String("name"),
		field.Bool("is_group").
			Immutable(),
	}
}

// Edges of the TelegramChat.
func (TelegramChat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Required(),
		edge.To("contracts", Contract.Type),
		edge.To("chains", Chain.Type),
		edge.To("address_trackers", AddressTracker.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
