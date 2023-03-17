package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// AddressTracker holds the schema definition for the AddressTracker entity.
type AddressTracker struct {
	ent.Schema
}

func (AddressTracker) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the AddressTracker.
func (AddressTracker) Fields() []ent.Field {
	return []ent.Field{
		field.String("address"),
	}
}

// Edges of the AddressTracker.
func (AddressTracker) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chain", Chain.Type).
			Ref("address_trackers").
			Unique().
			Required(),
		edge.From("discord_channel", DiscordChannel.Type).
			Ref("address_trackers").
			Unique(),
		edge.From("telegram_chat", TelegramChat.Type).
			Ref("address_trackers").
			Unique(),
		edge.To("chain_proposals", ChainProposal.Type),
	}
}

func (AddressTracker) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address").
			Edges("chain", "discord_channel", "telegram_chat").
			Unique(),
	}
}
