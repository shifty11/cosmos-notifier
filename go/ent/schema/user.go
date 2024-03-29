package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.String("name"),
		field.Enum("type").
			Values("telegram", "discord").
			Immutable(),
		field.Enum("role").
			Values("user", "admin").
			Default("user"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("telegram_chats", TelegramChat.Type).
			Ref("users").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("discord_channels", DiscordChannel.Type).
			Ref("users").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
