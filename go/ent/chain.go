// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/dao-dao-notifier/ent/chain"
)

// Chain is the model entity for the Chain schema.
type Chain struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// ChainID holds the value of the "chain_id" field.
	ChainID string `json:"chain_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// PrettyName holds the value of the "pretty_name" field.
	PrettyName string `json:"pretty_name,omitempty"`
	// IsEnabled holds the value of the "is_enabled" field.
	IsEnabled bool `json:"is_enabled,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// ThumbnailURL holds the value of the "thumbnail_url" field.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChainQuery when eager-loading is set.
	Edges ChainEdges `json:"edges"`
}

// ChainEdges holds the relations/edges for other nodes in the graph.
type ChainEdges struct {
	// ChainProposals holds the value of the chain_proposals edge.
	ChainProposals []*ChainProposal `json:"chain_proposals,omitempty"`
	// TelegramChats holds the value of the telegram_chats edge.
	TelegramChats []*TelegramChat `json:"telegram_chats,omitempty"`
	// DiscordChannels holds the value of the discord_channels edge.
	DiscordChannels []*DiscordChannel `json:"discord_channels,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ChainProposalsOrErr returns the ChainProposals value or an error if the edge
// was not loaded in eager-loading.
func (e ChainEdges) ChainProposalsOrErr() ([]*ChainProposal, error) {
	if e.loadedTypes[0] {
		return e.ChainProposals, nil
	}
	return nil, &NotLoadedError{edge: "chain_proposals"}
}

// TelegramChatsOrErr returns the TelegramChats value or an error if the edge
// was not loaded in eager-loading.
func (e ChainEdges) TelegramChatsOrErr() ([]*TelegramChat, error) {
	if e.loadedTypes[1] {
		return e.TelegramChats, nil
	}
	return nil, &NotLoadedError{edge: "telegram_chats"}
}

// DiscordChannelsOrErr returns the DiscordChannels value or an error if the edge
// was not loaded in eager-loading.
func (e ChainEdges) DiscordChannelsOrErr() ([]*DiscordChannel, error) {
	if e.loadedTypes[2] {
		return e.DiscordChannels, nil
	}
	return nil, &NotLoadedError{edge: "discord_channels"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Chain) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case chain.FieldIsEnabled:
			values[i] = new(sql.NullBool)
		case chain.FieldID:
			values[i] = new(sql.NullInt64)
		case chain.FieldChainID, chain.FieldName, chain.FieldPrettyName, chain.FieldImageURL, chain.FieldThumbnailURL:
			values[i] = new(sql.NullString)
		case chain.FieldCreateTime, chain.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Chain", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Chain fields.
func (c *Chain) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chain.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case chain.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case chain.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case chain.FieldChainID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_id", values[i])
			} else if value.Valid {
				c.ChainID = value.String
			}
		case chain.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case chain.FieldPrettyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pretty_name", values[i])
			} else if value.Valid {
				c.PrettyName = value.String
			}
		case chain.FieldIsEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_enabled", values[i])
			} else if value.Valid {
				c.IsEnabled = value.Bool
			}
		case chain.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				c.ImageURL = value.String
			}
		case chain.FieldThumbnailURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail_url", values[i])
			} else if value.Valid {
				c.ThumbnailURL = value.String
			}
		}
	}
	return nil
}

// QueryChainProposals queries the "chain_proposals" edge of the Chain entity.
func (c *Chain) QueryChainProposals() *ChainProposalQuery {
	return (&ChainClient{config: c.config}).QueryChainProposals(c)
}

// QueryTelegramChats queries the "telegram_chats" edge of the Chain entity.
func (c *Chain) QueryTelegramChats() *TelegramChatQuery {
	return (&ChainClient{config: c.config}).QueryTelegramChats(c)
}

// QueryDiscordChannels queries the "discord_channels" edge of the Chain entity.
func (c *Chain) QueryDiscordChannels() *DiscordChannelQuery {
	return (&ChainClient{config: c.config}).QueryDiscordChannels(c)
}

// Update returns a builder for updating this Chain.
// Note that you need to call Chain.Unwrap() before calling this method if this Chain
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Chain) Update() *ChainUpdateOne {
	return (&ChainClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Chain entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Chain) Unwrap() *Chain {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Chain is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Chain) String() string {
	var builder strings.Builder
	builder.WriteString("Chain(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("chain_id=")
	builder.WriteString(c.ChainID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("pretty_name=")
	builder.WriteString(c.PrettyName)
	builder.WriteString(", ")
	builder.WriteString("is_enabled=")
	builder.WriteString(fmt.Sprintf("%v", c.IsEnabled))
	builder.WriteString(", ")
	builder.WriteString("image_url=")
	builder.WriteString(c.ImageURL)
	builder.WriteString(", ")
	builder.WriteString("thumbnail_url=")
	builder.WriteString(c.ThumbnailURL)
	builder.WriteByte(')')
	return builder.String()
}

// Chains is a parsable slice of Chain.
type Chains []*Chain

func (c Chains) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}