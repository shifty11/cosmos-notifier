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
	"github.com/shifty11/cosmos-notifier/ent/chainproposal"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/predicate"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
	"github.com/shifty11/cosmos-notifier/ent/validator"
)

// ChainUpdate is the builder for updating Chain entities.
type ChainUpdate struct {
	config
	hooks    []Hook
	mutation *ChainMutation
}

// Where appends a list predicates to the ChainUpdate builder.
func (cu *ChainUpdate) Where(ps ...predicate.Chain) *ChainUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdateTime sets the "update_time" field.
func (cu *ChainUpdate) SetUpdateTime(t time.Time) *ChainUpdate {
	cu.mutation.SetUpdateTime(t)
	return cu
}

// SetChainID sets the "chain_id" field.
func (cu *ChainUpdate) SetChainID(s string) *ChainUpdate {
	cu.mutation.SetChainID(s)
	return cu
}

// SetName sets the "name" field.
func (cu *ChainUpdate) SetName(s string) *ChainUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetPrettyName sets the "pretty_name" field.
func (cu *ChainUpdate) SetPrettyName(s string) *ChainUpdate {
	cu.mutation.SetPrettyName(s)
	return cu
}

// SetPath sets the "path" field.
func (cu *ChainUpdate) SetPath(s string) *ChainUpdate {
	cu.mutation.SetPath(s)
	return cu
}

// SetDisplay sets the "display" field.
func (cu *ChainUpdate) SetDisplay(s string) *ChainUpdate {
	cu.mutation.SetDisplay(s)
	return cu
}

// SetNillableDisplay sets the "display" field if the given value is not nil.
func (cu *ChainUpdate) SetNillableDisplay(s *string) *ChainUpdate {
	if s != nil {
		cu.SetDisplay(*s)
	}
	return cu
}

// SetIsEnabled sets the "is_enabled" field.
func (cu *ChainUpdate) SetIsEnabled(b bool) *ChainUpdate {
	cu.mutation.SetIsEnabled(b)
	return cu
}

// SetNillableIsEnabled sets the "is_enabled" field if the given value is not nil.
func (cu *ChainUpdate) SetNillableIsEnabled(b *bool) *ChainUpdate {
	if b != nil {
		cu.SetIsEnabled(*b)
	}
	return cu
}

// SetImageURL sets the "image_url" field.
func (cu *ChainUpdate) SetImageURL(s string) *ChainUpdate {
	cu.mutation.SetImageURL(s)
	return cu
}

// SetThumbnailURL sets the "thumbnail_url" field.
func (cu *ChainUpdate) SetThumbnailURL(s string) *ChainUpdate {
	cu.mutation.SetThumbnailURL(s)
	return cu
}

// SetNillableThumbnailURL sets the "thumbnail_url" field if the given value is not nil.
func (cu *ChainUpdate) SetNillableThumbnailURL(s *string) *ChainUpdate {
	if s != nil {
		cu.SetThumbnailURL(*s)
	}
	return cu
}

// SetBech32Prefix sets the "bech32_prefix" field.
func (cu *ChainUpdate) SetBech32Prefix(s string) *ChainUpdate {
	cu.mutation.SetBech32Prefix(s)
	return cu
}

// AddChainProposalIDs adds the "chain_proposals" edge to the ChainProposal entity by IDs.
func (cu *ChainUpdate) AddChainProposalIDs(ids ...int) *ChainUpdate {
	cu.mutation.AddChainProposalIDs(ids...)
	return cu
}

// AddChainProposals adds the "chain_proposals" edges to the ChainProposal entity.
func (cu *ChainUpdate) AddChainProposals(c ...*ChainProposal) *ChainUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddChainProposalIDs(ids...)
}

// AddTelegramChatIDs adds the "telegram_chats" edge to the TelegramChat entity by IDs.
func (cu *ChainUpdate) AddTelegramChatIDs(ids ...int) *ChainUpdate {
	cu.mutation.AddTelegramChatIDs(ids...)
	return cu
}

// AddTelegramChats adds the "telegram_chats" edges to the TelegramChat entity.
func (cu *ChainUpdate) AddTelegramChats(t ...*TelegramChat) *ChainUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddTelegramChatIDs(ids...)
}

