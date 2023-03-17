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
	"github.com/shifty11/cosmos-notifier/ent/contract"
	"github.com/shifty11/cosmos-notifier/ent/contractproposal"
	"github.com/shifty11/cosmos-notifier/ent/predicate"
)

// ContractProposalUpdate is the builder for updating ContractProposal entities.
type ContractProposalUpdate struct {
	config
	hooks    []Hook
	mutation *ContractProposalMutation
}

// Where appends a list predicates to the ContractProposalUpdate builder.
func (cpu *ContractProposalUpdate) Where(ps ...predicate.ContractProposal) *ContractProposalUpdate {
	cpu.mutation.Where(ps...)
	return cpu
}

// SetUpdateTime sets the "update_time" field.
func (cpu *ContractProposalUpdate) SetUpdateTime(t time.Time) *ContractProposalUpdate {
	cpu.mutation.SetUpdateTime(t)
	return cpu
}

// SetProposalID sets the "proposal_id" field.
func (cpu *ContractProposalUpdate) SetProposalID(i int) *ContractProposalUpdate {
	cpu.mutation.ResetProposalID()
	cpu.mutation.SetProposalID(i)
	return cpu
}

// AddProposalID adds i to the "proposal_id" field.
func (cpu *ContractProposalUpdate) AddProposalID(i int) *ContractProposalUpdate {
	cpu.mutation.AddProposalID(i)
	return cpu
}

// SetTitle sets the "title" field.
func (cpu *ContractProposalUpdate) SetTitle(s string) *ContractProposalUpdate {
	cpu.mutation.SetTitle(s)
	return cpu
}

// SetDescription sets the "description" field.
func (cpu *ContractProposalUpdate) SetDescription(s string) *ContractProposalUpdate {
	cpu.mutation.SetDescription(s)
	return cpu
}

// SetExpiresAt sets the "expires_at" field.
func (cpu *ContractProposalUpdate) SetExpiresAt(t time.Time) *ContractProposalUpdate {
	cpu.mutation.SetExpiresAt(t)
	return cpu
}

// SetStatus sets the "status" field.
func (cpu *ContractProposalUpdate) SetStatus(c contractproposal.Status) *ContractProposalUpdate {
	cpu.mutation.SetStatus(c)
	return cpu
}

// SetContractID sets the "contract" edge to the Contract entity by ID.
func (cpu *ContractProposalUpdate) SetContractID(id int) *ContractProposalUpdate {
	cpu.mutation.SetContractID(id)
	return cpu
}

// SetNillableContractID sets the "contract" edge to the Contract entity by ID if the given value is not nil.
func (cpu *ContractProposalUpdate) SetNillableContractID(id *int) *ContractProposalUpdate {
	if id != nil {
		cpu = cpu.SetContractID(*id)
	}
	return cpu
}

// SetContract sets the "contract" edge to the Contract entity.
func (cpu *ContractProposalUpdate) SetContract(c *Contract) *ContractProposalUpdate {
	return cpu.SetContractID(c.ID)
}

// Mutation returns the ContractProposalMutation object of the builder.
func (cpu *ContractProposalUpdate) Mutation() *ContractProposalMutation {
	return cpu.mutation
}

