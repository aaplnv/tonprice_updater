// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"main/ent/migrate"

	"main/ent/rubchart"
	"main/ent/usdchart"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// RUBChart is the client for interacting with the RUBChart builders.
	RUBChart *RUBChartClient
	// USDChart is the client for interacting with the USDChart builders.
	USDChart *USDChartClient
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
	c.RUBChart = NewRUBChartClient(c.config)
	c.USDChart = NewUSDChartClient(c.config)
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
		ctx:      ctx,
		config:   cfg,
		RUBChart: NewRUBChartClient(cfg),
		USDChart: NewUSDChartClient(cfg),
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
		ctx:      ctx,
		config:   cfg,
		RUBChart: NewRUBChartClient(cfg),
		USDChart: NewUSDChartClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		RUBChart.
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
	c.RUBChart.Use(hooks...)
	c.USDChart.Use(hooks...)
}

// RUBChartClient is a client for the RUBChart schema.
type RUBChartClient struct {
	config
}

// NewRUBChartClient returns a client for the RUBChart from the given config.
func NewRUBChartClient(c config) *RUBChartClient {
	return &RUBChartClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `rubchart.Hooks(f(g(h())))`.
func (c *RUBChartClient) Use(hooks ...Hook) {
	c.hooks.RUBChart = append(c.hooks.RUBChart, hooks...)
}

// Create returns a create builder for RUBChart.
func (c *RUBChartClient) Create() *RUBChartCreate {
	mutation := newRUBChartMutation(c.config, OpCreate)
	return &RUBChartCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RUBChart entities.
func (c *RUBChartClient) CreateBulk(builders ...*RUBChartCreate) *RUBChartCreateBulk {
	return &RUBChartCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RUBChart.
func (c *RUBChartClient) Update() *RUBChartUpdate {
	mutation := newRUBChartMutation(c.config, OpUpdate)
	return &RUBChartUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RUBChartClient) UpdateOne(rc *RUBChart) *RUBChartUpdateOne {
	mutation := newRUBChartMutation(c.config, OpUpdateOne, withRUBChart(rc))
	return &RUBChartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RUBChartClient) UpdateOneID(id int) *RUBChartUpdateOne {
	mutation := newRUBChartMutation(c.config, OpUpdateOne, withRUBChartID(id))
	return &RUBChartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RUBChart.
func (c *RUBChartClient) Delete() *RUBChartDelete {
	mutation := newRUBChartMutation(c.config, OpDelete)
	return &RUBChartDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RUBChartClient) DeleteOne(rc *RUBChart) *RUBChartDeleteOne {
	return c.DeleteOneID(rc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RUBChartClient) DeleteOneID(id int) *RUBChartDeleteOne {
	builder := c.Delete().Where(rubchart.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RUBChartDeleteOne{builder}
}

// Query returns a query builder for RUBChart.
func (c *RUBChartClient) Query() *RUBChartQuery {
	return &RUBChartQuery{
		config: c.config,
	}
}

// Get returns a RUBChart entity by its id.
func (c *RUBChartClient) Get(ctx context.Context, id int) (*RUBChart, error) {
	return c.Query().Where(rubchart.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RUBChartClient) GetX(ctx context.Context, id int) *RUBChart {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RUBChartClient) Hooks() []Hook {
	return c.hooks.RUBChart
}

// USDChartClient is a client for the USDChart schema.
type USDChartClient struct {
	config
}

// NewUSDChartClient returns a client for the USDChart from the given config.
func NewUSDChartClient(c config) *USDChartClient {
	return &USDChartClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usdchart.Hooks(f(g(h())))`.
func (c *USDChartClient) Use(hooks ...Hook) {
	c.hooks.USDChart = append(c.hooks.USDChart, hooks...)
}

// Create returns a create builder for USDChart.
func (c *USDChartClient) Create() *USDChartCreate {
	mutation := newUSDChartMutation(c.config, OpCreate)
	return &USDChartCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of USDChart entities.
func (c *USDChartClient) CreateBulk(builders ...*USDChartCreate) *USDChartCreateBulk {
	return &USDChartCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for USDChart.
func (c *USDChartClient) Update() *USDChartUpdate {
	mutation := newUSDChartMutation(c.config, OpUpdate)
	return &USDChartUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *USDChartClient) UpdateOne(uc *USDChart) *USDChartUpdateOne {
	mutation := newUSDChartMutation(c.config, OpUpdateOne, withUSDChart(uc))
	return &USDChartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *USDChartClient) UpdateOneID(id int) *USDChartUpdateOne {
	mutation := newUSDChartMutation(c.config, OpUpdateOne, withUSDChartID(id))
	return &USDChartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for USDChart.
func (c *USDChartClient) Delete() *USDChartDelete {
	mutation := newUSDChartMutation(c.config, OpDelete)
	return &USDChartDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *USDChartClient) DeleteOne(uc *USDChart) *USDChartDeleteOne {
	return c.DeleteOneID(uc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *USDChartClient) DeleteOneID(id int) *USDChartDeleteOne {
	builder := c.Delete().Where(usdchart.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &USDChartDeleteOne{builder}
}

// Query returns a query builder for USDChart.
func (c *USDChartClient) Query() *USDChartQuery {
	return &USDChartQuery{
		config: c.config,
	}
}

// Get returns a USDChart entity by its id.
func (c *USDChartClient) Get(ctx context.Context, id int) (*USDChart, error) {
	return c.Query().Where(usdchart.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *USDChartClient) GetX(ctx context.Context, id int) *USDChart {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *USDChartClient) Hooks() []Hook {
	return c.hooks.USDChart
}
