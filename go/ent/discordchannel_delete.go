// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/predicate"
)

// DiscordChannelDelete is the builder for deleting a DiscordChannel entity.
type DiscordChannelDelete struct {
	config
	hooks    []Hook
	mutation *DiscordChannelMutation
}

// Where appends a list predicates to the DiscordChannelDelete builder.
func (dcd *DiscordChannelDelete) Where(ps ...predicate.DiscordChannel) *DiscordChannelDelete {
	dcd.mutation.Where(ps...)
	return dcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dcd *DiscordChannelDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dcd.hooks) == 0 {
		affected, err = dcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscordChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dcd.mutation = mutation
			affected, err = dcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dcd.hooks) - 1; i >= 0; i-- {
			if dcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcd *DiscordChannelDelete) ExecX(ctx context.Context) int {
	n, err := dcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dcd *DiscordChannelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: discordchannel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discordchannel.FieldID,
			},
		},
	}
	if ps := dcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DiscordChannelDeleteOne is the builder for deleting a single DiscordChannel entity.
type DiscordChannelDeleteOne struct {
	dcd *DiscordChannelDelete
}

// Exec executes the deletion query.
func (dcdo *DiscordChannelDeleteOne) Exec(ctx context.Context) error {
	n, err := dcdo.dcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{discordchannel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dcdo *DiscordChannelDeleteOne) ExecX(ctx context.Context) {
	dcdo.dcd.ExecX(ctx)
}
