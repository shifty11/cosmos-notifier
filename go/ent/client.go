// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/shifty11/dao-dao-notifier/ent/migrate"

	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/ent/discordchannel"
	"github.com/shifty11/dao-dao-notifier/ent/proposal"
	"github.com/shifty11/dao-dao-notifier/ent/telegramchat"
	"github.com/shifty11/dao-dao-notifier/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Contract is the client for interacting with the Contract builders.
	Contract *ContractClient
	// DiscordChannel is the client for interacting with the DiscordChannel builders.
	DiscordChannel *DiscordChannelClient
	// Proposal is the client for interacting with the Proposal builders.
	Proposal *ProposalClient
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
	c.Contract = NewContractClient(c.config)
	c.DiscordChannel = NewDiscordChannelClient(c.config)
	c.Proposal = NewProposalClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Contract:       NewContractClient(cfg),
		DiscordChannel: NewDiscordChannelClient(cfg),
		Proposal:       NewProposalClient(cfg),
		TelegramChat:   NewTelegramChatClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:            ctx,
		config:         cfg,
		Contract:       NewContractClient(cfg),
		DiscordChannel: NewDiscordChannelClient(cfg),
		Proposal:       NewProposalClient(cfg),
		TelegramChat:   NewTelegramChatClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Contract.
//		Query().
//		Count(ctx)
//
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
	c.Contract.Use(hooks...)
	c.DiscordChannel.Use(hooks...)
	c.Proposal.Use(hooks...)
	c.TelegramChat.Use(hooks...)
	c.User.Use(hooks...)
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

// DeleteOne returns a builder for deleting the given entity by its id.
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
func (c *ContractClient) QueryProposals(co *Contract) *ProposalQuery {
	query := &ProposalQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(contract.Table, contract.FieldID, id),
			sqlgraph.To(proposal.Table, proposal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, contract.ProposalsTable, contract.ProposalsColumn),
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

// DeleteOne returns a builder for deleting the given entity by its id.
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

// QueryUser queries the user edge of a DiscordChannel.
func (c *DiscordChannelClient) QueryUser(dc *DiscordChannel) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(discordchannel.Table, discordchannel.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, discordchannel.UserTable, discordchannel.UserColumn),
		)
		fromV = sqlgraph.Neighbors(dc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryContracts queries the contracts edge of a DiscordChannel.
func (c *DiscordChannelClient) QueryContracts(dc *DiscordChannel) *ContractQuery {
	query := &ContractQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(discordchannel.Table, discordchannel.FieldID, id),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, discordchannel.ContractsTable, discordchannel.ContractsColumn),
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

// ProposalClient is a client for the Proposal schema.
type ProposalClient struct {
	config
}

// NewProposalClient returns a client for the Proposal from the given config.
func NewProposalClient(c config) *ProposalClient {
	return &ProposalClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `proposal.Hooks(f(g(h())))`.
func (c *ProposalClient) Use(hooks ...Hook) {
	c.hooks.Proposal = append(c.hooks.Proposal, hooks...)
}

// Create returns a builder for creating a Proposal entity.
func (c *ProposalClient) Create() *ProposalCreate {
	mutation := newProposalMutation(c.config, OpCreate)
	return &ProposalCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Proposal entities.
func (c *ProposalClient) CreateBulk(builders ...*ProposalCreate) *ProposalCreateBulk {
	return &ProposalCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Proposal.
func (c *ProposalClient) Update() *ProposalUpdate {
	mutation := newProposalMutation(c.config, OpUpdate)
	return &ProposalUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProposalClient) UpdateOne(pr *Proposal) *ProposalUpdateOne {
	mutation := newProposalMutation(c.config, OpUpdateOne, withProposal(pr))
	return &ProposalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProposalClient) UpdateOneID(id int) *ProposalUpdateOne {
	mutation := newProposalMutation(c.config, OpUpdateOne, withProposalID(id))
	return &ProposalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Proposal.
func (c *ProposalClient) Delete() *ProposalDelete {
	mutation := newProposalMutation(c.config, OpDelete)
	return &ProposalDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProposalClient) DeleteOne(pr *Proposal) *ProposalDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ProposalClient) DeleteOneID(id int) *ProposalDeleteOne {
	builder := c.Delete().Where(proposal.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProposalDeleteOne{builder}
}

// Query returns a query builder for Proposal.
func (c *ProposalClient) Query() *ProposalQuery {
	return &ProposalQuery{
		config: c.config,
	}
}

// Get returns a Proposal entity by its id.
func (c *ProposalClient) Get(ctx context.Context, id int) (*Proposal, error) {
	return c.Query().Where(proposal.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProposalClient) GetX(ctx context.Context, id int) *Proposal {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryContract queries the contract edge of a Proposal.
func (c *ProposalClient) QueryContract(pr *Proposal) *ContractQuery {
	query := &ContractQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(proposal.Table, proposal.FieldID, id),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, proposal.ContractTable, proposal.ContractColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProposalClient) Hooks() []Hook {
	return c.hooks.Proposal
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

// DeleteOne returns a builder for deleting the given entity by its id.
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

// QueryUser queries the user edge of a TelegramChat.
func (c *TelegramChatClient) QueryUser(tc *TelegramChat) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := tc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(telegramchat.Table, telegramchat.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, telegramchat.UserTable, telegramchat.UserColumn),
		)
		fromV = sqlgraph.Neighbors(tc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryContracts queries the contracts edge of a TelegramChat.
func (c *TelegramChatClient) QueryContracts(tc *TelegramChat) *ContractQuery {
	query := &ContractQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := tc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(telegramchat.Table, telegramchat.FieldID, id),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, telegramchat.ContractsTable, telegramchat.ContractsColumn),
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

// DeleteOne returns a builder for deleting the given entity by its id.
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
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(telegramchat.Table, telegramchat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.TelegramChatsTable, user.TelegramChatsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDiscordChannels queries the discord_channels edge of a User.
func (c *UserClient) QueryDiscordChannels(u *User) *DiscordChannelQuery {
	query := &DiscordChannelQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(discordchannel.Table, discordchannel.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.DiscordChannelsTable, user.DiscordChannelsColumn),
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
