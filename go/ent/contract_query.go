// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-notifier/ent/contract"
	"github.com/shifty11/cosmos-notifier/ent/contractproposal"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/predicate"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
)

// ContractQuery is the builder for querying Contract entities.
type ContractQuery struct {
	config
	limit               *int
	offset              *int
	unique              *bool
	order               []OrderFunc
	fields              []string
	predicates          []predicate.Contract
	withProposals       *ContractProposalQuery
	withTelegramChats   *TelegramChatQuery
	withDiscordChannels *DiscordChannelQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ContractQuery builder.
func (cq *ContractQuery) Where(ps ...predicate.Contract) *ContractQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *ContractQuery) Limit(limit int) *ContractQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *ContractQuery) Offset(offset int) *ContractQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ContractQuery) Unique(unique bool) *ContractQuery {
	cq.unique = &unique
	return cq
}

// Order adds an order step to the query.
func (cq *ContractQuery) Order(o ...OrderFunc) *ContractQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryProposals chains the current query on the "proposals" edge.
func (cq *ContractQuery) QueryProposals() *ContractProposalQuery {
	query := &ContractProposalQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contract.Table, contract.FieldID, selector),
			sqlgraph.To(contractproposal.Table, contractproposal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, contract.ProposalsTable, contract.ProposalsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTelegramChats chains the current query on the "telegram_chats" edge.
func (cq *ContractQuery) QueryTelegramChats() *TelegramChatQuery {
	query := &TelegramChatQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contract.Table, contract.FieldID, selector),
			sqlgraph.To(telegramchat.Table, telegramchat.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, contract.TelegramChatsTable, contract.TelegramChatsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDiscordChannels chains the current query on the "discord_channels" edge.
func (cq *ContractQuery) QueryDiscordChannels() *DiscordChannelQuery {
	query := &DiscordChannelQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contract.Table, contract.FieldID, selector),
			sqlgraph.To(discordchannel.Table, discordchannel.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, contract.DiscordChannelsTable, contract.DiscordChannelsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Contract entity from the query.
// Returns a *NotFoundError when no Contract was found.
func (cq *ContractQuery) First(ctx context.Context) (*Contract, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{contract.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ContractQuery) FirstX(ctx context.Context) *Contract {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Contract ID from the query.
// Returns a *NotFoundError when no Contract ID was found.
func (cq *ContractQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{contract.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ContractQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Contract entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Contract entity is found.
// Returns a *NotFoundError when no Contract entities are found.
func (cq *ContractQuery) Only(ctx context.Context) (*Contract, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{contract.Label}
	default:
		return nil, &NotSingularError{contract.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ContractQuery) OnlyX(ctx context.Context) *Contract {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Contract ID in the query.
// Returns a *NotSingularError when more than one Contract ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ContractQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{contract.Label}
	default:
		err = &NotSingularError{contract.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ContractQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Contracts.
func (cq *ContractQuery) All(ctx context.Context) ([]*Contract, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *ContractQuery) AllX(ctx context.Context) []*Contract {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Contract IDs.
func (cq *ContractQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cq.Select(contract.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ContractQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ContractQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ContractQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ContractQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ContractQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ContractQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ContractQuery) Clone() *ContractQuery {
	if cq == nil {
		return nil
	}
	return &ContractQuery{
		config:              cq.config,
		limit:               cq.limit,
		offset:              cq.offset,
		order:               append([]OrderFunc{}, cq.order...),
		predicates:          append([]predicate.Contract{}, cq.predicates...),
		withProposals:       cq.withProposals.Clone(),
		withTelegramChats:   cq.withTelegramChats.Clone(),
		withDiscordChannels: cq.withDiscordChannels.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithProposals tells the query-builder to eager-load the nodes that are connected to
// the "proposals" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ContractQuery) WithProposals(opts ...func(*ContractProposalQuery)) *ContractQuery {
	query := &ContractProposalQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withProposals = query
	return cq
}

// WithTelegramChats tells the query-builder to eager-load the nodes that are connected to
// the "telegram_chats" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ContractQuery) WithTelegramChats(opts ...func(*TelegramChatQuery)) *ContractQuery {
	query := &TelegramChatQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withTelegramChats = query
	return cq
}

// WithDiscordChannels tells the query-builder to eager-load the nodes that are connected to
// the "discord_channels" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ContractQuery) WithDiscordChannels(opts ...func(*DiscordChannelQuery)) *ContractQuery {
	query := &DiscordChannelQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withDiscordChannels = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Contract.Query().
//		GroupBy(contract.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ContractQuery) GroupBy(field string, fields ...string) *ContractGroupBy {
	grbuild := &ContractGroupBy{config: cq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(ctx), nil
	}
	grbuild.label = contract.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Contract.Query().
//		Select(contract.FieldCreateTime).
//		Scan(ctx, &v)
func (cq *ContractQuery) Select(fields ...string) *ContractSelect {
	cq.fields = append(cq.fields, fields...)
	selbuild := &ContractSelect{ContractQuery: cq}
	selbuild.label = contract.Label
	selbuild.flds, selbuild.scan = &cq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a ContractSelect configured with the given aggregations.
func (cq *ContractQuery) Aggregate(fns ...AggregateFunc) *ContractSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ContractQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !contract.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ContractQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Contract, error) {
	var (
		nodes       = []*Contract{}
		_spec       = cq.querySpec()
		loadedTypes = [3]bool{
			cq.withProposals != nil,
			cq.withTelegramChats != nil,
			cq.withDiscordChannels != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Contract).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Contract{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withProposals; query != nil {
		if err := cq.loadProposals(ctx, query, nodes,
			func(n *Contract) { n.Edges.Proposals = []*ContractProposal{} },
			func(n *Contract, e *ContractProposal) { n.Edges.Proposals = append(n.Edges.Proposals, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withTelegramChats; query != nil {
		if err := cq.loadTelegramChats(ctx, query, nodes,
			func(n *Contract) { n.Edges.TelegramChats = []*TelegramChat{} },
			func(n *Contract, e *TelegramChat) { n.Edges.TelegramChats = append(n.Edges.TelegramChats, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withDiscordChannels; query != nil {
		if err := cq.loadDiscordChannels(ctx, query, nodes,
			func(n *Contract) { n.Edges.DiscordChannels = []*DiscordChannel{} },
			func(n *Contract, e *DiscordChannel) { n.Edges.DiscordChannels = append(n.Edges.DiscordChannels, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ContractQuery) loadProposals(ctx context.Context, query *ContractProposalQuery, nodes []*Contract, init func(*Contract), assign func(*Contract, *ContractProposal)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Contract)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ContractProposal(func(s *sql.Selector) {
		s.Where(sql.InValues(contract.ProposalsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.contract_proposals
		if fk == nil {
			return fmt.Errorf(`foreign-key "contract_proposals" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "contract_proposals" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ContractQuery) loadTelegramChats(ctx context.Context, query *TelegramChatQuery, nodes []*Contract, init func(*Contract), assign func(*Contract, *TelegramChat)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Contract)
	nids := make(map[int]map[*Contract]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(contract.TelegramChatsTable)
		s.Join(joinT).On(s.C(telegramchat.FieldID), joinT.C(contract.TelegramChatsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(contract.TelegramChatsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(contract.TelegramChatsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Contract]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "telegram_chats" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *ContractQuery) loadDiscordChannels(ctx context.Context, query *DiscordChannelQuery, nodes []*Contract, init func(*Contract), assign func(*Contract, *DiscordChannel)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Contract)
	nids := make(map[int]map[*Contract]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(contract.DiscordChannelsTable)
		s.Join(joinT).On(s.C(discordchannel.FieldID), joinT.C(contract.DiscordChannelsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(contract.DiscordChannelsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(contract.DiscordChannelsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Contract]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "discord_channels" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *ContractQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ContractQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (cq *ContractQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contract.Table,
			Columns: contract.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contract.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contract.FieldID)
		for i := range fields {
			if fields[i] != contract.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ContractQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(contract.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = contract.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ContractGroupBy is the group-by builder for Contract entities.
type ContractGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ContractGroupBy) Aggregate(fns ...AggregateFunc) *ContractGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *ContractGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

func (cgb *ContractGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cgb.fields {
		if !contract.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *ContractGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql.Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
		for _, f := range cgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cgb.fields...)...)
}

// ContractSelect is the builder for selecting fields of Contract entities.
type ContractSelect struct {
	*ContractQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ContractSelect) Aggregate(fns ...AggregateFunc) *ContractSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ContractSelect) Scan(ctx context.Context, v any) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.ContractQuery.sqlQuery(ctx)
	return cs.sqlScan(ctx, v)
}

func (cs *ContractSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(cs.sql))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		cs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		cs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := cs.sql.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