// AddDiscordChannelIDs adds the "discord_channels" edge to the DiscordChannel entity by IDs.
func (cu *ChainUpdate) AddDiscordChannelIDs(ids ...int) *ChainUpdate {
	cu.mutation.AddDiscordChannelIDs(ids...)
	return cu
}

// AddDiscordChannels adds the "discord_channels" edges to the DiscordChannel entity.
func (cu *ChainUpdate) AddDiscordChannels(d ...*DiscordChannel) *ChainUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cu.AddDiscordChannelIDs(ids...)
}

// AddAddressTrackerIDs adds the "address_trackers" edge to the AddressTracker entity by IDs.
func (cu *ChainUpdate) AddAddressTrackerIDs(ids ...int) *ChainUpdate {
	cu.mutation.AddAddressTrackerIDs(ids...)
	return cu
}

// AddAddressTrackers adds the "address_trackers" edges to the AddressTracker entity.
func (cu *ChainUpdate) AddAddressTrackers(a ...*AddressTracker) *ChainUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.AddAddressTrackerIDs(ids...)
}

// AddValidatorIDs adds the "validators" edge to the Validator entity by IDs.
func (cu *ChainUpdate) AddValidatorIDs(ids ...int) *ChainUpdate {
	cu.mutation.AddValidatorIDs(ids...)
	return cu
}

// AddValidators adds the "validators" edges to the Validator entity.
func (cu *ChainUpdate) AddValidators(v ...*Validator) *ChainUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.AddValidatorIDs(ids...)
}

// Mutation returns the ChainMutation object of the builder.
func (cu *ChainUpdate) Mutation() *ChainMutation {
	return cu.mutation
}

// ClearChainProposals clears all "chain_proposals" edges to the ChainProposal entity.
func (cu *ChainUpdate) ClearChainProposals() *ChainUpdate {
	cu.mutation.ClearChainProposals()
	return cu
}

// RemoveChainProposalIDs removes the "chain_proposals" edge to ChainProposal entities by IDs.
func (cu *ChainUpdate) RemoveChainProposalIDs(ids ...int) *ChainUpdate {
	cu.mutation.RemoveChainProposalIDs(ids...)
	return cu
}

// RemoveChainProposals removes "chain_proposals" edges to ChainProposal entities.
func (cu *ChainUpdate) RemoveChainProposals(c ...*ChainProposal) *ChainUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveChainProposalIDs(ids...)
}

// ClearTelegramChats clears all "telegram_chats" edges to the TelegramChat entity.
func (cu *ChainUpdate) ClearTelegramChats() *ChainUpdate {
	cu.mutation.ClearTelegramChats()
	return cu
}

// RemoveTelegramChatIDs removes the "telegram_chats" edge to TelegramChat entities by IDs.
func (cu *ChainUpdate) RemoveTelegramChatIDs(ids ...int) *ChainUpdate {
	cu.mutation.RemoveTelegramChatIDs(ids...)
	return cu
}

// RemoveTelegramChats removes "telegram_chats" edges to TelegramChat entities.
func (cu *ChainUpdate) RemoveTelegramChats(t ...*TelegramChat) *ChainUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveTelegramChatIDs(ids...)
}

// ClearDiscordChannels clears all "discord_channels" edges to the DiscordChannel entity.
func (cu *ChainUpdate) ClearDiscordChannels() *ChainUpdate {
	cu.mutation.ClearDiscordChannels()
	return cu
}

// RemoveDiscordChannelIDs removes the "discord_channels" edge to DiscordChannel entities by IDs.
func (cu *ChainUpdate) RemoveDiscordChannelIDs(ids ...int) *ChainUpdate {
	cu.mutation.RemoveDiscordChannelIDs(ids...)
	return cu
}

// RemoveDiscordChannels removes "discord_channels" edges to DiscordChannel entities.
func (cu *ChainUpdate) RemoveDiscordChannels(d ...*DiscordChannel) *ChainUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cu.RemoveDiscordChannelIDs(ids...)
}

// ClearAddressTrackers clears all "address_trackers" edges to the AddressTracker entity.
func (cu *ChainUpdate) ClearAddressTrackers() *ChainUpdate {
	cu.mutation.ClearAddressTrackers()
	return cu
}

