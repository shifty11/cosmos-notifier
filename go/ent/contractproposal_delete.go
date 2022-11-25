// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-notifier/ent/contractproposal"
	"github.com/shifty11/cosmos-notifier/ent/predicate"
)

// ContractProposalDelete is the builder for deleting a ContractProposal entity.
type ContractProposalDelete struct {
	config
	hooks    []Hook
	mutation *ContractProposalMutation
}

// Where appends a list predicates to the ContractProposalDelete builder.
func (cpd *ContractProposalDelete) Where(ps ...predicate.ContractProposal) *ContractProposalDelete {
	cpd.mutation.Where(ps...)
	return cpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cpd *ContractProposalDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cpd.hooks) == 0 {
		affected, err = cpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContractProposalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cpd.mutation = mutation
			affected, err = cpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cpd.hooks) - 1; i >= 0; i-- {
			if cpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpd *ContractProposalDelete) ExecX(ctx context.Context) int {
	n, err := cpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cpd *ContractProposalDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: contractproposal.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contractproposal.FieldID,
			},
		},
	}
	if ps := cpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ContractProposalDeleteOne is the builder for deleting a single ContractProposal entity.
type ContractProposalDeleteOne struct {
	cpd *ContractProposalDelete
}

// Exec executes the deletion query.
func (cpdo *ContractProposalDeleteOne) Exec(ctx context.Context) error {
	n, err := cpdo.cpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{contractproposal.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cpdo *ContractProposalDeleteOne) ExecX(ctx context.Context) {
	cpdo.cpd.ExecX(ctx)
}
