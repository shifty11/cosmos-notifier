// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-notifier/ent/addresstracker"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/chainproposal"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
)

// AddressTrackerCreate is the builder for creating a AddressTracker entity.
type AddressTrackerCreate struct {
	config
	mutation *AddressTrackerMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (atc *AddressTrackerCreate) SetCreateTime(t time.Time) *AddressTrackerCreate {
	atc.mutation.SetCreateTime(t)
	return atc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (atc *AddressTrackerCreate) SetNillableCreateTime(t *time.Time) *AddressTrackerCreate {
	if t != nil {
		atc.SetCreateTime(*t)
	}
	return atc
}

// SetUpdateTime sets the "update_time" field.
func (atc *AddressTrackerCreate) SetUpdateTime(t time.Time) *AddressTrackerCreate {
	atc.mutation.SetUpdateTime(t)
	return atc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (atc *AddressTrackerCreate) SetNillableUpdateTime(t *time.Time) *AddressTrackerCreate {
	if t != nil {
		atc.SetUpdateTime(*t)
	}
	return atc
}

// SetAddress sets the "address" field.
func (atc *AddressTrackerCreate) SetAddress(s string) *AddressTrackerCreate {
	atc.mutation.SetAddress(s)
	return atc
}

// SetChainID sets the "chain" edge to the Chain entity by ID.
func (atc *AddressTrackerCreate) SetChainID(id int) *AddressTrackerCreate {
	atc.mutation.SetChainID(id)
	return atc
}

// SetChain sets the "chain" edge to the Chain entity.
func (atc *AddressTrackerCreate) SetChain(c *Chain) *AddressTrackerCreate {
	return atc.SetChainID(c.ID)
}

// SetDiscordChannelID sets the "discord_channel" edge to the DiscordChannel entity by ID.
func (atc *AddressTrackerCreate) SetDiscordChannelID(id int) *AddressTrackerCreate {
	atc.mutation.SetDiscordChannelID(id)
	return atc
}

// SetNillableDiscordChannelID sets the "discord_channel" edge to the DiscordChannel entity by ID if the given value is not nil.
func (atc *AddressTrackerCreate) SetNillableDiscordChannelID(id *int) *AddressTrackerCreate {
	if id != nil {
		atc = atc.SetDiscordChannelID(*id)
	}
	return atc
}

// SetDiscordChannel sets the "discord_channel" edge to the DiscordChannel entity.
func (atc *AddressTrackerCreate) SetDiscordChannel(d *DiscordChannel) *AddressTrackerCreate {
	return atc.SetDiscordChannelID(d.ID)
}

// SetTelegramChatID sets the "telegram_chat" edge to the TelegramChat entity by ID.
func (atc *AddressTrackerCreate) SetTelegramChatID(id int) *AddressTrackerCreate {
	atc.mutation.SetTelegramChatID(id)
	return atc
}

// SetNillableTelegramChatID sets the "telegram_chat" edge to the TelegramChat entity by ID if the given value is not nil.
func (atc *AddressTrackerCreate) SetNillableTelegramChatID(id *int) *AddressTrackerCreate {
	if id != nil {
		atc = atc.SetTelegramChatID(*id)
	}
	return atc
}

// SetTelegramChat sets the "telegram_chat" edge to the TelegramChat entity.
func (atc *AddressTrackerCreate) SetTelegramChat(t *TelegramChat) *AddressTrackerCreate {
	return atc.SetTelegramChatID(t.ID)
}

// AddChainProposalIDs adds the "chain_proposals" edge to the ChainProposal entity by IDs.
func (atc *AddressTrackerCreate) AddChainProposalIDs(ids ...int) *AddressTrackerCreate {
	atc.mutation.AddChainProposalIDs(ids...)
	return atc
}

// AddChainProposals adds the "chain_proposals" edges to the ChainProposal entity.
func (atc *AddressTrackerCreate) AddChainProposals(c ...*ChainProposal) *AddressTrackerCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return atc.AddChainProposalIDs(ids...)
}

// Mutation returns the AddressTrackerMutation object of the builder.
func (atc *AddressTrackerCreate) Mutation() *AddressTrackerMutation {
	return atc.mutation
}

// Save creates the AddressTracker in the database.
func (atc *AddressTrackerCreate) Save(ctx context.Context) (*AddressTracker, error) {
	atc.defaults()
	return withHooks[*AddressTracker, AddressTrackerMutation](ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *AddressTrackerCreate) SaveX(ctx context.Context) *AddressTracker {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *AddressTrackerCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *AddressTrackerCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *AddressTrackerCreate) defaults() {
	if _, ok := atc.mutation.CreateTime(); !ok {
		v := addresstracker.DefaultCreateTime()
		atc.mutation.SetCreateTime(v)
	}
	if _, ok := atc.mutation.UpdateTime(); !ok {
		v := addresstracker.DefaultUpdateTime()
		atc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *AddressTrackerCreate) check() error {
	if _, ok := atc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "AddressTracker.create_time"`)}
	}
	if _, ok := atc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "AddressTracker.update_time"`)}
	}
	if _, ok := atc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "AddressTracker.address"`)}
	}
	if _, ok := atc.mutation.ChainID(); !ok {
		return &ValidationError{Name: "chain", err: errors.New(`ent: missing required edge "AddressTracker.chain"`)}
	}
	return nil
}

func (atc *AddressTrackerCreate) sqlSave(ctx context.Context) (*AddressTracker, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *AddressTrackerCreate) createSpec() (*AddressTracker, *sqlgraph.CreateSpec) {
	var (
		_node = &AddressTracker{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(addresstracker.Table, sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt))
	)
	if value, ok := atc.mutation.CreateTime(); ok {
		_spec.SetField(addresstracker.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := atc.mutation.UpdateTime(); ok {
		_spec.SetField(addresstracker.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := atc.mutation.Address(); ok {
		_spec.SetField(addresstracker.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if nodes := atc.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.ChainTable,
			Columns: []string{addresstracker.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.chain_address_trackers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atc.mutation.DiscordChannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.DiscordChannelTable,
			Columns: []string{addresstracker.DiscordChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.discord_channel_address_trackers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atc.mutation.TelegramChatIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.TelegramChatTable,
			Columns: []string{addresstracker.TelegramChatColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.telegram_chat_address_trackers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atc.mutation.ChainProposalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addresstracker.ChainProposalsTable,
			Columns: addresstracker.ChainProposalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AddressTrackerCreateBulk is the builder for creating many AddressTracker entities in bulk.
type AddressTrackerCreateBulk struct {
	config
	builders []*AddressTrackerCreate
}

// Save creates the AddressTracker entities in the database.
func (atcb *AddressTrackerCreateBulk) Save(ctx context.Context) ([]*AddressTracker, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*AddressTracker, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AddressTrackerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *AddressTrackerCreateBulk) SaveX(ctx context.Context) []*AddressTracker {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *AddressTrackerCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *AddressTrackerCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}
