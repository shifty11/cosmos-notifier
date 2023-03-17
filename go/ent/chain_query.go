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
	"github.com/shifty11/cosmos-notifier/ent/addresstracker"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/chainproposal"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/predicate"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
)

// ChainQuery is the builder for querying Chain entities.
type ChainQuery struct {
	config
	ctx                 *QueryContext
	order               []OrderFunc
	inters              []Interceptor
	predicates          []predicate.Chain
	withChainProposals  *ChainProposalQuery
	withTelegramChats   *TelegramChatQuery
	withDiscordChannels *DiscordChannelQuery
	withAddressTrackers *AddressTrackerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChainQuery builder.
func (cq *ChainQuery) Where(ps ...predicate.Chain) *ChainQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ChainQuery) Limit(limit int) *ChainQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ChainQuery) Offset(offset int) *ChainQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ChainQuery) Unique(unique bool) *ChainQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ChainQuery) Order(o ...OrderFunc) *ChainQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryChainProposals chains the current query on the "chain_proposals" edge.
func (cq *ChainQuery) QueryChainProposals() *ChainProposalQuery {
	query := (&ChainProposalClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chain.Table, chain.FieldID, selector),
			sqlgraph.To(chainproposal.Table, chainproposal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, chain.ChainProposalsTable, chain.ChainProposalsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTelegramChats chains the current query on the "telegram_chats" edge.
func (cq *ChainQuery) QueryTelegramChats() *TelegramChatQuery {
	query := (&TelegramChatClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chain.Table, chain.FieldID, selector),
			sqlgraph.To(telegramchat.Table, telegramchat.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, chain.TelegramChatsTable, chain.TelegramChatsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDiscordChannels chains the current query on the "discord_channels" edge.
func (cq *ChainQuery) QueryDiscordChannels() *DiscordChannelQuery {
	query := (&DiscordChannelClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chain.Table, chain.FieldID, selector),
			sqlgraph.To(discordchannel.Table, discordchannel.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, chain.DiscordChannelsTable, chain.DiscordChannelsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAddressTrackers chains the current query on the "address_trackers" edge.
func (cq *ChainQuery) QueryAddressTrackers() *AddressTrackerQuery {
	query := (&AddressTrackerClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chain.Table, chain.FieldID, selector),
			sqlgraph.To(addresstracker.Table, addresstracker.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, chain.AddressTrackersTable, chain.AddressTrackersColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Chain entity from the query.
// Returns a *NotFoundError when no Chain was found.
func (cq *ChainQuery) First(ctx context.Context) (*Chain, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chain.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ChainQuery) FirstX(ctx context.Context) *Chain {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Chain ID from the query.
// Returns a *NotFoundError when no Chain ID was found.
func (cq *ChainQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chain.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ChainQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Chain entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Chain entity is found.
// Returns a *NotFoundError when no Chain entities are found.
func (cq *ChainQuery) Only(ctx context.Context) (*Chain, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chain.Label}
	default:
		return nil, &NotSingularError{chain.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ChainQuery) OnlyX(ctx context.Context) *Chain {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Chain ID in the query.
// Returns a *NotSingularError when more than one Chain ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ChainQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chain.Label}
	default:
		err = &NotSingularError{chain.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ChainQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Chains.
func (cq *ChainQuery) All(ctx context.Context) ([]*Chain, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Chain, *ChainQuery]()
	return withInterceptors[[]*Chain](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ChainQuery) AllX(ctx context.Context) []*Chain {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Chain IDs.
func (cq *ChainQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(chain.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ChainQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ChainQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ChainQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ChainQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ChainQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ChainQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChainQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ChainQuery) Clone() *ChainQuery {
	if cq == nil {
		return nil
	}
	return &ChainQuery{
		config:              cq.config,
		ctx:                 cq.ctx.Clone(),
		order:               append([]OrderFunc{}, cq.order...),
		inters:              append([]Interceptor{}, cq.inters...),
		predicates:          append([]predicate.Chain{}, cq.predicates...),
		withChainProposals:  cq.withChainProposals.Clone(),
		withTelegramChats:   cq.withTelegramChats.Clone(),
		withDiscordChannels: cq.withDiscordChannels.Clone(),
		withAddressTrackers: cq.withAddressTrackers.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithChainProposals tells the query-builder to eager-load the nodes that are connected to
// the "chain_proposals" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChainQuery) WithChainProposals(opts ...func(*ChainProposalQuery)) *ChainQuery {
	query := (&ChainProposalClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withChainProposals = query
	return cq
}

// WithTelegramChats tells the query-builder to eager-load the nodes that are connected to
// the "telegram_chats" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChainQuery) WithTelegramChats(opts ...func(*TelegramChatQuery)) *ChainQuery {
	query := (&TelegramChatClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withTelegramChats = query
	return cq
}

// WithDiscordChannels tells the query-builder to eager-load the nodes that are connected to
// the "discord_channels" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChainQuery) WithDiscordChannels(opts ...func(*DiscordChannelQuery)) *ChainQuery {
	query := (&DiscordChannelClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withDiscordChannels = query
	return cq
}

// WithAddressTrackers tells the query-builder to eager-load the nodes that are connected to
// the "address_trackers" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChainQuery) WithAddressTrackers(opts ...func(*AddressTrackerQuery)) *ChainQuery {
	query := (&AddressTrackerClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withAddressTrackers = query
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
//	client.Chain.Query().
//		GroupBy(chain.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ChainQuery) GroupBy(field string, fields ...string) *ChainGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChainGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = chain.Label
	grbuild.scan = grbuild.Scan
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
//	client.Chain.Query().
//		Select(chain.FieldCreateTime).
//		Scan(ctx, &v)
func (cq *ChainQuery) Select(fields ...string) *ChainSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ChainSelect{ChainQuery: cq}
	sbuild.label = chain.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChainSelect configured with the given aggregations.
func (cq *ChainQuery) Aggregate(fns ...AggregateFunc) *ChainSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ChainQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !chain.ValidColumn(f) {
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

func (cq *ChainQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Chain, error) {
	var (
		nodes       = []*Chain{}
		_spec       = cq.querySpec()
		loadedTypes = [4]bool{
			cq.withChainProposals != nil,
			cq.withTelegramChats != nil,
			cq.withDiscordChannels != nil,
			cq.withAddressTrackers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Chain).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Chain{config: cq.config}
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
	if query := cq.withChainProposals; query != nil {
		if err := cq.loadChainProposals(ctx, query, nodes,
			func(n *Chain) { n.Edges.ChainProposals = []*ChainProposal{} },
			func(n *Chain, e *ChainProposal) { n.Edges.ChainProposals = append(n.Edges.ChainProposals, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withTelegramChats; query != nil {
		if err := cq.loadTelegramChats(ctx, query, nodes,
			func(n *Chain) { n.Edges.TelegramChats = []*TelegramChat{} },
			func(n *Chain, e *TelegramChat) { n.Edges.TelegramChats = append(n.Edges.TelegramChats, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withDiscordChannels; query != nil {
		if err := cq.loadDiscordChannels(ctx, query, nodes,
			func(n *Chain) { n.Edges.DiscordChannels = []*DiscordChannel{} },
			func(n *Chain, e *DiscordChannel) { n.Edges.DiscordChannels = append(n.Edges.DiscordChannels, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withAddressTrackers; query != nil {
		if err := cq.loadAddressTrackers(ctx, query, nodes,
			func(n *Chain) { n.Edges.AddressTrackers = []*AddressTracker{} },
			func(n *Chain, e *AddressTracker) { n.Edges.AddressTrackers = append(n.Edges.AddressTrackers, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ChainQuery) loadChainProposals(ctx context.Context, query *ChainProposalQuery, nodes []*Chain, init func(*Chain), assign func(*Chain, *ChainProposal)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Chain)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ChainProposal(func(s *sql.Selector) {
		s.Where(sql.InValues(chain.ChainProposalsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.chain_chain_proposals
		if fk == nil {
			return fmt.Errorf(`foreign-key "chain_chain_proposals" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chain_chain_proposals" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ChainQuery) loadTelegramChats(ctx context.Context, query *TelegramChatQuery, nodes []*Chain, init func(*Chain), assign func(*Chain, *TelegramChat)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Chain)
	nids := make(map[int]map[*Chain]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(chain.TelegramChatsTable)
		s.Join(joinT).On(s.C(telegramchat.FieldID), joinT.C(chain.TelegramChatsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(chain.TelegramChatsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(chain.TelegramChatsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
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
					nids[inValue] = map[*Chain]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*TelegramChat](ctx, query, qr, query.inters)
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
func (cq *ChainQuery) loadDiscordChannels(ctx context.Context, query *DiscordChannelQuery, nodes []*Chain, init func(*Chain), assign func(*Chain, *DiscordChannel)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Chain)
	nids := make(map[int]map[*Chain]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(chain.DiscordChannelsTable)
		s.Join(joinT).On(s.C(discordchannel.FieldID), joinT.C(chain.DiscordChannelsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(chain.DiscordChannelsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(chain.DiscordChannelsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
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
					nids[inValue] = map[*Chain]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*DiscordChannel](ctx, query, qr, query.inters)
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
func (cq *ChainQuery) loadAddressTrackers(ctx context.Context, query *AddressTrackerQuery, nodes []*Chain, init func(*Chain), assign func(*Chain, *AddressTracker)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Chain)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AddressTracker(func(s *sql.Selector) {
		s.Where(sql.InValues(chain.AddressTrackersColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.chain_address_trackers
		if fk == nil {
			return fmt.Errorf(`foreign-key "chain_address_trackers" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chain_address_trackers" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *ChainQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ChainQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(chain.Table, chain.Columns, sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chain.FieldID)
		for i := range fields {
			if fields[i] != chain.FieldID {
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
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
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

func (cq *ChainQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(chain.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = chain.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChainGroupBy is the group-by builder for Chain entities.
type ChainGroupBy struct {
	selector
	build *ChainQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ChainGroupBy) Aggregate(fns ...AggregateFunc) *ChainGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ChainGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChainQuery, *ChainGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ChainGroupBy) sqlScan(ctx context.Context, root *ChainQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChainSelect is the builder for selecting fields of Chain entities.
type ChainSelect struct {
	*ChainQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ChainSelect) Aggregate(fns ...AggregateFunc) *ChainSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ChainSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChainQuery, *ChainSelect](ctx, cs.ChainQuery, cs, cs.inters, v)
}

func (cs *ChainSelect) sqlScan(ctx context.Context, root *ChainQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
