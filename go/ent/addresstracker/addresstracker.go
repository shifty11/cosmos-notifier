// Code generated by ent, DO NOT EDIT.

package addresstracker

import (
	"time"
)

const (
	// Label holds the string label denoting the addresstracker type in the database.
	Label = "address_tracker"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// EdgeChain holds the string denoting the chain edge name in mutations.
	EdgeChain = "chain"
	// EdgeDiscordChannel holds the string denoting the discord_channel edge name in mutations.
	EdgeDiscordChannel = "discord_channel"
	// EdgeTelegramChat holds the string denoting the telegram_chat edge name in mutations.
	EdgeTelegramChat = "telegram_chat"
	// EdgeChainProposals holds the string denoting the chain_proposals edge name in mutations.
	EdgeChainProposals = "chain_proposals"
	// Table holds the table name of the addresstracker in the database.
	Table = "address_trackers"
	// ChainTable is the table that holds the chain relation/edge.
	ChainTable = "address_trackers"
	// ChainInverseTable is the table name for the Chain entity.
	// It exists in this package in order to avoid circular dependency with the "chain" package.
	ChainInverseTable = "chains"
	// ChainColumn is the table column denoting the chain relation/edge.
	ChainColumn = "chain_address_trackers"
	// DiscordChannelTable is the table that holds the discord_channel relation/edge.
	DiscordChannelTable = "address_trackers"
	// DiscordChannelInverseTable is the table name for the DiscordChannel entity.
	// It exists in this package in order to avoid circular dependency with the "discordchannel" package.
	DiscordChannelInverseTable = "discord_channels"
	// DiscordChannelColumn is the table column denoting the discord_channel relation/edge.
	DiscordChannelColumn = "discord_channel_address_trackers"
	// TelegramChatTable is the table that holds the telegram_chat relation/edge.
	TelegramChatTable = "address_trackers"
	// TelegramChatInverseTable is the table name for the TelegramChat entity.
	// It exists in this package in order to avoid circular dependency with the "telegramchat" package.
	TelegramChatInverseTable = "telegram_chats"
	// TelegramChatColumn is the table column denoting the telegram_chat relation/edge.
	TelegramChatColumn = "telegram_chat_address_trackers"
	// ChainProposalsTable is the table that holds the chain_proposals relation/edge. The primary key declared below.
	ChainProposalsTable = "address_tracker_chain_proposals"
	// ChainProposalsInverseTable is the table name for the ChainProposal entity.
	// It exists in this package in order to avoid circular dependency with the "chainproposal" package.
	ChainProposalsInverseTable = "chain_proposals"
)

// Columns holds all SQL columns for addresstracker fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldAddress,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "address_trackers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"chain_address_trackers",
	"discord_channel_address_trackers",
	"telegram_chat_address_trackers",
}

var (
	// ChainProposalsPrimaryKey and ChainProposalsColumn2 are the table columns denoting the
	// primary key for the chain_proposals relation (M2M).
	ChainProposalsPrimaryKey = []string{"address_tracker_id", "chain_proposal_id"}
)

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
