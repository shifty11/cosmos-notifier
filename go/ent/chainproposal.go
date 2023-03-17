// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/chainproposal"
)

// ChainProposal is the model entity for the ChainProposal schema.
type ChainProposal struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// ProposalID holds the value of the "proposal_id" field.
	ProposalID int `json:"proposal_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// VotingStartTime holds the value of the "voting_start_time" field.
	VotingStartTime time.Time `json:"voting_start_time,omitempty"`
	// VotingEndTime holds the value of the "voting_end_time" field.
	VotingEndTime time.Time `json:"voting_end_time,omitempty"`
	// Status holds the value of the "status" field.
	Status chainproposal.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChainProposalQuery when eager-loading is set.
	Edges                 ChainProposalEdges `json:"edges"`
	chain_chain_proposals *int
}

// ChainProposalEdges holds the relations/edges for other nodes in the graph.
type ChainProposalEdges struct {
	// Chain holds the value of the chain edge.
	Chain *Chain `json:"chain,omitempty"`
	// AddressTracker holds the value of the address_tracker edge.
	AddressTracker []*AddressTracker `json:"address_tracker,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChainOrErr returns the Chain value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChainProposalEdges) ChainOrErr() (*Chain, error) {
	if e.loadedTypes[0] {
		if e.Chain == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: chain.Label}
		}
		return e.Chain, nil
	}
	return nil, &NotLoadedError{edge: "chain"}
}

// AddressTrackerOrErr returns the AddressTracker value or an error if the edge
// was not loaded in eager-loading.
func (e ChainProposalEdges) AddressTrackerOrErr() ([]*AddressTracker, error) {
	if e.loadedTypes[1] {
		return e.AddressTracker, nil
	}
	return nil, &NotLoadedError{edge: "address_tracker"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChainProposal) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chainproposal.FieldID, chainproposal.FieldProposalID:
			values[i] = new(sql.NullInt64)
		case chainproposal.FieldTitle, chainproposal.FieldDescription, chainproposal.FieldStatus:
			values[i] = new(sql.NullString)
		case chainproposal.FieldCreateTime, chainproposal.FieldUpdateTime, chainproposal.FieldVotingStartTime, chainproposal.FieldVotingEndTime:
			values[i] = new(sql.NullTime)
		case chainproposal.ForeignKeys[0]: // chain_chain_proposals
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ChainProposal", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChainProposal fields.
func (cp *ChainProposal) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chainproposal.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cp.ID = int(value.Int64)
		case chainproposal.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				cp.CreateTime = value.Time
			}
		case chainproposal.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				cp.UpdateTime = value.Time
			}
		case chainproposal.FieldProposalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field proposal_id", values[i])
			} else if value.Valid {
				cp.ProposalID = int(value.Int64)
			}
		case chainproposal.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				cp.Title = value.String
			}
		case chainproposal.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				cp.Description = value.String
			}
		case chainproposal.FieldVotingStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field voting_start_time", values[i])
			} else if value.Valid {
				cp.VotingStartTime = value.Time
			}
		case chainproposal.FieldVotingEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field voting_end_time", values[i])
			} else if value.Valid {
				cp.VotingEndTime = value.Time
			}
		case chainproposal.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				cp.Status = chainproposal.Status(value.String)
			}
		case chainproposal.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field chain_chain_proposals", value)
			} else if value.Valid {
				cp.chain_chain_proposals = new(int)
				*cp.chain_chain_proposals = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryChain queries the "chain" edge of the ChainProposal entity.
func (cp *ChainProposal) QueryChain() *ChainQuery {
	return NewChainProposalClient(cp.config).QueryChain(cp)
}

// QueryAddressTracker queries the "address_tracker" edge of the ChainProposal entity.
func (cp *ChainProposal) QueryAddressTracker() *AddressTrackerQuery {
	return NewChainProposalClient(cp.config).QueryAddressTracker(cp)
}

// Update returns a builder for updating this ChainProposal.
// Note that you need to call ChainProposal.Unwrap() before calling this method if this ChainProposal
// was returned from a transaction, and the transaction was committed or rolled back.
func (cp *ChainProposal) Update() *ChainProposalUpdateOne {
	return NewChainProposalClient(cp.config).UpdateOne(cp)
}

// Unwrap unwraps the ChainProposal entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cp *ChainProposal) Unwrap() *ChainProposal {
	_tx, ok := cp.config.driver.(*txDriver)
	if !ok {
		panic("ent: ChainProposal is not a transactional entity")
	}
	cp.config.driver = _tx.drv
	return cp
}

// String implements the fmt.Stringer.
func (cp *ChainProposal) String() string {
	var builder strings.Builder
	builder.WriteString("ChainProposal(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cp.ID))
	builder.WriteString("create_time=")
	builder.WriteString(cp.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(cp.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("proposal_id=")
	builder.WriteString(fmt.Sprintf("%v", cp.ProposalID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(cp.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(cp.Description)
	builder.WriteString(", ")
	builder.WriteString("voting_start_time=")
	builder.WriteString(cp.VotingStartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("voting_end_time=")
	builder.WriteString(cp.VotingEndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", cp.Status))
	builder.WriteByte(')')
	return builder.String()
}

// ChainProposals is a parsable slice of ChainProposal.
type ChainProposals []*ChainProposal
