package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Chain holds the schema definition for the Chain entity.
type Chain struct {
	ent.Schema
}

func (Chain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Chain.
func (Chain) Fields() []ent.Field {
	return []ent.Field{
		field.String("chain_id"),
		field.String("name"),
		field.String("pretty_name"),
		field.String("path"),
		field.String("display").
			Default(""),
		field.Bool("is_enabled").
			Default(true),
		field.String("image_url"),
		field.String("thumbnail_url").
			Default(""),
		field.String("bech32_prefix"),
	}
}

// Edges of the Chain.
func (Chain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chain_proposals", ChainProposal.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("telegram_chats", TelegramChat.Type).
			Ref("chains").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("discord_channels", DiscordChannel.Type).
			Ref("chains").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("address_trackers", AddressTracker.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("validators", Validator.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Chain) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}
