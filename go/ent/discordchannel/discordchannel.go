// Code generated by ent, DO NOT EDIT.

package discordchannel

import (
	"time"
)

const (
	// Label holds the string label denoting the discordchannel type in the database.
	Label = "discord_channel"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldChannelID holds the string denoting the channel_id field in the database.
	FieldChannelID = "channel_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIsGroup holds the string denoting the is_group field in the database.
	FieldIsGroup = "is_group"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeChains holds the string denoting the chains edge name in mutations.
	EdgeChains = "chains"
	// Table holds the table name of the discordchannel in the database.
	Table = "discord_channels"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "discord_channels"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "discord_channel_user"
	// ChainsTable is the table that holds the chains relation/edge.
	ChainsTable = "contracts"
	// ChainsInverseTable is the table name for the Contract entity.
	// It exists in this package in order to avoid circular dependency with the "contract" package.
	ChainsInverseTable = "contracts"
	// ChainsColumn is the table column denoting the chains relation/edge.
	ChainsColumn = "discord_channel_chains"
)

// Columns holds all SQL columns for discordchannel fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldChannelID,
	FieldName,
	FieldIsGroup,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "discord_channels"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"discord_channel_user",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
