package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// UserWithZeroId holds the schema definition for the UserWithZeroId entity.
type UserWithZeroId struct {
	ent.Schema
}

func (UserWithZeroId) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// TODO: remove after migration
// Fields of the UserWithZeroId.
func (UserWithZeroId) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("telegram", "discord").
			Immutable(),
		field.Int64("chat_or_channel_id"),
		field.String("chat_or_channel_name"),
		field.Bool("is_group"),
		field.String("chain_id"),
	}
}