// RemoveAddressTrackerIDs removes the "address_trackers" edge to AddressTracker entities by IDs.
func (cu *ChainUpdate) RemoveAddressTrackerIDs(ids ...int) *ChainUpdate {
	cu.mutation.RemoveAddressTrackerIDs(ids...)
	return cu
}

// RemoveAddressTrackers removes "address_trackers" edges to AddressTracker entities.
func (cu *ChainUpdate) RemoveAddressTrackers(a ...*AddressTracker) *ChainUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.RemoveAddressTrackerIDs(ids...)
}

// ClearValidators clears all "validators" edges to the Validator entity.
func (cu *ChainUpdate) ClearValidators() *ChainUpdate {
	cu.mutation.ClearValidators()
	return cu
}

// RemoveValidatorIDs removes the "validators" edge to Validator entities by IDs.
func (cu *ChainUpdate) RemoveValidatorIDs(ids ...int) *ChainUpdate {
	cu.mutation.RemoveValidatorIDs(ids...)
	return cu
}

// RemoveValidators removes "validators" edges to Validator entities.
func (cu *ChainUpdate) RemoveValidators(v ...*Validator) *ChainUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.RemoveValidatorIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChainUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks[int, ChainMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChainUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChainUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChainUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ChainUpdate) defaults() {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := chain.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
}

