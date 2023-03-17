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
	"github.com/shifty11/cosmos-notifier/ent/addresstracker"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/contract"
	"github.com/shifty11/cosmos-notifier/ent/predicate"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
	"github.com/shifty11/cosmos-notifier/ent/user"
)

// TelegramChatUpdate is the builder for updating TelegramChat entities.
type TelegramChatUpdate struct {
	config
	hooks    []Hook
	mutation *TelegramChatMutation
}

// Where appends a list predicates to the TelegramChatUpdate builder.
func (tcu *TelegramChatUpdate) Where(ps ...predicate.TelegramChat) *TelegramChatUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetUpdateTime sets the "update_time" field.
func (tcu *TelegramChatUpdate) SetUpdateTime(t time.Time) *TelegramChatUpdate {
	tcu.mutation.SetUpdateTime(t)
	return tcu
}

// SetChatID sets the "chat_id" field.
func (tcu *TelegramChatUpdate) SetChatID(i int64) *TelegramChatUpdate {
	tcu.mutation.ResetChatID()
	tcu.mutation.SetChatID(i)
	return tcu
}

// AddChatID adds i to the "chat_id" field.
func (tcu *TelegramChatUpdate) AddChatID(i int64) *TelegramChatUpdate {
	tcu.mutation.AddChatID(i)
	return tcu
}

// SetName sets the "name" field.
func (tcu *TelegramChatUpdate) SetName(s string) *TelegramChatUpdate {
	tcu.mutation.SetName(s)
	return tcu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (tcu *TelegramChatUpdate) AddUserIDs(ids ...int) *TelegramChatUpdate {
	tcu.mutation.AddUserIDs(ids...)
	return tcu
}

// AddUsers adds the "users" edges to the User entity.
func (tcu *TelegramChatUpdate) AddUsers(u ...*User) *TelegramChatUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tcu.AddUserIDs(ids...)
}

// AddContractIDs adds the "contracts" edge to the Contract entity by IDs.
func (tcu *TelegramChatUpdate) AddContractIDs(ids ...int) *TelegramChatUpdate {
	tcu.mutation.AddContractIDs(ids...)
	return tcu
}

// AddContracts adds the "contracts" edges to the Contract entity.
func (tcu *TelegramChatUpdate) AddContracts(c ...*Contract) *TelegramChatUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcu.AddContractIDs(ids...)
}

// AddChainIDs adds the "chains" edge to the Chain entity by IDs.
func (tcu *TelegramChatUpdate) AddChainIDs(ids ...int) *TelegramChatUpdate {
	tcu.mutation.AddChainIDs(ids...)
	return tcu
}

// AddChains adds the "chains" edges to the Chain entity.
func (tcu *TelegramChatUpdate) AddChains(c ...*Chain) *TelegramChatUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcu.AddChainIDs(ids...)
}

// AddAddressTrackerIDs adds the "address_trackers" edge to the AddressTracker entity by IDs.
func (tcu *TelegramChatUpdate) AddAddressTrackerIDs(ids ...int) *TelegramChatUpdate {
	tcu.mutation.AddAddressTrackerIDs(ids...)
	return tcu
}

// AddAddressTrackers adds the "address_trackers" edges to the AddressTracker entity.
func (tcu *TelegramChatUpdate) AddAddressTrackers(a ...*AddressTracker) *TelegramChatUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tcu.AddAddressTrackerIDs(ids...)
}

// Mutation returns the TelegramChatMutation object of the builder.
func (tcu *TelegramChatUpdate) Mutation() *TelegramChatMutation {
	return tcu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (tcu *TelegramChatUpdate) ClearUsers() *TelegramChatUpdate {
	tcu.mutation.ClearUsers()
	return tcu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (tcu *TelegramChatUpdate) RemoveUserIDs(ids ...int) *TelegramChatUpdate {
	tcu.mutation.RemoveUserIDs(ids...)
	return tcu
}

// RemoveUsers removes "users" edges to User entities.
func (tcu *TelegramChatUpdate) RemoveUsers(u ...*User) *TelegramChatUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tcu.RemoveUserIDs(ids...)
}

// ClearContracts clears all "contracts" edges to the Contract entity.
func (tcu *TelegramChatUpdate) ClearContracts() *TelegramChatUpdate {
	tcu.mutation.ClearContracts()
	return tcu
}

// RemoveContractIDs removes the "contracts" edge to Contract entities by IDs.
func (tcu *TelegramChatUpdate) RemoveContractIDs(ids ...int) *TelegramChatUpdate {
	tcu.mutation.RemoveContractIDs(ids...)
	return tcu
}

// RemoveContracts removes "contracts" edges to Contract entities.
func (tcu *TelegramChatUpdate) RemoveContracts(c ...*Contract) *TelegramChatUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcu.RemoveContractIDs(ids...)
}

