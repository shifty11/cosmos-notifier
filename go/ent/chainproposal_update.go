// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/dao-dao-notifier/ent/chain"
	"github.com/shifty11/dao-dao-notifier/ent/chainproposal"
	"github.com/shifty11/dao-dao-notifier/ent/predicate"
)

// ChainProposalUpdate is the builder for updating ChainProposal entities.
type ChainProposalUpdate struct {
	config
	hooks    []Hook
	mutation *ChainProposalMutation
}

// Where appends a list predicates to the ChainProposalUpdate builder.
func (cpu *ChainProposalUpdate) Where(ps ...predicate.ChainProposal) *ChainProposalUpdate {
	cpu.mutation.Where(ps...)
	return cpu
}

// SetUpdateTime sets the "update_time" field.
func (cpu *ChainProposalUpdate) SetUpdateTime(t time.Time) *ChainProposalUpdate {
	cpu.mutation.SetUpdateTime(t)
	return cpu
}

// SetProposalID sets the "proposal_id" field.
func (cpu *ChainProposalUpdate) SetProposalID(i int) *ChainProposalUpdate {
	cpu.mutation.ResetProposalID()
	cpu.mutation.SetProposalID(i)
	return cpu
}

// AddProposalID adds i to the "proposal_id" field.
func (cpu *ChainProposalUpdate) AddProposalID(i int) *ChainProposalUpdate {
	cpu.mutation.AddProposalID(i)
	return cpu
}

// SetTitle sets the "title" field.
func (cpu *ChainProposalUpdate) SetTitle(s string) *ChainProposalUpdate {
	cpu.mutation.SetTitle(s)
	return cpu
}

// SetDescription sets the "description" field.
func (cpu *ChainProposalUpdate) SetDescription(s string) *ChainProposalUpdate {
	cpu.mutation.SetDescription(s)
	return cpu
}

// SetVotingStartTime sets the "voting_start_time" field.
func (cpu *ChainProposalUpdate) SetVotingStartTime(t time.Time) *ChainProposalUpdate {
	cpu.mutation.SetVotingStartTime(t)
	return cpu
}

// SetVotingEndTime sets the "voting_end_time" field.
func (cpu *ChainProposalUpdate) SetVotingEndTime(t time.Time) *ChainProposalUpdate {
	cpu.mutation.SetVotingEndTime(t)
	return cpu
}

// SetStatus sets the "status" field.
func (cpu *ChainProposalUpdate) SetStatus(c chainproposal.Status) *ChainProposalUpdate {
	cpu.mutation.SetStatus(c)
	return cpu
}

// SetChainID sets the "chain" edge to the Chain entity by ID.
func (cpu *ChainProposalUpdate) SetChainID(id int) *ChainProposalUpdate {
	cpu.mutation.SetChainID(id)
	return cpu
}

// SetNillableChainID sets the "chain" edge to the Chain entity by ID if the given value is not nil.
func (cpu *ChainProposalUpdate) SetNillableChainID(id *int) *ChainProposalUpdate {
	if id != nil {
		cpu = cpu.SetChainID(*id)
	}
	return cpu
}

// SetChain sets the "chain" edge to the Chain entity.
func (cpu *ChainProposalUpdate) SetChain(c *Chain) *ChainProposalUpdate {
	return cpu.SetChainID(c.ID)
}

// Mutation returns the ChainProposalMutation object of the builder.
func (cpu *ChainProposalUpdate) Mutation() *ChainProposalMutation {
	return cpu.mutation
}

