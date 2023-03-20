// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-notifier/ent/addresstracker"
	"github.com/shifty11/cosmos-notifier/ent/predicate"
)

// AddressTrackerDelete is the builder for deleting a AddressTracker entity.
type AddressTrackerDelete struct {
	config
	hooks    []Hook
	mutation *AddressTrackerMutation
}

// Where appends a list predicates to the AddressTrackerDelete builder.
func (atd *AddressTrackerDelete) Where(ps ...predicate.AddressTracker) *AddressTrackerDelete {
	atd.mutation.Where(ps...)
	return atd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atd *AddressTrackerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, AddressTrackerMutation](ctx, atd.sqlExec, atd.mutation, atd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (atd *AddressTrackerDelete) ExecX(ctx context.Context) int {
	n, err := atd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atd *AddressTrackerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(addresstracker.Table, sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt))
	if ps := atd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, atd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	atd.mutation.done = true
	return affected, err
}

// AddressTrackerDeleteOne is the builder for deleting a single AddressTracker entity.
type AddressTrackerDeleteOne struct {
	atd *AddressTrackerDelete
}

// Where appends a list predicates to the AddressTrackerDelete builder.
func (atdo *AddressTrackerDeleteOne) Where(ps ...predicate.AddressTracker) *AddressTrackerDeleteOne {
	atdo.atd.mutation.Where(ps...)
	return atdo
}

// Exec executes the deletion query.
func (atdo *AddressTrackerDeleteOne) Exec(ctx context.Context) error {
	n, err := atdo.atd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{addresstracker.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atdo *AddressTrackerDeleteOne) ExecX(ctx context.Context) {
	if err := atdo.Exec(ctx); err != nil {
		panic(err)
	}
}