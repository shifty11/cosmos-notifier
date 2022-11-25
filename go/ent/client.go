// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/shifty11/cosmos-notifier/ent/migrate"

	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/chainproposal"
	"github.com/shifty11/cosmos-notifier/ent/contract"
	"github.com/shifty11/cosmos-notifier/ent/contractproposal"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
	"github.com/shifty11/cosmos-notifier/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Chain is the client for interacting with the Chain builders.
	Chain *ChainClient
	// ChainProposal is the client for interacting with the ChainProposal builders.
	ChainProposal *ChainProposalClient
	// Contract is the client for interacting with the Contract builders.
	Contract *ContractClient
	// ContractProposal is the client for interacting with the ContractProposal builders.
	ContractProposal *ContractProposalClient
	// DiscordChannel is the client for interacting with the DiscordChannel builders.
	DiscordChannel *DiscordChannelClient
	// TelegramChat is the client for interacting with the TelegramChat builders.
	TelegramChat *TelegramChatClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Chain = NewChainClient(c.config)
	c.ChainProposal = NewChainProposalClient(c.config)
	c.Contract = NewContractClient(c.config)
	c.ContractProposal = NewContractProposalClient(c.config)
	c.DiscordChannel = NewDiscordChannelClient(c.config)
	c.TelegramChat = NewTelegramChatClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		Chain:            NewChainClient(cfg),
		ChainProposal:    NewChainProposalClient(cfg),
		Contract:         NewContractClient(cfg),
		ContractProposal: NewContractProposalClient(cfg),
		DiscordChannel:   NewDiscordChannelClient(cfg),
		TelegramChat:     NewTelegramChatClient(cfg),
		User:             NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		Chain:            NewChainClient(cfg),
		ChainProposal:    NewChainProposalClient(cfg),
		Contract:         NewContractClient(cfg),
		ContractProposal: NewContractProposalClient(cfg),
		DiscordChannel:   NewDiscordChannelClient(cfg),
		TelegramChat:     NewTelegramChatClient(cfg),
		User:             NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Chain.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Chain.Use(hooks...)
	c.ChainProposal.Use(hooks...)
	c.Contract.Use(hooks...)
	c.ContractProposal.Use(hooks...)
	c.DiscordChannel.Use(hooks...)
	c.TelegramChat.Use(hooks...)
	c.User.Use(hooks...)
}

// ChainClient is a client for the Chain schema.
type ChainClient struct {
	config
}