// ClearChains clears all "chains" edges to the Chain entity.
func (tcu *TelegramChatUpdate) ClearChains() *TelegramChatUpdate {
	tcu.mutation.ClearChains()
	return tcu
}

// RemoveChainIDs removes the "chains" edge to Chain entities by IDs.
func (tcu *TelegramChatUpdate) RemoveChainIDs(ids ...int) *TelegramChatUpdate {
	tcu.mutation.RemoveChainIDs(ids...)
	return tcu
}

// RemoveChains removes "chains" edges to Chain entities.
func (tcu *TelegramChatUpdate) RemoveChains(c ...*Chain) *TelegramChatUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcu.RemoveChainIDs(ids...)
}

// ClearAddressTrackers clears all "address_trackers" edges to the AddressTracker entity.
func (tcu *TelegramChatUpdate) ClearAddressTrackers() *TelegramChatUpdate {
	tcu.mutation.ClearAddressTrackers()
	return tcu
}

// RemoveAddressTrackerIDs removes the "address_trackers" edge to AddressTracker entities by IDs.
func (tcu *TelegramChatUpdate) RemoveAddressTrackerIDs(ids ...int) *TelegramChatUpdate {
	tcu.mutation.RemoveAddressTrackerIDs(ids...)
	return tcu
}

