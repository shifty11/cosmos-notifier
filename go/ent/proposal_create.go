// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/ent/proposal"
)

// ProposalCreate is the builder for creating a Proposal entity.
type ProposalCreate struct {
	config
	mutation *ProposalMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *ProposalCreate) SetCreateTime(t time.Time) *ProposalCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *ProposalCreate) SetNillableCreateTime(t *time.Time) *ProposalCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *ProposalCreate) SetUpdateTime(t time.Time) *ProposalCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *ProposalCreate) SetNillableUpdateTime(t *time.Time) *ProposalCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetProposalID sets the "proposal_id" field.
func (pc *ProposalCreate) SetProposalID(s string) *ProposalCreate {
	pc.mutation.SetProposalID(s)
	return pc
}

// SetTitle sets the "title" field.
func (pc *ProposalCreate) SetTitle(s string) *ProposalCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProposalCreate) SetDescription(s string) *ProposalCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetExpiresAt sets the "expires_at" field.
func (pc *ProposalCreate) SetExpiresAt(t time.Time) *ProposalCreate {
	pc.mutation.SetExpiresAt(t)
	return pc
}

// SetStatus sets the "status" field.
func (pc *ProposalCreate) SetStatus(pr proposal.Status) *ProposalCreate {
	pc.mutation.SetStatus(pr)
	return pc
}

// SetContractID sets the "contract" edge to the Contract entity by ID.
func (pc *ProposalCreate) SetContractID(id int) *ProposalCreate {
	pc.mutation.SetContractID(id)
	return pc
}

// SetNillableContractID sets the "contract" edge to the Contract entity by ID if the given value is not nil.
func (pc *ProposalCreate) SetNillableContractID(id *int) *ProposalCreate {
	if id != nil {
		pc = pc.SetContractID(*id)
	}
	return pc
}

// SetContract sets the "contract" edge to the Contract entity.
func (pc *ProposalCreate) SetContract(c *Contract) *ProposalCreate {
	return pc.SetContractID(c.ID)
}

// Mutation returns the ProposalMutation object of the builder.
func (pc *ProposalCreate) Mutation() *ProposalMutation {
	return pc.mutation
}

// Save creates the Proposal in the database.
func (pc *ProposalCreate) Save(ctx context.Context) (*Proposal, error) {
	var (
		err  error
		node *Proposal
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProposalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Proposal)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProposalMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProposalCreate) SaveX(ctx context.Context) *Proposal {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProposalCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProposalCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProposalCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := proposal.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := proposal.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProposalCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Proposal.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Proposal.update_time"`)}
	}
	if _, ok := pc.mutation.ProposalID(); !ok {
		return &ValidationError{Name: "proposal_id", err: errors.New(`ent: missing required field "Proposal.proposal_id"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Proposal.title"`)}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Proposal.description"`)}
	}
	if _, ok := pc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "Proposal.expires_at"`)}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Proposal.status"`)}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := proposal.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Proposal.status": %w`, err)}
		}
	}
	return nil
}

func (pc *ProposalCreate) sqlSave(ctx context.Context) (*Proposal, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *ProposalCreate) createSpec() (*Proposal, *sqlgraph.CreateSpec) {
	var (
		_node = &Proposal{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: proposal.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: proposal.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: proposal.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: proposal.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.ProposalID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: proposal.FieldProposalID,
		})
		_node.ProposalID = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: proposal.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: proposal.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := pc.mutation.ExpiresAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: proposal.FieldExpiresAt,
		})
		_node.ExpiresAt = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: proposal.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := pc.mutation.ContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proposal.ContractTable,
			Columns: []string{proposal.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contract.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.contract_proposals = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProposalCreateBulk is the builder for creating many Proposal entities in bulk.
type ProposalCreateBulk struct {
	config
	builders []*ProposalCreate
}

// Save creates the Proposal entities in the database.
func (pcb *ProposalCreateBulk) Save(ctx context.Context) ([]*Proposal, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Proposal, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProposalMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProposalCreateBulk) SaveX(ctx context.Context) []*Proposal {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProposalCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProposalCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