func (cu *ChainUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(chain.Table, chain.Columns, sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.SetField(chain.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cu.mutation.ChainID(); ok {
		_spec.SetField(chain.FieldChainID, field.TypeString, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(chain.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.PrettyName(); ok {
		_spec.SetField(chain.FieldPrettyName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Path(); ok {
		_spec.SetField(chain.FieldPath, field.TypeString, value)
	}
	if value, ok := cu.mutation.Display(); ok {
		_spec.SetField(chain.FieldDisplay, field.TypeString, value)
	}
	if value, ok := cu.mutation.IsEnabled(); ok {
		_spec.SetField(chain.FieldIsEnabled, field.TypeBool, value)
	}
	if value, ok := cu.mutation.ImageURL(); ok {
		_spec.SetField(chain.FieldImageURL, field.TypeString, value)
	}
	if value, ok := cu.mutation.ThumbnailURL(); ok {
		_spec.SetField(chain.FieldThumbnailURL, field.TypeString, value)
	}
	if value, ok := cu.mutation.Bech32Prefix(); ok {
		_spec.SetField(chain.FieldBech32Prefix, field.TypeString, value)
	}
	if cu.mutation.ChainProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ChainProposalsTable,
			Columns: []string{chain.ChainProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedChainProposalsIDs(); len(nodes) > 0 && !cu.mutation.ChainProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ChainProposalsTable,
			Columns: []string{chain.ChainProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ChainProposalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ChainProposalsTable,
			Columns: []string{chain.ChainProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.TelegramChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.TelegramChatsTable,
			Columns: chain.TelegramChatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedTelegramChatsIDs(); len(nodes) > 0 && !cu.mutation.TelegramChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.TelegramChatsTable,
			Columns: chain.TelegramChatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TelegramChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.TelegramChatsTable,
			Columns: chain.TelegramChatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.DiscordChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.DiscordChannelsTable,
			Columns: chain.DiscordChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedDiscordChannelsIDs(); len(nodes) > 0 && !cu.mutation.DiscordChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.DiscordChannelsTable,
			Columns: chain.DiscordChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.DiscordChannelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.DiscordChannelsTable,
			Columns: chain.DiscordChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.AddressTrackersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.AddressTrackersTable,
			Columns: []string{chain.AddressTrackersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedAddressTrackersIDs(); len(nodes) > 0 && !cu.mutation.AddressTrackersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.AddressTrackersTable,
			Columns: []string{chain.AddressTrackersColumn},
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
	if nodes := cu.mutation.AddressTrackersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.AddressTrackersTable,
			Columns: []string{chain.AddressTrackersColumn},
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
	if cu.mutation.ValidatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ValidatorsTable,
			Columns: []string{chain.ValidatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(validator.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedValidatorsIDs(); len(nodes) > 0 && !cu.mutation.ValidatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ValidatorsTable,
			Columns: []string{chain.ValidatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(validator.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ValidatorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ValidatorsTable,
			Columns: []string{chain.ValidatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(validator.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chain.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChainUpdateOne is the builder for updating a single Chain entity.
type ChainUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChainMutation
}

// SetUpdateTime sets the "update_time" field.
func (cuo *ChainUpdateOne) SetUpdateTime(t time.Time) *ChainUpdateOne {
	cuo.mutation.SetUpdateTime(t)
	return cuo
}

// SetChainID sets the "chain_id" field.
func (cuo *ChainUpdateOne) SetChainID(s string) *ChainUpdateOne {
	cuo.mutation.SetChainID(s)
	return cuo
}

// SetName sets the "name" field.
func (cuo *ChainUpdateOne) SetName(s string) *ChainUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetPrettyName sets the "pretty_name" field.
func (cuo *ChainUpdateOne) SetPrettyName(s string) *ChainUpdateOne {
	cuo.mutation.SetPrettyName(s)
	return cuo
}

// SetPath sets the "path" field.
func (cuo *ChainUpdateOne) SetPath(s string) *ChainUpdateOne {
	cuo.mutation.SetPath(s)
	return cuo
}

// SetDisplay sets the "display" field.
func (cuo *ChainUpdateOne) SetDisplay(s string) *ChainUpdateOne {
	cuo.mutation.SetDisplay(s)
	return cuo
}

// SetNillableDisplay sets the "display" field if the given value is not nil.
func (cuo *ChainUpdateOne) SetNillableDisplay(s *string) *ChainUpdateOne {
	if s != nil {
		cuo.SetDisplay(*s)
	}
	return cuo
}

// SetIsEnabled sets the "is_enabled" field.
func (cuo *ChainUpdateOne) SetIsEnabled(b bool) *ChainUpdateOne {
	cuo.mutation.SetIsEnabled(b)
	return cuo
}

// SetNillableIsEnabled sets the "is_enabled" field if the given value is not nil.
func (cuo *ChainUpdateOne) SetNillableIsEnabled(b *bool) *ChainUpdateOne {
	if b != nil {
		cuo.SetIsEnabled(*b)
	}
	return cuo
}

// SetImageURL sets the "image_url" field.
func (cuo *ChainUpdateOne) SetImageURL(s string) *ChainUpdateOne {
	cuo.mutation.SetImageURL(s)
	return cuo
}

// SetThumbnailURL sets the "thumbnail_url" field.
func (cuo *ChainUpdateOne) SetThumbnailURL(s string) *ChainUpdateOne {
	cuo.mutation.SetThumbnailURL(s)
	return cuo
}

// SetNillableThumbnailURL sets the "thumbnail_url" field if the given value is not nil.
func (cuo *ChainUpdateOne) SetNillableThumbnailURL(s *string) *ChainUpdateOne {
	if s != nil {
		cuo.SetThumbnailURL(*s)
	}
	return cuo
}

// SetBech32Prefix sets the "bech32_prefix" field.
func (cuo *ChainUpdateOne) SetBech32Prefix(s string) *ChainUpdateOne {
	cuo.mutation.SetBech32Prefix(s)
	return cuo
}

// AddChainProposalIDs adds the "chain_proposals" edge to the ChainProposal entity by IDs.
func (cuo *ChainUpdateOne) AddChainProposalIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.AddChainProposalIDs(ids...)
	return cuo
}

// AddChainProposals adds the "chain_proposals" edges to the ChainProposal entity.
func (cuo *ChainUpdateOne) AddChainProposals(c ...*ChainProposal) *ChainUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddChainProposalIDs(ids...)
}

// AddTelegramChatIDs adds the "telegram_chats" edge to the TelegramChat entity by IDs.
func (cuo *ChainUpdateOne) AddTelegramChatIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.AddTelegramChatIDs(ids...)
	return cuo
}

// AddTelegramChats adds the "telegram_chats" edges to the TelegramChat entity.
func (cuo *ChainUpdateOne) AddTelegramChats(t ...*TelegramChat) *ChainUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddTelegramChatIDs(ids...)
}

// AddDiscordChannelIDs adds the "discord_channels" edge to the DiscordChannel entity by IDs.
func (cuo *ChainUpdateOne) AddDiscordChannelIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.AddDiscordChannelIDs(ids...)
	return cuo
}

// AddDiscordChannels adds the "discord_channels" edges to the DiscordChannel entity.
func (cuo *ChainUpdateOne) AddDiscordChannels(d ...*DiscordChannel) *ChainUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cuo.AddDiscordChannelIDs(ids...)
}

// AddAddressTrackerIDs adds the "address_trackers" edge to the AddressTracker entity by IDs.
func (cuo *ChainUpdateOne) AddAddressTrackerIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.AddAddressTrackerIDs(ids...)
	return cuo
}

// AddAddressTrackers adds the "address_trackers" edges to the AddressTracker entity.
func (cuo *ChainUpdateOne) AddAddressTrackers(a ...*AddressTracker) *ChainUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.AddAddressTrackerIDs(ids...)
}

// AddValidatorIDs adds the "validators" edge to the Validator entity by IDs.
func (cuo *ChainUpdateOne) AddValidatorIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.AddValidatorIDs(ids...)
	return cuo
}

// AddValidators adds the "validators" edges to the Validator entity.
func (cuo *ChainUpdateOne) AddValidators(v ...*Validator) *ChainUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.AddValidatorIDs(ids...)
}

// Mutation returns the ChainMutation object of the builder.
func (cuo *ChainUpdateOne) Mutation() *ChainMutation {
	return cuo.mutation
}

// ClearChainProposals clears all "chain_proposals" edges to the ChainProposal entity.
func (cuo *ChainUpdateOne) ClearChainProposals() *ChainUpdateOne {
	cuo.mutation.ClearChainProposals()
	return cuo
}

// RemoveChainProposalIDs removes the "chain_proposals" edge to ChainProposal entities by IDs.
func (cuo *ChainUpdateOne) RemoveChainProposalIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.RemoveChainProposalIDs(ids...)
	return cuo
}

// RemoveChainProposals removes "chain_proposals" edges to ChainProposal entities.
func (cuo *ChainUpdateOne) RemoveChainProposals(c ...*ChainProposal) *ChainUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveChainProposalIDs(ids...)
}

// ClearTelegramChats clears all "telegram_chats" edges to the TelegramChat entity.
func (cuo *ChainUpdateOne) ClearTelegramChats() *ChainUpdateOne {
	cuo.mutation.ClearTelegramChats()
	return cuo
}

// RemoveTelegramChatIDs removes the "telegram_chats" edge to TelegramChat entities by IDs.
func (cuo *ChainUpdateOne) RemoveTelegramChatIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.RemoveTelegramChatIDs(ids...)
	return cuo
}

// RemoveTelegramChats removes "telegram_chats" edges to TelegramChat entities.
func (cuo *ChainUpdateOne) RemoveTelegramChats(t ...*TelegramChat) *ChainUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveTelegramChatIDs(ids...)
}

// ClearDiscordChannels clears all "discord_channels" edges to the DiscordChannel entity.
func (cuo *ChainUpdateOne) ClearDiscordChannels() *ChainUpdateOne {
	cuo.mutation.ClearDiscordChannels()
	return cuo
}

// RemoveDiscordChannelIDs removes the "discord_channels" edge to DiscordChannel entities by IDs.
func (cuo *ChainUpdateOne) RemoveDiscordChannelIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.RemoveDiscordChannelIDs(ids...)
	return cuo
}

// RemoveDiscordChannels removes "discord_channels" edges to DiscordChannel entities.
func (cuo *ChainUpdateOne) RemoveDiscordChannels(d ...*DiscordChannel) *ChainUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cuo.RemoveDiscordChannelIDs(ids...)
}

// ClearAddressTrackers clears all "address_trackers" edges to the AddressTracker entity.
func (cuo *ChainUpdateOne) ClearAddressTrackers() *ChainUpdateOne {
	cuo.mutation.ClearAddressTrackers()
	return cuo
}

// RemoveAddressTrackerIDs removes the "address_trackers" edge to AddressTracker entities by IDs.
func (cuo *ChainUpdateOne) RemoveAddressTrackerIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.RemoveAddressTrackerIDs(ids...)
	return cuo
}