// RemoveAddressTrackers removes "address_trackers" edges to AddressTracker entities.
func (tcu *TelegramChatUpdate) RemoveAddressTrackers(a ...*AddressTracker) *TelegramChatUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tcu.RemoveAddressTrackerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TelegramChatUpdate) Save(ctx context.Context) (int, error) {
	tcu.defaults()
	return withHooks[int, TelegramChatMutation](ctx, tcu.sqlSave, tcu.mutation, tcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TelegramChatUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TelegramChatUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TelegramChatUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcu *TelegramChatUpdate) defaults() {
	if _, ok := tcu.mutation.UpdateTime(); !ok {
		v := telegramchat.UpdateDefaultUpdateTime()
		tcu.mutation.SetUpdateTime(v)
	}
}

func (tcu *TelegramChatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(telegramchat.Table, telegramchat.Columns, sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt))
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.UpdateTime(); ok {
		_spec.SetField(telegramchat.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tcu.mutation.ChatID(); ok {
		_spec.SetField(telegramchat.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := tcu.mutation.AddedChatID(); ok {
		_spec.AddField(telegramchat.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := tcu.mutation.Name(); ok {
		_spec.SetField(telegramchat.FieldName, field.TypeString, value)
	}
	if tcu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.UsersTable,
			Columns: telegramchat.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !tcu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.UsersTable,
			Columns: telegramchat.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.UsersTable,
			Columns: telegramchat.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tcu.mutation.ContractsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ContractsTable,
			Columns: telegramchat.ContractsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.RemovedContractsIDs(); len(nodes) > 0 && !tcu.mutation.ContractsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ContractsTable,
			Columns: telegramchat.ContractsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.ContractsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ContractsTable,
			Columns: telegramchat.ContractsPrimaryKey,
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
	if tcu.mutation.ChainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ChainsTable,
			Columns: telegramchat.ChainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.RemovedChainsIDs(); len(nodes) > 0 && !tcu.mutation.ChainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ChainsTable,
			Columns: telegramchat.ChainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.ChainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ChainsTable,
			Columns: telegramchat.ChainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tcu.mutation.AddressTrackersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   telegramchat.AddressTrackersTable,
			Columns: []string{telegramchat.AddressTrackersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.RemovedAddressTrackersIDs(); len(nodes) > 0 && !tcu.mutation.AddressTrackersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   telegramchat.AddressTrackersTable,
			Columns: []string{telegramchat.AddressTrackersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.AddressTrackersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   telegramchat.AddressTrackersTable,
			Columns: []string{telegramchat.AddressTrackersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{telegramchat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tcu.mutation.done = true
	return n, nil
}

// TelegramChatUpdateOne is the builder for updating a single TelegramChat entity.
type TelegramChatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TelegramChatMutation
}

// SetUpdateTime sets the "update_time" field.
func (tcuo *TelegramChatUpdateOne) SetUpdateTime(t time.Time) *TelegramChatUpdateOne {
	tcuo.mutation.SetUpdateTime(t)
	return tcuo
}

// SetChatID sets the "chat_id" field.
func (tcuo *TelegramChatUpdateOne) SetChatID(i int64) *TelegramChatUpdateOne {
	tcuo.mutation.ResetChatID()
	tcuo.mutation.SetChatID(i)
	return tcuo
}

// AddChatID adds i to the "chat_id" field.
func (tcuo *TelegramChatUpdateOne) AddChatID(i int64) *TelegramChatUpdateOne {
	tcuo.mutation.AddChatID(i)
	return tcuo
}

// SetName sets the "name" field.
func (tcuo *TelegramChatUpdateOne) SetName(s string) *TelegramChatUpdateOne {
	tcuo.mutation.SetName(s)
	return tcuo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (tcuo *TelegramChatUpdateOne) AddUserIDs(ids ...int) *TelegramChatUpdateOne {
	tcuo.mutation.AddUserIDs(ids...)
	return tcuo
}

// AddUsers adds the "users" edges to the User entity.
func (tcuo *TelegramChatUpdateOne) AddUsers(u ...*User) *TelegramChatUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tcuo.AddUserIDs(ids...)
}

// AddContractIDs adds the "contracts" edge to the Contract entity by IDs.
func (tcuo *TelegramChatUpdateOne) AddContractIDs(ids ...int) *TelegramChatUpdateOne {
	tcuo.mutation.AddContractIDs(ids...)
	return tcuo
}

// AddContracts adds the "contracts" edges to the Contract entity.
func (tcuo *TelegramChatUpdateOne) AddContracts(c ...*Contract) *TelegramChatUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcuo.AddContractIDs(ids...)
}

// AddChainIDs adds the "chains" edge to the Chain entity by IDs.
func (tcuo *TelegramChatUpdateOne) AddChainIDs(ids ...int) *TelegramChatUpdateOne {
	tcuo.mutation.AddChainIDs(ids...)
	return tcuo
}

// AddChains adds the "chains" edges to the Chain entity.
func (tcuo *TelegramChatUpdateOne) AddChains(c ...*Chain) *TelegramChatUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcuo.AddChainIDs(ids...)
}

// AddAddressTrackerIDs adds the "address_trackers" edge to the AddressTracker entity by IDs.
func (tcuo *TelegramChatUpdateOne) AddAddressTrackerIDs(ids ...int) *TelegramChatUpdateOne {
	tcuo.mutation.AddAddressTrackerIDs(ids...)
	return tcuo
}

// AddAddressTrackers adds the "address_trackers" edges to the AddressTracker entity.
func (tcuo *TelegramChatUpdateOne) AddAddressTrackers(a ...*AddressTracker) *TelegramChatUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tcuo.AddAddressTrackerIDs(ids...)
}

// Mutation returns the TelegramChatMutation object of the builder.
func (tcuo *TelegramChatUpdateOne) Mutation() *TelegramChatMutation {
	return tcuo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (tcuo *TelegramChatUpdateOne) ClearUsers() *TelegramChatUpdateOne {
	tcuo.mutation.ClearUsers()
	return tcuo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (tcuo *TelegramChatUpdateOne) RemoveUserIDs(ids ...int) *TelegramChatUpdateOne {
	tcuo.mutation.RemoveUserIDs(ids...)
	return tcuo
}

// RemoveUsers removes "users" edges to User entities.
func (tcuo *TelegramChatUpdateOne) RemoveUsers(u ...*User) *TelegramChatUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tcuo.RemoveUserIDs(ids...)
}

// ClearContracts clears all "contracts" edges to the Contract entity.
func (tcuo *TelegramChatUpdateOne) ClearContracts() *TelegramChatUpdateOne {
	tcuo.mutation.ClearContracts()
	return tcuo
}

// RemoveContractIDs removes the "contracts" edge to Contract entities by IDs.
func (tcuo *TelegramChatUpdateOne) RemoveContractIDs(ids ...int) *TelegramChatUpdateOne {
	tcuo.mutation.RemoveContractIDs(ids...)
	return tcuo
}

// RemoveContracts removes "contracts" edges to Contract entities.
func (tcuo *TelegramChatUpdateOne) RemoveContracts(c ...*Contract) *TelegramChatUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcuo.RemoveContractIDs(ids...)
}

// ClearChains clears all "chains" edges to the Chain entity.
func (tcuo *TelegramChatUpdateOne) ClearChains() *TelegramChatUpdateOne {
	tcuo.mutation.ClearChains()
	return tcuo
}

// RemoveChainIDs removes the "chains" edge to Chain entities by IDs.
func (tcuo *TelegramChatUpdateOne) RemoveChainIDs(ids ...int) *TelegramChatUpdateOne {
	tcuo.mutation.RemoveChainIDs(ids...)
	return tcuo
}

// RemoveChains removes "chains" edges to Chain entities.
func (tcuo *TelegramChatUpdateOne) RemoveChains(c ...*Chain) *TelegramChatUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcuo.RemoveChainIDs(ids...)
}

// ClearAddressTrackers clears all "address_trackers" edges to the AddressTracker entity.
func (tcuo *TelegramChatUpdateOne) ClearAddressTrackers() *TelegramChatUpdateOne {
	tcuo.mutation.ClearAddressTrackers()
	return tcuo
}

// RemoveAddressTrackerIDs removes the "address_trackers" edge to AddressTracker entities by IDs.
func (tcuo *TelegramChatUpdateOne) RemoveAddressTrackerIDs(ids ...int) *TelegramChatUpdateOne {
	tcuo.mutation.RemoveAddressTrackerIDs(ids...)
	return tcuo
}

// RemoveAddressTrackers removes "address_trackers" edges to AddressTracker entities.
func (tcuo *TelegramChatUpdateOne) RemoveAddressTrackers(a ...*AddressTracker) *TelegramChatUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tcuo.RemoveAddressTrackerIDs(ids...)
}

// Where appends a list predicates to the TelegramChatUpdate builder.
func (tcuo *TelegramChatUpdateOne) Where(ps ...predicate.TelegramChat) *TelegramChatUpdateOne {
	tcuo.mutation.Where(ps...)
	return tcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TelegramChatUpdateOne) Select(field string, fields ...string) *TelegramChatUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TelegramChat entity.
func (tcuo *TelegramChatUpdateOne) Save(ctx context.Context) (*TelegramChat, error) {
	tcuo.defaults()
	return withHooks[*TelegramChat, TelegramChatMutation](ctx, tcuo.sqlSave, tcuo.mutation, tcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TelegramChatUpdateOne) SaveX(ctx context.Context) *TelegramChat {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TelegramChatUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TelegramChatUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcuo *TelegramChatUpdateOne) defaults() {
	if _, ok := tcuo.mutation.UpdateTime(); !ok {
		v := telegramchat.UpdateDefaultUpdateTime()
		tcuo.mutation.SetUpdateTime(v)
	}
}

func (tcuo *TelegramChatUpdateOne) sqlSave(ctx context.Context) (_node *TelegramChat, err error) {
	_spec := sqlgraph.NewUpdateSpec(telegramchat.Table, telegramchat.Columns, sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt))
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TelegramChat.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, telegramchat.FieldID)
		for _, f := range fields {
			if !telegramchat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != telegramchat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.UpdateTime(); ok {
		_spec.SetField(telegramchat.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tcuo.mutation.ChatID(); ok {
		_spec.SetField(telegramchat.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := tcuo.mutation.AddedChatID(); ok {
		_spec.AddField(telegramchat.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := tcuo.mutation.Name(); ok {
		_spec.SetField(telegramchat.FieldName, field.TypeString, value)
	}
	if tcuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.UsersTable,
			Columns: telegramchat.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !tcuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.UsersTable,
			Columns: telegramchat.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.UsersTable,
			Columns: telegramchat.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tcuo.mutation.ContractsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ContractsTable,
			Columns: telegramchat.ContractsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.RemovedContractsIDs(); len(nodes) > 0 && !tcuo.mutation.ContractsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ContractsTable,
			Columns: telegramchat.ContractsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.ContractsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ContractsTable,
			Columns: telegramchat.ContractsPrimaryKey,
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
	if tcuo.mutation.ChainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ChainsTable,
			Columns: telegramchat.ChainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.RemovedChainsIDs(); len(nodes) > 0 && !tcuo.mutation.ChainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ChainsTable,
			Columns: telegramchat.ChainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.ChainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ChainsTable,
			Columns: telegramchat.ChainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tcuo.mutation.AddressTrackersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   telegramchat.AddressTrackersTable,
			Columns: []string{telegramchat.AddressTrackersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.RemovedAddressTrackersIDs(); len(nodes) > 0 && !tcuo.mutation.AddressTrackersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   telegramchat.AddressTrackersTable,
			Columns: []string{telegramchat.AddressTrackersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.AddressTrackersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   telegramchat.AddressTrackersTable,
			Columns: []string{telegramchat.AddressTrackersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TelegramChat{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{telegramchat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tcuo.mutation.done = true
	return _node, nil
}
