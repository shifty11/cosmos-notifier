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
		field.String("chain_id").
			Unique(),
		field.String("name").
			Unique(),
		field.String("pretty_name").
			Unique(),
		field.String("path").
			Default(""), //TODO: remove this field after migration
		field.String("display").
			Default(""), //TODO: remove this field after migration
		field.Bool("is_enabled").
			Default(true),
		field.String("image_url"),
		field.String("thumbnail_url").
			Default(""),
	}
}

// Edges of the Chain.
func (Chain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chain_proposals", ChainProposal.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		//edge.To("draft_proposals", DraftProposal.Type).
		//	Annotations(entsql.Annotation{
		//		OnDelete: entsql.Cascade,
		//	}),
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
	}
}

func (Chain) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
	}
}