// ClearContract clears the "contract" edge to the Contract entity.
func (cpu *ContractProposalUpdate) ClearContract() *ContractProposalUpdate {
	cpu.mutation.ClearContract()
	return cpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cpu *ContractProposalUpdate) Save(ctx context.Context) (int, error) {
	cpu.defaults()
	return withHooks[int, ContractProposalMutation](ctx, cpu.sqlSave, cpu.mutation, cpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cpu *ContractProposalUpdate) SaveX(ctx context.Context) int {
	affected, err := cpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cpu *ContractProposalUpdate) Exec(ctx context.Context) error {
	_, err := cpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpu *ContractProposalUpdate) ExecX(ctx context.Context) {
	if err := cpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpu *ContractProposalUpdate) defaults() {
	if _, ok := cpu.mutation.UpdateTime(); !ok {
		v := contractproposal.UpdateDefaultUpdateTime()
		cpu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpu *ContractProposalUpdate) check() error {
	if v, ok := cpu.mutation.Status(); ok {
		if err := contractproposal.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ContractProposal.status": %w`, err)}
		}
	}
	return nil
}

func (cpu *ContractProposalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(contractproposal.Table, contractproposal.Columns, sqlgraph.NewFieldSpec(contractproposal.FieldID, field.TypeInt))
	if ps := cpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpu.mutation.UpdateTime(); ok {
		_spec.SetField(contractproposal.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cpu.mutation.ProposalID(); ok {
		_spec.SetField(contractproposal.FieldProposalID, field.TypeInt, value)
	}
	if value, ok := cpu.mutation.AddedProposalID(); ok {
		_spec.AddField(contractproposal.FieldProposalID, field.TypeInt, value)
	}
	if value, ok := cpu.mutation.Title(); ok {
		_spec.SetField(contractproposal.FieldTitle, field.TypeString, value)
	}
	if value, ok := cpu.mutation.Description(); ok {
		_spec.SetField(contractproposal.FieldDescription, field.TypeString, value)
	}
	if value, ok := cpu.mutation.ExpiresAt(); ok {
		_spec.SetField(contractproposal.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := cpu.mutation.Status(); ok {
		_spec.SetField(contractproposal.FieldStatus, field.TypeEnum, value)
	}
	if cpu.mutation.ContractCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contractproposal.ContractTable,
			Columns: []string{contractproposal.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.ContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contractproposal.ContractTable,
			Columns: []string{contractproposal.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contractproposal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cpu.mutation.done = true
	return n, nil
}

// ContractProposalUpdateOne is the builder for updating a single ContractProposal entity.
type ContractProposalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContractProposalMutation
}

// SetUpdateTime sets the "update_time" field.
func (cpuo *ContractProposalUpdateOne) SetUpdateTime(t time.Time) *ContractProposalUpdateOne {
	cpuo.mutation.SetUpdateTime(t)
	return cpuo
}

// SetProposalID sets the "proposal_id" field.
func (cpuo *ContractProposalUpdateOne) SetProposalID(i int) *ContractProposalUpdateOne {
	cpuo.mutation.ResetProposalID()
	cpuo.mutation.SetProposalID(i)
	return cpuo
}

// AddProposalID adds i to the "proposal_id" field.
func (cpuo *ContractProposalUpdateOne) AddProposalID(i int) *ContractProposalUpdateOne {
	cpuo.mutation.AddProposalID(i)
	return cpuo
}

// SetTitle sets the "title" field.
func (cpuo *ContractProposalUpdateOne) SetTitle(s string) *ContractProposalUpdateOne {
	cpuo.mutation.SetTitle(s)
	return cpuo
}

// SetDescription sets the "description" field.
func (cpuo *ContractProposalUpdateOne) SetDescription(s string) *ContractProposalUpdateOne {
	cpuo.mutation.SetDescription(s)
	return cpuo
}

// SetExpiresAt sets the "expires_at" field.
func (cpuo *ContractProposalUpdateOne) SetExpiresAt(t time.Time) *ContractProposalUpdateOne {
	cpuo.mutation.SetExpiresAt(t)
	return cpuo
}

// SetStatus sets the "status" field.
func (cpuo *ContractProposalUpdateOne) SetStatus(c contractproposal.Status) *ContractProposalUpdateOne {
	cpuo.mutation.SetStatus(c)
	return cpuo
}

// SetContractID sets the "contract" edge to the Contract entity by ID.
func (cpuo *ContractProposalUpdateOne) SetContractID(id int) *ContractProposalUpdateOne {
	cpuo.mutation.SetContractID(id)
	return cpuo
}

// SetNillableContractID sets the "contract" edge to the Contract entity by ID if the given value is not nil.
func (cpuo *ContractProposalUpdateOne) SetNillableContractID(id *int) *ContractProposalUpdateOne {
	if id != nil {
		cpuo = cpuo.SetContractID(*id)
	}
	return cpuo
}

// SetContract sets the "contract" edge to the Contract entity.
func (cpuo *ContractProposalUpdateOne) SetContract(c *Contract) *ContractProposalUpdateOne {
	return cpuo.SetContractID(c.ID)
}

// Mutation returns the ContractProposalMutation object of the builder.
func (cpuo *ContractProposalUpdateOne) Mutation() *ContractProposalMutation {
	return cpuo.mutation
}

// ClearContract clears the "contract" edge to the Contract entity.
func (cpuo *ContractProposalUpdateOne) ClearContract() *ContractProposalUpdateOne {
	cpuo.mutation.ClearContract()
	return cpuo
}

// Where appends a list predicates to the ContractProposalUpdate builder.
func (cpuo *ContractProposalUpdateOne) Where(ps ...predicate.ContractProposal) *ContractProposalUpdateOne {
	cpuo.mutation.Where(ps...)
	return cpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cpuo *ContractProposalUpdateOne) Select(field string, fields ...string) *ContractProposalUpdateOne {
	cpuo.fields = append([]string{field}, fields...)
	return cpuo
}

// Save executes the query and returns the updated ContractProposal entity.
func (cpuo *ContractProposalUpdateOne) Save(ctx context.Context) (*ContractProposal, error) {
	cpuo.defaults()
	return withHooks[*ContractProposal, ContractProposalMutation](ctx, cpuo.sqlSave, cpuo.mutation, cpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cpuo *ContractProposalUpdateOne) SaveX(ctx context.Context) *ContractProposal {
	node, err := cpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cpuo *ContractProposalUpdateOne) Exec(ctx context.Context) error {
	_, err := cpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpuo *ContractProposalUpdateOne) ExecX(ctx context.Context) {
	if err := cpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpuo *ContractProposalUpdateOne) defaults() {
	if _, ok := cpuo.mutation.UpdateTime(); !ok {
		v := contractproposal.UpdateDefaultUpdateTime()
		cpuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpuo *ContractProposalUpdateOne) check() error {
	if v, ok := cpuo.mutation.Status(); ok {
		if err := contractproposal.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ContractProposal.status": %w`, err)}
		}
	}
	return nil
}

func (cpuo *ContractProposalUpdateOne) sqlSave(ctx context.Context) (_node *ContractProposal, err error) {
	if err := cpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(contractproposal.Table, contractproposal.Columns, sqlgraph.NewFieldSpec(contractproposal.FieldID, field.TypeInt))
	id, ok := cpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ContractProposal.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contractproposal.FieldID)
		for _, f := range fields {
			if !contractproposal.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contractproposal.FieldID {
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
		_spec.SetField(contractproposal.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cpuo.mutation.ProposalID(); ok {
		_spec.SetField(contractproposal.FieldProposalID, field.TypeInt, value)
	}
	if value, ok := cpuo.mutation.AddedProposalID(); ok {
		_spec.AddField(contractproposal.FieldProposalID, field.TypeInt, value)
	}
	if value, ok := cpuo.mutation.Title(); ok {
		_spec.SetField(contractproposal.FieldTitle, field.TypeString, value)
	}
	if value, ok := cpuo.mutation.Description(); ok {
		_spec.SetField(contractproposal.FieldDescription, field.TypeString, value)
	}
	if value, ok := cpuo.mutation.ExpiresAt(); ok {
		_spec.SetField(contractproposal.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := cpuo.mutation.Status(); ok {
		_spec.SetField(contractproposal.FieldStatus, field.TypeEnum, value)
	}
	if cpuo.mutation.ContractCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contractproposal.ContractTable,
			Columns: []string{contractproposal.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.ContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contractproposal.ContractTable,
			Columns: []string{contractproposal.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ContractProposal{config: cpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contractproposal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cpuo.mutation.done = true
	return _node, nil
}
