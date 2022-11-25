package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/shifty11/cosmos-notifier/types"
)

// Contract holds the schema definition for the Contract entity.
type Contract struct {
	ent.Schema
}

func (Contract) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Contract.
func (Contract) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").
			Unique(),
		field.String("name"),
		field.String("description"),
		field.String("image_url"),
		field.String("thumbnail_url").
			Default(""),
		field.String("rpc_endpoint").
			Default("https://rpc.cosmos.directory/juno"),
		field.Enum("config_version").
			Values(
				types.ContractVersionUnknown.String(),
				types.ContractVersionV1.String(),
				types.ContractVersionV2.String(),
			).
			Default(types.ContractVersionUnknown.String()),
		field.String("get_proposals_query").
			Default("{\"list_proposals\":{}}"),
	}
}

// Edges of the Contract.
func (Contract) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("proposals", ContractProposal.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("telegram_chats", TelegramChat.Type).
			Ref("contracts").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("discord_channels", DiscordChannel.Type).
			Ref("contracts").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Contract) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address").
			Unique(),
	}
}