// ClearChain clears the "chain" edge to the Chain entity.
func (cpu *ChainProposalUpdate) ClearChain() *ChainProposalUpdate {
	cpu.mutation.ClearChain()
	return cpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cpu *ChainProposalUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cpu.defaults()
	if len(cpu.hooks) == 0 {
		if err = cpu.check(); err != nil {
			return 0, err
		}
		affected, err = cpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChainProposalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpu.check(); err != nil {
				return 0, err
			}
			cpu.mutation = mutation
			affected, err = cpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cpu.hooks) - 1; i >= 0; i-- {
			if cpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cpu *ChainProposalUpdate) SaveX(ctx context.Context) int {
	affected, err := cpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cpu *ChainProposalUpdate) Exec(ctx context.Context) error {
	_, err := cpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpu *ChainProposalUpdate) ExecX(ctx context.Context) {
	if err := cpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpu *ChainProposalUpdate) defaults() {
	if _, ok := cpu.mutation.UpdateTime(); !ok {
		v := chainproposal.UpdateDefaultUpdateTime()
		cpu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpu *ChainProposalUpdate) check() error {
	if v, ok := cpu.mutation.Status(); ok {
		if err := chainproposal.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ChainProposal.status": %w`, err)}
		}
	}
	return nil
}

func (cpu *ChainProposalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chainproposal.Table,
			Columns: chainproposal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chainproposal.FieldID,
			},
		},
	}
	if ps := cpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chainproposal.FieldUpdateTime,
		})
	}
	if value, ok := cpu.mutation.ProposalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: chainproposal.FieldProposalID,
		})
	}
	if value, ok := cpu.mutation.AddedProposalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: chainproposal.FieldProposalID,
		})
	}
	if value, ok := cpu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chainproposal.FieldTitle,
		})
	}
	if value, ok := cpu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chainproposal.FieldDescription,
		})
	}
	if value, ok := cpu.mutation.VotingStartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chainproposal.FieldVotingStartTime,
		})
	}
	if value, ok := cpu.mutation.VotingEndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chainproposal.FieldVotingEndTime,
		})
	}
	if value, ok := cpu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: chainproposal.FieldStatus,
		})
	}
	if cpu.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chainproposal.ChainTable,
			Columns: []string{chainproposal.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chain.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chainproposal.ChainTable,
			Columns: []string{chainproposal.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chainproposal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ChainProposalUpdateOne is the builder for updating a single ChainProposal entity.
type ChainProposalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChainProposalMutation
}

// SetUpdateTime sets the "update_time" field.
func (cpuo *ChainProposalUpdateOne) SetUpdateTime(t time.Time) *ChainProposalUpdateOne {
	cpuo.mutation.SetUpdateTime(t)
	return cpuo
}

// SetProposalID sets the "proposal_id" field.
func (cpuo *ChainProposalUpdateOne) SetProposalID(i int) *ChainProposalUpdateOne {
	cpuo.mutation.ResetProposalID()
	cpuo.mutation.SetProposalID(i)
	return cpuo
}

// AddProposalID adds i to the "proposal_id" field.
func (cpuo *ChainProposalUpdateOne) AddProposalID(i int) *ChainProposalUpdateOne {
	cpuo.mutation.AddProposalID(i)
	return cpuo
}

// SetTitle sets the "title" field.
func (cpuo *ChainProposalUpdateOne) SetTitle(s string) *ChainProposalUpdateOne {
	cpuo.mutation.SetTitle(s)
	return cpuo
}

// SetDescription sets the "description" field.
func (cpuo *ChainProposalUpdateOne) SetDescription(s string) *ChainProposalUpdateOne {
	cpuo.mutation.SetDescription(s)
	return cpuo
}

// SetVotingStartTime sets the "voting_start_time" field.
func (cpuo *ChainProposalUpdateOne) SetVotingStartTime(t time.Time) *ChainProposalUpdateOne {
	cpuo.mutation.SetVotingStartTime(t)
	return cpuo
}

// SetVotingEndTime sets the "voting_end_time" field.
func (cpuo *ChainProposalUpdateOne) SetVotingEndTime(t time.Time) *ChainProposalUpdateOne {
	cpuo.mutation.SetVotingEndTime(t)
	return cpuo
}

// SetStatus sets the "status" field.
func (cpuo *ChainProposalUpdateOne) SetStatus(c chainproposal.Status) *ChainProposalUpdateOne {
	cpuo.mutation.SetStatus(c)
	return cpuo
}

// SetChainID sets the "chain" edge to the Chain entity by ID.
func (cpuo *ChainProposalUpdateOne) SetChainID(id int) *ChainProposalUpdateOne {
	cpuo.mutation.SetChainID(id)
	return cpuo
}

// SetNillableChainID sets the "chain" edge to the Chain entity by ID if the given value is not nil.
func (cpuo *ChainProposalUpdateOne) SetNillableChainID(id *int) *ChainProposalUpdateOne {
	if id != nil {
		cpuo = cpuo.SetChainID(*id)
	}
	return cpuo
}

// SetChain sets the "chain" edge to the Chain entity.
func (cpuo *ChainProposalUpdateOne) SetChain(c *Chain) *ChainProposalUpdateOne {
	return cpuo.SetChainID(c.ID)
}

// Mutation returns the ChainProposalMutation object of the builder.
func (cpuo *ChainProposalUpdateOne) Mutation() *ChainProposalMutation {
	return cpuo.mutation
}

// ClearChain clears the "chain" edge to the Chain entity.
func (cpuo *ChainProposalUpdateOne) ClearChain() *ChainProposalUpdateOne {
	cpuo.mutation.ClearChain()
	return cpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cpuo *ChainProposalUpdateOne) Select(field string, fields ...string) *ChainProposalUpdateOne {
	cpuo.fields = append([]string{field}, fields...)
	return cpuo
}

// Save executes the query and returns the updated ChainProposal entity.
func (cpuo *ChainProposalUpdateOne) Save(ctx context.Context) (*ChainProposal, error) {
	var (
		err  error
		node *ChainProposal
	)
	cpuo.defaults()
	if len(cpuo.hooks) == 0 {
		if err = cpuo.check(); err != nil {
			return nil, err
		}
		node, err = cpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChainProposalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpuo.check(); err != nil {
				return nil, err
			}
			cpuo.mutation = mutation
			node, err = cpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cpuo.hooks) - 1; i >= 0; i-- {
			if cpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cpuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ChainProposal)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ChainProposalMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cpuo *ChainProposalUpdateOne) SaveX(ctx context.Context) *ChainProposal {
	node, err := cpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cpuo *ChainProposalUpdateOne) Exec(ctx context.Context) error {
	_, err := cpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpuo *ChainProposalUpdateOne) ExecX(ctx context.Context) {
	if err := cpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpuo *ChainProposalUpdateOne) defaults() {
	if _, ok := cpuo.mutation.UpdateTime(); !ok {
		v := chainproposal.UpdateDefaultUpdateTime()
		cpuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpuo *ChainProposalUpdateOne) check() error {
	if v, ok := cpuo.mutation.Status(); ok {
		if err := chainproposal.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ChainProposal.status": %w`, err)}
		}
	}
	return nil
}

func (cpuo *ChainProposalUpdateOne) sqlSave(ctx context.Context) (_node *ChainProposal, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chainproposal.Table,
			Columns: chainproposal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chainproposal.FieldID,
			},
		},
	}
	id, ok := cpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ChainProposal.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chainproposal.FieldID)
		for _, f := range fields {
			if !chainproposal.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chainproposal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chainproposal.FieldUpdateTime,
		})
	}
	if value, ok := cpuo.mutation.ProposalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: chainproposal.FieldProposalID,
		})
	}
	if value, ok := cpuo.mutation.AddedProposalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: chainproposal.FieldProposalID,
		})
	}
	if value, ok := cpuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chainproposal.FieldTitle,
		})
	}
	if value, ok := cpuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chainproposal.FieldDescription,
		})
	}
	if value, ok := cpuo.mutation.VotingStartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chainproposal.FieldVotingStartTime,
		})
	}
	if value, ok := cpuo.mutation.VotingEndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chainproposal.FieldVotingEndTime,
		})
	}
	if value, ok := cpuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: chainproposal.FieldStatus,
		})
	}
	if cpuo.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chainproposal.ChainTable,
			Columns: []string{chainproposal.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chain.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chainproposal.ChainTable,
			Columns: []string{chainproposal.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ChainProposal{config: cpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chainproposal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
