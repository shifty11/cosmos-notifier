package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"errors"
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
			Immutable().
			Unique().
			Validate(func(s string) error {
				if s == "" {
					return errors.New("address is empty")
				}
				return nil
			}),
		field.String("moniker").
			Validate(func(s string) error {
				if s == "" {
					return errors.New("moniker is empty")
				}
				return nil
			}),
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
		edge.To("telegram_chats", TelegramChat.Type),
		edge.To("discord_channels", DiscordChannel.Type),
	}
}

func (Validator) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address"),
		index.Fields("moniker"),
		index.Fields("moniker", "address").
			Edges("chain").
			Unique(),
	}
}