// NewChainClient returns a client for the Chain from the given config.
func NewChainClient(c config) *ChainClient {
	return &ChainClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chain.Hooks(f(g(h())))`.
func (c *ChainClient) Use(hooks ...Hook) {
	c.hooks.Chain = append(c.hooks.Chain, hooks...)
}

// Create returns a builder for creating a Chain entity.
func (c *ChainClient) Create() *ChainCreate {
	mutation := newChainMutation(c.config, OpCreate)
	return &ChainCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Chain entities.
func (c *ChainClient) CreateBulk(builders ...*ChainCreate) *ChainCreateBulk {
	return &ChainCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Chain.
func (c *ChainClient) Update() *ChainUpdate {
	mutation := newChainMutation(c.config, OpUpdate)
	return &ChainUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChainClient) UpdateOne(ch *Chain) *ChainUpdateOne {
	mutation := newChainMutation(c.config, OpUpdateOne, withChain(ch))
	return &ChainUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChainClient) UpdateOneID(id int) *ChainUpdateOne {
	mutation := newChainMutation(c.config, OpUpdateOne, withChainID(id))
	return &ChainUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Chain.
func (c *ChainClient) Delete() *ChainDelete {
	mutation := newChainMutation(c.config, OpDelete)
	return &ChainDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChainClient) DeleteOne(ch *Chain) *ChainDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChainClient) DeleteOneID(id int) *ChainDeleteOne {
	builder := c.Delete().Where(chain.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChainDeleteOne{builder}
}

// Query returns a query builder for Chain.
func (c *ChainClient) Query() *ChainQuery {
	return &ChainQuery{
		config: c.config,
	}
}

// Get returns a Chain entity by its id.
func (c *ChainClient) Get(ctx context.Context, id int) (*Chain, error) {
	return c.Query().Where(chain.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChainClient) GetX(ctx context.Context, id int) *Chain {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChainProposals queries the chain_proposals edge of a Chain.
func (c *ChainClient) QueryChainProposals(ch *Chain) *ChainProposalQuery {
	query := &ChainProposalQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chain.Table, chain.FieldID, id),
			sqlgraph.To(chainproposal.Table, chainproposal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, chain.ChainProposalsTable, chain.ChainProposalsColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTelegramChats queries the telegram_chats edge of a Chain.
func (c *ChainClient) QueryTelegramChats(ch *Chain) *TelegramChatQuery {
	query := &TelegramChatQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chain.Table, chain.FieldID, id),
			sqlgraph.To(telegramchat.Table, telegramchat.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, chain.TelegramChatsTable, chain.TelegramChatsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDiscordChannels queries the discord_channels edge of a Chain.
func (c *ChainClient) QueryDiscordChannels(ch *Chain) *DiscordChannelQuery {
	query := &DiscordChannelQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chain.Table, chain.FieldID, id),
			sqlgraph.To(discordchannel.Table, discordchannel.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, chain.DiscordChannelsTable, chain.DiscordChannelsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChainClient) Hooks() []Hook {
	return c.hooks.Chain
}

// ChainProposalClient is a client for the ChainProposal schema.
type ChainProposalClient struct {
	config
}

// NewChainProposalClient returns a client for the ChainProposal from the given config.
func NewChainProposalClient(c config) *ChainProposalClient {
	return &ChainProposalClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chainproposal.Hooks(f(g(h())))`.
func (c *ChainProposalClient) Use(hooks ...Hook) {
	c.hooks.ChainProposal = append(c.hooks.ChainProposal, hooks...)
}

// Create returns a builder for creating a ChainProposal entity.
func (c *ChainProposalClient) Create() *ChainProposalCreate {
	mutation := newChainProposalMutation(c.config, OpCreate)
	return &ChainProposalCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChainProposal entities.
func (c *ChainProposalClient) CreateBulk(builders ...*ChainProposalCreate) *ChainProposalCreateBulk {
	return &ChainProposalCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChainProposal.
func (c *ChainProposalClient) Update() *ChainProposalUpdate {
	mutation := newChainProposalMutation(c.config, OpUpdate)
	return &ChainProposalUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChainProposalClient) UpdateOne(cp *ChainProposal) *ChainProposalUpdateOne {
	mutation := newChainProposalMutation(c.config, OpUpdateOne, withChainProposal(cp))
	return &ChainProposalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChainProposalClient) UpdateOneID(id int) *ChainProposalUpdateOne {
	mutation := newChainProposalMutation(c.config, OpUpdateOne, withChainProposalID(id))
	return &ChainProposalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChainProposal.
func (c *ChainProposalClient) Delete() *ChainProposalDelete {
	mutation := newChainProposalMutation(c.config, OpDelete)
	return &ChainProposalDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChainProposalClient) DeleteOne(cp *ChainProposal) *ChainProposalDeleteOne {
	return c.DeleteOneID(cp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChainProposalClient) DeleteOneID(id int) *ChainProposalDeleteOne {
	builder := c.Delete().Where(chainproposal.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChainProposalDeleteOne{builder}
}

// Query returns a query builder for ChainProposal.
func (c *ChainProposalClient) Query() *ChainProposalQuery {
	return &ChainProposalQuery{
		config: c.config,
	}
}

// Get returns a ChainProposal entity by its id.
func (c *ChainProposalClient) Get(ctx context.Context, id int) (*ChainProposal, error) {
	return c.Query().Where(chainproposal.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChainProposalClient) GetX(ctx context.Context, id int) *ChainProposal {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChain queries the chain edge of a ChainProposal.
func (c *ChainProposalClient) QueryChain(cp *ChainProposal) *ChainQuery {
	query := &ChainQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chainproposal.Table, chainproposal.FieldID, id),
			sqlgraph.To(chain.Table, chain.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, chainproposal.ChainTable, chainproposal.ChainColumn),
		)
		fromV = sqlgraph.Neighbors(cp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChainProposalClient) Hooks() []Hook {
	return c.hooks.ChainProposal
}

// ContractClient is a client for the Contract schema.
type ContractClient struct {
	config
}

// NewContractClient returns a client for the Contract from the given config.
func NewContractClient(c config) *ContractClient {
	return &ContractClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `contract.Hooks(f(g(h())))`.
func (c *ContractClient) Use(hooks ...Hook) {
	c.hooks.Contract = append(c.hooks.Contract, hooks...)
}

// Create returns a builder for creating a Contract entity.
func (c *ContractClient) Create() *ContractCreate {
	mutation := newContractMutation(c.config, OpCreate)
	return &ContractCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Contract entities.
func (c *ContractClient) CreateBulk(builders ...*ContractCreate) *ContractCreateBulk {
	return &ContractCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Contract.
func (c *ContractClient) Update() *ContractUpdate {
	mutation := newContractMutation(c.config, OpUpdate)
	return &ContractUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ContractClient) UpdateOne(co *Contract) *ContractUpdateOne {
	mutation := newContractMutation(c.config, OpUpdateOne, withContract(co))
	return &ContractUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ContractClient) UpdateOneID(id int) *ContractUpdateOne {
	mutation := newContractMutation(c.config, OpUpdateOne, withContractID(id))
	return &ContractUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Contract.
func (c *ContractClient) Delete() *ContractDelete {
	mutation := newContractMutation(c.config, OpDelete)
	return &ContractDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ContractClient) DeleteOne(co *Contract) *ContractDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ContractClient) DeleteOneID(id int) *ContractDeleteOne {
	builder := c.Delete().Where(contract.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ContractDeleteOne{builder}
}

// Query returns a query builder for Contract.
func (c *ContractClient) Query() *ContractQuery {
	return &ContractQuery{
		config: c.config,
	}
}

// Get returns a Contract entity by its id.
func (c *ContractClient) Get(ctx context.Context, id int) (*Contract, error) {
	return c.Query().Where(contract.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ContractClient) GetX(ctx context.Context, id int) *Contract {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProposals queries the proposals edge of a Contract.
func (c *ContractClient) QueryProposals(co *Contract) *ContractProposalQuery {
	query := &ContractProposalQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(contract.Table, contract.FieldID, id),
			sqlgraph.To(contractproposal.Table, contractproposal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, contract.ProposalsTable, contract.ProposalsColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTelegramChats queries the telegram_chats edge of a Contract.
func (c *ContractClient) QueryTelegramChats(co *Contract) *TelegramChatQuery {
	query := &TelegramChatQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(contract.Table, contract.FieldID, id),
			sqlgraph.To(telegramchat.Table, telegramchat.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, contract.TelegramChatsTable, contract.TelegramChatsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDiscordChannels queries the discord_channels edge of a Contract.
func (c *ContractClient) QueryDiscordChannels(co *Contract) *DiscordChannelQuery {
	query := &DiscordChannelQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(contract.Table, contract.FieldID, id),
			sqlgraph.To(discordchannel.Table, discordchannel.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, contract.DiscordChannelsTable, contract.DiscordChannelsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ContractClient) Hooks() []Hook {
	return c.hooks.Contract
}

// ContractProposalClient is a client for the ContractProposal schema.
type ContractProposalClient struct {
	config
}

// NewContractProposalClient returns a client for the ContractProposal from the given config.
func NewContractProposalClient(c config) *ContractProposalClient {
	return &ContractProposalClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `contractproposal.Hooks(f(g(h())))`.
func (c *ContractProposalClient) Use(hooks ...Hook) {
	c.hooks.ContractProposal = append(c.hooks.ContractProposal, hooks...)
}

// Create returns a builder for creating a ContractProposal entity.
func (c *ContractProposalClient) Create() *ContractProposalCreate {
	mutation := newContractProposalMutation(c.config, OpCreate)
	return &ContractProposalCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ContractProposal entities.
func (c *ContractProposalClient) CreateBulk(builders ...*ContractProposalCreate) *ContractProposalCreateBulk {
	return &ContractProposalCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ContractProposal.
func (c *ContractProposalClient) Update() *ContractProposalUpdate {
	mutation := newContractProposalMutation(c.config, OpUpdate)
	return &ContractProposalUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ContractProposalClient) UpdateOne(cp *ContractProposal) *ContractProposalUpdateOne {
	mutation := newContractProposalMutation(c.config, OpUpdateOne, withContractProposal(cp))
	return &ContractProposalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ContractProposalClient) UpdateOneID(id int) *ContractProposalUpdateOne {
	mutation := newContractProposalMutation(c.config, OpUpdateOne, withContractProposalID(id))
	return &ContractProposalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ContractProposal.
func (c *ContractProposalClient) Delete() *ContractProposalDelete {
	mutation := newContractProposalMutation(c.config, OpDelete)
	return &ContractProposalDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ContractProposalClient) DeleteOne(cp *ContractProposal) *ContractProposalDeleteOne {
	return c.DeleteOneID(cp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ContractProposalClient) DeleteOneID(id int) *ContractProposalDeleteOne {
	builder := c.Delete().Where(contractproposal.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ContractProposalDeleteOne{builder}
}

// Query returns a query builder for ContractProposal.
func (c *ContractProposalClient) Query() *ContractProposalQuery {
	return &ContractProposalQuery{
		config: c.config,
	}
}

// Get returns a ContractProposal entity by its id.
func (c *ContractProposalClient) Get(ctx context.Context, id int) (*ContractProposal, error) {
	return c.Query().Where(contractproposal.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ContractProposalClient) GetX(ctx context.Context, id int) *ContractProposal {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryContract queries the contract edge of a ContractProposal.
func (c *ContractProposalClient) QueryContract(cp *ContractProposal) *ContractQuery {
	query := &ContractQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(contractproposal.Table, contractproposal.FieldID, id),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, contractproposal.ContractTable, contractproposal.ContractColumn),
		)
		fromV = sqlgraph.Neighbors(cp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ContractProposalClient) Hooks() []Hook {
	return c.hooks.ContractProposal
}

// DiscordChannelClient is a client for the DiscordChannel schema.
type DiscordChannelClient struct {
	config
}

// NewDiscordChannelClient returns a client for the DiscordChannel from the given config.
func NewDiscordChannelClient(c config) *DiscordChannelClient {
	return &DiscordChannelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `discordchannel.Hooks(f(g(h())))`.
func (c *DiscordChannelClient) Use(hooks ...Hook) {
	c.hooks.DiscordChannel = append(c.hooks.DiscordChannel, hooks...)
}

// Create returns a builder for creating a DiscordChannel entity.
func (c *DiscordChannelClient) Create() *DiscordChannelCreate {
	mutation := newDiscordChannelMutation(c.config, OpCreate)
	return &DiscordChannelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DiscordChannel entities.
func (c *DiscordChannelClient) CreateBulk(builders ...*DiscordChannelCreate) *DiscordChannelCreateBulk {
	return &DiscordChannelCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DiscordChannel.
func (c *DiscordChannelClient) Update() *DiscordChannelUpdate {
	mutation := newDiscordChannelMutation(c.config, OpUpdate)
	return &DiscordChannelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DiscordChannelClient) UpdateOne(dc *DiscordChannel) *DiscordChannelUpdateOne {
	mutation := newDiscordChannelMutation(c.config, OpUpdateOne, withDiscordChannel(dc))
	return &DiscordChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DiscordChannelClient) UpdateOneID(id int) *DiscordChannelUpdateOne {
	mutation := newDiscordChannelMutation(c.config, OpUpdateOne, withDiscordChannelID(id))
	return &DiscordChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DiscordChannel.
func (c *DiscordChannelClient) Delete() *DiscordChannelDelete {
	mutation := newDiscordChannelMutation(c.config, OpDelete)
	return &DiscordChannelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DiscordChannelClient) DeleteOne(dc *DiscordChannel) *DiscordChannelDeleteOne {
	return c.DeleteOneID(dc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DiscordChannelClient) DeleteOneID(id int) *DiscordChannelDeleteOne {
	builder := c.Delete().Where(discordchannel.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DiscordChannelDeleteOne{builder}
}

// Query returns a query builder for DiscordChannel.
func (c *DiscordChannelClient) Query() *DiscordChannelQuery {
	return &DiscordChannelQuery{
		config: c.config,
	}
}

// Get returns a DiscordChannel entity by its id.
func (c *DiscordChannelClient) Get(ctx context.Context, id int) (*DiscordChannel, error) {
	return c.Query().Where(discordchannel.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DiscordChannelClient) GetX(ctx context.Context, id int) *DiscordChannel {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a DiscordChannel.
func (c *DiscordChannelClient) QueryUsers(dc *DiscordChannel) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := dc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(discordchannel.Table, discordchannel.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, discordchannel.UsersTable, discordchannel.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(dc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryContracts queries the contracts edge of a DiscordChannel.
func (c *DiscordChannelClient) QueryContracts(dc *DiscordChannel) *ContractQuery {
	query := &ContractQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := dc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(discordchannel.Table, discordchannel.FieldID, id),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, discordchannel.ContractsTable, discordchannel.ContractsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(dc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChains queries the chains edge of a DiscordChannel.
func (c *DiscordChannelClient) QueryChains(dc *DiscordChannel) *ChainQuery {
	query := &ChainQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := dc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(discordchannel.Table, discordchannel.FieldID, id),
			sqlgraph.To(chain.Table, chain.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, discordchannel.ChainsTable, discordchannel.ChainsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(dc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DiscordChannelClient) Hooks() []Hook {
	return c.hooks.DiscordChannel
}

// TelegramChatClient is a client for the TelegramChat schema.
type TelegramChatClient struct {
	config
}

// NewTelegramChatClient returns a client for the TelegramChat from the given config.
func NewTelegramChatClient(c config) *TelegramChatClient {
	return &TelegramChatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `telegramchat.Hooks(f(g(h())))`.
func (c *TelegramChatClient) Use(hooks ...Hook) {
	c.hooks.TelegramChat = append(c.hooks.TelegramChat, hooks...)
}

// Create returns a builder for creating a TelegramChat entity.
func (c *TelegramChatClient) Create() *TelegramChatCreate {
	mutation := newTelegramChatMutation(c.config, OpCreate)
	return &TelegramChatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TelegramChat entities.
func (c *TelegramChatClient) CreateBulk(builders ...*TelegramChatCreate) *TelegramChatCreateBulk {
	return &TelegramChatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TelegramChat.
func (c *TelegramChatClient) Update() *TelegramChatUpdate {
	mutation := newTelegramChatMutation(c.config, OpUpdate)
	return &TelegramChatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TelegramChatClient) UpdateOne(tc *TelegramChat) *TelegramChatUpdateOne {
	mutation := newTelegramChatMutation(c.config, OpUpdateOne, withTelegramChat(tc))
	return &TelegramChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TelegramChatClient) UpdateOneID(id int) *TelegramChatUpdateOne {
	mutation := newTelegramChatMutation(c.config, OpUpdateOne, withTelegramChatID(id))
	return &TelegramChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TelegramChat.
func (c *TelegramChatClient) Delete() *TelegramChatDelete {
	mutation := newTelegramChatMutation(c.config, OpDelete)
	return &TelegramChatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TelegramChatClient) DeleteOne(tc *TelegramChat) *TelegramChatDeleteOne {
	return c.DeleteOneID(tc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TelegramChatClient) DeleteOneID(id int) *TelegramChatDeleteOne {
	builder := c.Delete().Where(telegramchat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TelegramChatDeleteOne{builder}
}

// Query returns a query builder for TelegramChat.
func (c *TelegramChatClient) Query() *TelegramChatQuery {
	return &TelegramChatQuery{
		config: c.config,
	}
}

// Get returns a TelegramChat entity by its id.
func (c *TelegramChatClient) Get(ctx context.Context, id int) (*TelegramChat, error) {
	return c.Query().Where(telegramchat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TelegramChatClient) GetX(ctx context.Context, id int) *TelegramChat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a TelegramChat.
func (c *TelegramChatClient) QueryUsers(tc *TelegramChat) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := tc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(telegramchat.Table, telegramchat.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, telegramchat.UsersTable, telegramchat.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(tc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryContracts queries the contracts edge of a TelegramChat.
func (c *TelegramChatClient) QueryContracts(tc *TelegramChat) *ContractQuery {
	query := &ContractQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := tc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(telegramchat.Table, telegramchat.FieldID, id),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, telegramchat.ContractsTable, telegramchat.ContractsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(tc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChains queries the chains edge of a TelegramChat.
func (c *TelegramChatClient) QueryChains(tc *TelegramChat) *ChainQuery {
	query := &ChainQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := tc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(telegramchat.Table, telegramchat.FieldID, id),
			sqlgraph.To(chain.Table, chain.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, telegramchat.ChainsTable, telegramchat.ChainsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(tc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TelegramChatClient) Hooks() []Hook {
	return c.hooks.TelegramChat
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTelegramChats queries the telegram_chats edge of a User.
func (c *UserClient) QueryTelegramChats(u *User) *TelegramChatQuery {
	query := &TelegramChatQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(telegramchat.Table, telegramchat.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.TelegramChatsTable, user.TelegramChatsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDiscordChannels queries the discord_channels edge of a User.
func (c *UserClient) QueryDiscordChannels(u *User) *DiscordChannelQuery {
	query := &DiscordChannelQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(discordchannel.Table, discordchannel.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.DiscordChannelsTable, user.DiscordChannelsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