// RemoveAddressTrackers removes "address_trackers" edges to AddressTracker entities.
func (cuo *ChainUpdateOne) RemoveAddressTrackers(a ...*AddressTracker) *ChainUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.RemoveAddressTrackerIDs(ids...)
}

// ClearValidators clears all "validators" edges to the Validator entity.
func (cuo *ChainUpdateOne) ClearValidators() *ChainUpdateOne {
	cuo.mutation.ClearValidators()
	return cuo
}

// RemoveValidatorIDs removes the "validators" edge to Validator entities by IDs.
func (cuo *ChainUpdateOne) RemoveValidatorIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.RemoveValidatorIDs(ids...)
	return cuo
}

// RemoveValidators removes "validators" edges to Validator entities.
func (cuo *ChainUpdateOne) RemoveValidators(v ...*Validator) *ChainUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.RemoveValidatorIDs(ids...)
}

// Where appends a list predicates to the ChainUpdate builder.
func (cuo *ChainUpdateOne) Where(ps ...predicate.Chain) *ChainUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChainUpdateOne) Select(field string, fields ...string) *ChainUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chain entity.
func (cuo *ChainUpdateOne) Save(ctx context.Context) (*Chain, error) {
	cuo.defaults()
	return withHooks[*Chain, ChainMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChainUpdateOne) SaveX(ctx context.Context) *Chain {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChainUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChainUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ChainUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := chain.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
}

func (cuo *ChainUpdateOne) sqlSave(ctx context.Context) (_node *Chain, err error) {
	_spec := sqlgraph.NewUpdateSpec(chain.Table, chain.Columns, sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Chain.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chain.FieldID)
		for _, f := range fields {
			if !chain.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chain.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.SetField(chain.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.ChainID(); ok {
		_spec.SetField(chain.FieldChainID, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(chain.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.PrettyName(); ok {
		_spec.SetField(chain.FieldPrettyName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Path(); ok {
		_spec.SetField(chain.FieldPath, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Display(); ok {
		_spec.SetField(chain.FieldDisplay, field.TypeString, value)
	}
	if value, ok := cuo.mutation.IsEnabled(); ok {
		_spec.SetField(chain.FieldIsEnabled, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.ImageURL(); ok {
		_spec.SetField(chain.FieldImageURL, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ThumbnailURL(); ok {
		_spec.SetField(chain.FieldThumbnailURL, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Bech32Prefix(); ok {
		_spec.SetField(chain.FieldBech32Prefix, field.TypeString, value)
	}
	if cuo.mutation.ChainProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ChainProposalsTable,
			Columns: []string{chain.ChainProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedChainProposalsIDs(); len(nodes) > 0 && !cuo.mutation.ChainProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ChainProposalsTable,
			Columns: []string{chain.ChainProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ChainProposalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ChainProposalsTable,
			Columns: []string{chain.ChainProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.TelegramChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.TelegramChatsTable,
			Columns: chain.TelegramChatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedTelegramChatsIDs(); len(nodes) > 0 && !cuo.mutation.TelegramChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.TelegramChatsTable,
			Columns: chain.TelegramChatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TelegramChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.TelegramChatsTable,
			Columns: chain.TelegramChatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.DiscordChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.DiscordChannelsTable,
			Columns: chain.DiscordChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedDiscordChannelsIDs(); len(nodes) > 0 && !cuo.mutation.DiscordChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.DiscordChannelsTable,
			Columns: chain.DiscordChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.DiscordChannelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.DiscordChannelsTable,
			Columns: chain.DiscordChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.AddressTrackersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.AddressTrackersTable,
			Columns: []string{chain.AddressTrackersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedAddressTrackersIDs(); len(nodes) > 0 && !cuo.mutation.AddressTrackersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.AddressTrackersTable,
			Columns: []string{chain.AddressTrackersColumn},
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
	if nodes := cuo.mutation.AddressTrackersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.AddressTrackersTable,
			Columns: []string{chain.AddressTrackersColumn},
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
	if cuo.mutation.ValidatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ValidatorsTable,
			Columns: []string{chain.ValidatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(validator.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedValidatorsIDs(); len(nodes) > 0 && !cuo.mutation.ValidatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ValidatorsTable,
			Columns: []string{chain.ValidatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(validator.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ValidatorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ValidatorsTable,
			Columns: []string{chain.ValidatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(validator.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Chain{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chain.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
