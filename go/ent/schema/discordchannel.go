package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// DiscordChannel holds the schema definition for the DiscordChannel entity.
type DiscordChannel struct {
	ent.Schema
}

func (DiscordChannel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the DiscordChannel.
func (DiscordChannel) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("channel_id").
			Unique(),
		field.String("name"),
		field.Bool("is_group").
			Immutable(),
	}
}

// Edges of the DiscordChannel.
func (DiscordChannel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Required(),
		edge.To("contracts", Contract.Type),
		edge.To("chains", Chain.Type),
		edge.To("address_trackers", AddressTracker.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("validators", Validator.Type).
			Ref("discord_channels"),
	}
}
