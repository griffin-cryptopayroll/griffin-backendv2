// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"griffin-dao/ent/migrate"

	"griffin-dao/ent/crypto_currency"
	"griffin-dao/ent/crypto_prc_source"
	"griffin-dao/ent/employ_type"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer_user_info"
	"griffin-dao/ent/payment_history"
	"griffin-dao/ent/tr_log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CRYPTO_CURRENCY is the client for interacting with the CRYPTO_CURRENCY builders.
	CRYPTO_CURRENCY *CRYPTO_CURRENCYClient
	// CRYPTO_PRC_SOURCE is the client for interacting with the CRYPTO_PRC_SOURCE builders.
	CRYPTO_PRC_SOURCE *CRYPTO_PRC_SOURCEClient
	// EMPLOYEE is the client for interacting with the EMPLOYEE builders.
	EMPLOYEE *EMPLOYEEClient
	// EMPLOYER_USER_INFO is the client for interacting with the EMPLOYER_USER_INFO builders.
	EMPLOYER_USER_INFO *EMPLOYER_USER_INFOClient
	// EMPLOY_TYPE is the client for interacting with the EMPLOY_TYPE builders.
	EMPLOY_TYPE *EMPLOY_TYPEClient
	// PAYMENT_HISTORY is the client for interacting with the PAYMENT_HISTORY builders.
	PAYMENT_HISTORY *PAYMENT_HISTORYClient
	// Tr_log is the client for interacting with the Tr_log builders.
	Tr_log *Tr_logClient
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
	c.CRYPTO_CURRENCY = NewCRYPTO_CURRENCYClient(c.config)
	c.CRYPTO_PRC_SOURCE = NewCRYPTO_PRC_SOURCEClient(c.config)
	c.EMPLOYEE = NewEMPLOYEEClient(c.config)
	c.EMPLOYER_USER_INFO = NewEMPLOYER_USER_INFOClient(c.config)
	c.EMPLOY_TYPE = NewEMPLOY_TYPEClient(c.config)
	c.PAYMENT_HISTORY = NewPAYMENT_HISTORYClient(c.config)
	c.Tr_log = NewTr_logClient(c.config)
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
		ctx:                ctx,
		config:             cfg,
		CRYPTO_CURRENCY:    NewCRYPTO_CURRENCYClient(cfg),
		CRYPTO_PRC_SOURCE:  NewCRYPTO_PRC_SOURCEClient(cfg),
		EMPLOYEE:           NewEMPLOYEEClient(cfg),
		EMPLOYER_USER_INFO: NewEMPLOYER_USER_INFOClient(cfg),
		EMPLOY_TYPE:        NewEMPLOY_TYPEClient(cfg),
		PAYMENT_HISTORY:    NewPAYMENT_HISTORYClient(cfg),
		Tr_log:             NewTr_logClient(cfg),
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
		ctx:                ctx,
		config:             cfg,
		CRYPTO_CURRENCY:    NewCRYPTO_CURRENCYClient(cfg),
		CRYPTO_PRC_SOURCE:  NewCRYPTO_PRC_SOURCEClient(cfg),
		EMPLOYEE:           NewEMPLOYEEClient(cfg),
		EMPLOYER_USER_INFO: NewEMPLOYER_USER_INFOClient(cfg),
		EMPLOY_TYPE:        NewEMPLOY_TYPEClient(cfg),
		PAYMENT_HISTORY:    NewPAYMENT_HISTORYClient(cfg),
		Tr_log:             NewTr_logClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CRYPTO_CURRENCY.
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
	c.CRYPTO_CURRENCY.Use(hooks...)
	c.CRYPTO_PRC_SOURCE.Use(hooks...)
	c.EMPLOYEE.Use(hooks...)
	c.EMPLOYER_USER_INFO.Use(hooks...)
	c.EMPLOY_TYPE.Use(hooks...)
	c.PAYMENT_HISTORY.Use(hooks...)
	c.Tr_log.Use(hooks...)
}

// CRYPTO_CURRENCYClient is a client for the CRYPTO_CURRENCY schema.
type CRYPTO_CURRENCYClient struct {
	config
}

// NewCRYPTO_CURRENCYClient returns a client for the CRYPTO_CURRENCY from the given config.
func NewCRYPTO_CURRENCYClient(c config) *CRYPTO_CURRENCYClient {
	return &CRYPTO_CURRENCYClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `crypto_currency.Hooks(f(g(h())))`.
func (c *CRYPTO_CURRENCYClient) Use(hooks ...Hook) {
	c.hooks.CRYPTO_CURRENCY = append(c.hooks.CRYPTO_CURRENCY, hooks...)
}

// Create returns a builder for creating a CRYPTO_CURRENCY entity.
func (c *CRYPTO_CURRENCYClient) Create() *CRYPTOCURRENCYCreate {
	mutation := newCRYPTOCURRENCYMutation(c.config, OpCreate)
	return &CRYPTOCURRENCYCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CRYPTO_CURRENCY entities.
func (c *CRYPTO_CURRENCYClient) CreateBulk(builders ...*CRYPTOCURRENCYCreate) *CRYPTOCURRENCYCreateBulk {
	return &CRYPTOCURRENCYCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CRYPTO_CURRENCY.
func (c *CRYPTO_CURRENCYClient) Update() *CRYPTOCURRENCYUpdate {
	mutation := newCRYPTOCURRENCYMutation(c.config, OpUpdate)
	return &CRYPTOCURRENCYUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CRYPTO_CURRENCYClient) UpdateOne(cc *CRYPTO_CURRENCY) *CRYPTOCURRENCYUpdateOne {
	mutation := newCRYPTOCURRENCYMutation(c.config, OpUpdateOne, withCRYPTO_CURRENCY(cc))
	return &CRYPTOCURRENCYUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CRYPTO_CURRENCYClient) UpdateOneID(id int) *CRYPTOCURRENCYUpdateOne {
	mutation := newCRYPTOCURRENCYMutation(c.config, OpUpdateOne, withCRYPTO_CURRENCYID(id))
	return &CRYPTOCURRENCYUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CRYPTO_CURRENCY.
func (c *CRYPTO_CURRENCYClient) Delete() *CRYPTOCURRENCYDelete {
	mutation := newCRYPTOCURRENCYMutation(c.config, OpDelete)
	return &CRYPTOCURRENCYDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CRYPTO_CURRENCYClient) DeleteOne(cc *CRYPTO_CURRENCY) *CRYPTOCURRENCYDeleteOne {
	return c.DeleteOneID(cc.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CRYPTO_CURRENCYClient) DeleteOneID(id int) *CRYPTOCURRENCYDeleteOne {
	builder := c.Delete().Where(crypto_currency.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CRYPTOCURRENCYDeleteOne{builder}
}

// Query returns a query builder for CRYPTO_CURRENCY.
func (c *CRYPTO_CURRENCYClient) Query() *CRYPTOCURRENCYQuery {
	return &CRYPTOCURRENCYQuery{
		config: c.config,
	}
}

// Get returns a CRYPTO_CURRENCY entity by its id.
func (c *CRYPTO_CURRENCYClient) Get(ctx context.Context, id int) (*CRYPTO_CURRENCY, error) {
	return c.Query().Where(crypto_currency.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CRYPTO_CURRENCYClient) GetX(ctx context.Context, id int) *CRYPTO_CURRENCY {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySourceOf queries the source_of edge of a CRYPTO_CURRENCY.
func (c *CRYPTO_CURRENCYClient) QuerySourceOf(cc *CRYPTO_CURRENCY) *CRYPTOPRCSOURCEQuery {
	query := &CRYPTOPRCSOURCEQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(crypto_currency.Table, crypto_currency.FieldID, id),
			sqlgraph.To(crypto_prc_source.Table, crypto_prc_source.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, crypto_currency.SourceOfTable, crypto_currency.SourceOfColumn),
		)
		fromV = sqlgraph.Neighbors(cc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEmployeePaid queries the employee_paid edge of a CRYPTO_CURRENCY.
func (c *CRYPTO_CURRENCYClient) QueryEmployeePaid(cc *CRYPTO_CURRENCY) *EMPLOYEEQuery {
	query := &EMPLOYEEQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(crypto_currency.Table, crypto_currency.FieldID, id),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, crypto_currency.EmployeePaidTable, crypto_currency.EmployeePaidColumn),
		)
		fromV = sqlgraph.Neighbors(cc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CRYPTO_CURRENCYClient) Hooks() []Hook {
	return c.hooks.CRYPTO_CURRENCY
}

// CRYPTO_PRC_SOURCEClient is a client for the CRYPTO_PRC_SOURCE schema.
type CRYPTO_PRC_SOURCEClient struct {
	config
}

// NewCRYPTO_PRC_SOURCEClient returns a client for the CRYPTO_PRC_SOURCE from the given config.
func NewCRYPTO_PRC_SOURCEClient(c config) *CRYPTO_PRC_SOURCEClient {
	return &CRYPTO_PRC_SOURCEClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `crypto_prc_source.Hooks(f(g(h())))`.
func (c *CRYPTO_PRC_SOURCEClient) Use(hooks ...Hook) {
	c.hooks.CRYPTO_PRC_SOURCE = append(c.hooks.CRYPTO_PRC_SOURCE, hooks...)
}

// Create returns a builder for creating a CRYPTO_PRC_SOURCE entity.
func (c *CRYPTO_PRC_SOURCEClient) Create() *CRYPTOPRCSOURCECreate {
	mutation := newCRYPTOPRCSOURCEMutation(c.config, OpCreate)
	return &CRYPTOPRCSOURCECreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CRYPTO_PRC_SOURCE entities.
func (c *CRYPTO_PRC_SOURCEClient) CreateBulk(builders ...*CRYPTOPRCSOURCECreate) *CRYPTOPRCSOURCECreateBulk {
	return &CRYPTOPRCSOURCECreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CRYPTO_PRC_SOURCE.
func (c *CRYPTO_PRC_SOURCEClient) Update() *CRYPTOPRCSOURCEUpdate {
	mutation := newCRYPTOPRCSOURCEMutation(c.config, OpUpdate)
	return &CRYPTOPRCSOURCEUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CRYPTO_PRC_SOURCEClient) UpdateOne(cps *CRYPTO_PRC_SOURCE) *CRYPTOPRCSOURCEUpdateOne {
	mutation := newCRYPTOPRCSOURCEMutation(c.config, OpUpdateOne, withCRYPTO_PRC_SOURCE(cps))
	return &CRYPTOPRCSOURCEUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CRYPTO_PRC_SOURCEClient) UpdateOneID(id int) *CRYPTOPRCSOURCEUpdateOne {
	mutation := newCRYPTOPRCSOURCEMutation(c.config, OpUpdateOne, withCRYPTO_PRC_SOURCEID(id))
	return &CRYPTOPRCSOURCEUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CRYPTO_PRC_SOURCE.
func (c *CRYPTO_PRC_SOURCEClient) Delete() *CRYPTOPRCSOURCEDelete {
	mutation := newCRYPTOPRCSOURCEMutation(c.config, OpDelete)
	return &CRYPTOPRCSOURCEDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CRYPTO_PRC_SOURCEClient) DeleteOne(cps *CRYPTO_PRC_SOURCE) *CRYPTOPRCSOURCEDeleteOne {
	return c.DeleteOneID(cps.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CRYPTO_PRC_SOURCEClient) DeleteOneID(id int) *CRYPTOPRCSOURCEDeleteOne {
	builder := c.Delete().Where(crypto_prc_source.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CRYPTOPRCSOURCEDeleteOne{builder}
}

// Query returns a query builder for CRYPTO_PRC_SOURCE.
func (c *CRYPTO_PRC_SOURCEClient) Query() *CRYPTOPRCSOURCEQuery {
	return &CRYPTOPRCSOURCEQuery{
		config: c.config,
	}
}

// Get returns a CRYPTO_PRC_SOURCE entity by its id.
func (c *CRYPTO_PRC_SOURCEClient) Get(ctx context.Context, id int) (*CRYPTO_PRC_SOURCE, error) {
	return c.Query().Where(crypto_prc_source.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CRYPTO_PRC_SOURCEClient) GetX(ctx context.Context, id int) *CRYPTO_PRC_SOURCE {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPriceOf queries the price_of edge of a CRYPTO_PRC_SOURCE.
func (c *CRYPTO_PRC_SOURCEClient) QueryPriceOf(cps *CRYPTO_PRC_SOURCE) *CRYPTOCURRENCYQuery {
	query := &CRYPTOCURRENCYQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cps.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(crypto_prc_source.Table, crypto_prc_source.FieldID, id),
			sqlgraph.To(crypto_currency.Table, crypto_currency.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, crypto_prc_source.PriceOfTable, crypto_prc_source.PriceOfColumn),
		)
		fromV = sqlgraph.Neighbors(cps.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CRYPTO_PRC_SOURCEClient) Hooks() []Hook {
	return c.hooks.CRYPTO_PRC_SOURCE
}

// EMPLOYEEClient is a client for the EMPLOYEE schema.
type EMPLOYEEClient struct {
	config
}

// NewEMPLOYEEClient returns a client for the EMPLOYEE from the given config.
func NewEMPLOYEEClient(c config) *EMPLOYEEClient {
	return &EMPLOYEEClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `employee.Hooks(f(g(h())))`.
func (c *EMPLOYEEClient) Use(hooks ...Hook) {
	c.hooks.EMPLOYEE = append(c.hooks.EMPLOYEE, hooks...)
}

// Create returns a builder for creating a EMPLOYEE entity.
func (c *EMPLOYEEClient) Create() *EMPLOYEECreate {
	mutation := newEMPLOYEEMutation(c.config, OpCreate)
	return &EMPLOYEECreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EMPLOYEE entities.
func (c *EMPLOYEEClient) CreateBulk(builders ...*EMPLOYEECreate) *EMPLOYEECreateBulk {
	return &EMPLOYEECreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EMPLOYEE.
func (c *EMPLOYEEClient) Update() *EMPLOYEEUpdate {
	mutation := newEMPLOYEEMutation(c.config, OpUpdate)
	return &EMPLOYEEUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EMPLOYEEClient) UpdateOne(e *EMPLOYEE) *EMPLOYEEUpdateOne {
	mutation := newEMPLOYEEMutation(c.config, OpUpdateOne, withEMPLOYEE(e))
	return &EMPLOYEEUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EMPLOYEEClient) UpdateOneID(id int) *EMPLOYEEUpdateOne {
	mutation := newEMPLOYEEMutation(c.config, OpUpdateOne, withEMPLOYEEID(id))
	return &EMPLOYEEUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EMPLOYEE.
func (c *EMPLOYEEClient) Delete() *EMPLOYEEDelete {
	mutation := newEMPLOYEEMutation(c.config, OpDelete)
	return &EMPLOYEEDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EMPLOYEEClient) DeleteOne(e *EMPLOYEE) *EMPLOYEEDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *EMPLOYEEClient) DeleteOneID(id int) *EMPLOYEEDeleteOne {
	builder := c.Delete().Where(employee.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EMPLOYEEDeleteOne{builder}
}

// Query returns a query builder for EMPLOYEE.
func (c *EMPLOYEEClient) Query() *EMPLOYEEQuery {
	return &EMPLOYEEQuery{
		config: c.config,
	}
}

// Get returns a EMPLOYEE entity by its id.
func (c *EMPLOYEEClient) Get(ctx context.Context, id int) (*EMPLOYEE, error) {
	return c.Query().Where(employee.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EMPLOYEEClient) GetX(ctx context.Context, id int) *EMPLOYEE {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEmployeeGets queries the employee_gets edge of a EMPLOYEE.
func (c *EMPLOYEEClient) QueryEmployeeGets(e *EMPLOYEE) *CRYPTOCURRENCYQuery {
	query := &CRYPTOCURRENCYQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(employee.Table, employee.FieldID, id),
			sqlgraph.To(crypto_currency.Table, crypto_currency.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, employee.EmployeeGetsTable, employee.EmployeeGetsColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEmployeeTypeFrom queries the employee_type_from edge of a EMPLOYEE.
func (c *EMPLOYEEClient) QueryEmployeeTypeFrom(e *EMPLOYEE) *EMPLOYTYPEQuery {
	query := &EMPLOYTYPEQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(employee.Table, employee.FieldID, id),
			sqlgraph.To(employ_type.Table, employ_type.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, employee.EmployeeTypeFromTable, employee.EmployeeTypeFromColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWorkFor queries the work_for edge of a EMPLOYEE.
func (c *EMPLOYEEClient) QueryWorkFor(e *EMPLOYEE) *EMPLOYERUSERINFOQuery {
	query := &EMPLOYERUSERINFOQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(employee.Table, employee.FieldID, id),
			sqlgraph.To(employer_user_info.Table, employer_user_info.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, employee.WorkForTable, employee.WorkForColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPaymentHistory queries the payment_history edge of a EMPLOYEE.
func (c *EMPLOYEEClient) QueryPaymentHistory(e *EMPLOYEE) *PAYMENTHISTORYQuery {
	query := &PAYMENTHISTORYQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(employee.Table, employee.FieldID, id),
			sqlgraph.To(payment_history.Table, payment_history.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, employee.PaymentHistoryTable, employee.PaymentHistoryColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EMPLOYEEClient) Hooks() []Hook {
	return c.hooks.EMPLOYEE
}

// EMPLOYER_USER_INFOClient is a client for the EMPLOYER_USER_INFO schema.
type EMPLOYER_USER_INFOClient struct {
	config
}

// NewEMPLOYER_USER_INFOClient returns a client for the EMPLOYER_USER_INFO from the given config.
func NewEMPLOYER_USER_INFOClient(c config) *EMPLOYER_USER_INFOClient {
	return &EMPLOYER_USER_INFOClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `employer_user_info.Hooks(f(g(h())))`.
func (c *EMPLOYER_USER_INFOClient) Use(hooks ...Hook) {
	c.hooks.EMPLOYER_USER_INFO = append(c.hooks.EMPLOYER_USER_INFO, hooks...)
}

// Create returns a builder for creating a EMPLOYER_USER_INFO entity.
func (c *EMPLOYER_USER_INFOClient) Create() *EMPLOYERUSERINFOCreate {
	mutation := newEMPLOYERUSERINFOMutation(c.config, OpCreate)
	return &EMPLOYERUSERINFOCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EMPLOYER_USER_INFO entities.
func (c *EMPLOYER_USER_INFOClient) CreateBulk(builders ...*EMPLOYERUSERINFOCreate) *EMPLOYERUSERINFOCreateBulk {
	return &EMPLOYERUSERINFOCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EMPLOYER_USER_INFO.
func (c *EMPLOYER_USER_INFOClient) Update() *EMPLOYERUSERINFOUpdate {
	mutation := newEMPLOYERUSERINFOMutation(c.config, OpUpdate)
	return &EMPLOYERUSERINFOUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EMPLOYER_USER_INFOClient) UpdateOne(eui *EMPLOYER_USER_INFO) *EMPLOYERUSERINFOUpdateOne {
	mutation := newEMPLOYERUSERINFOMutation(c.config, OpUpdateOne, withEMPLOYER_USER_INFO(eui))
	return &EMPLOYERUSERINFOUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EMPLOYER_USER_INFOClient) UpdateOneID(id int) *EMPLOYERUSERINFOUpdateOne {
	mutation := newEMPLOYERUSERINFOMutation(c.config, OpUpdateOne, withEMPLOYER_USER_INFOID(id))
	return &EMPLOYERUSERINFOUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EMPLOYER_USER_INFO.
func (c *EMPLOYER_USER_INFOClient) Delete() *EMPLOYERUSERINFODelete {
	mutation := newEMPLOYERUSERINFOMutation(c.config, OpDelete)
	return &EMPLOYERUSERINFODelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EMPLOYER_USER_INFOClient) DeleteOne(eui *EMPLOYER_USER_INFO) *EMPLOYERUSERINFODeleteOne {
	return c.DeleteOneID(eui.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *EMPLOYER_USER_INFOClient) DeleteOneID(id int) *EMPLOYERUSERINFODeleteOne {
	builder := c.Delete().Where(employer_user_info.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EMPLOYERUSERINFODeleteOne{builder}
}

// Query returns a query builder for EMPLOYER_USER_INFO.
func (c *EMPLOYER_USER_INFOClient) Query() *EMPLOYERUSERINFOQuery {
	return &EMPLOYERUSERINFOQuery{
		config: c.config,
	}
}

// Get returns a EMPLOYER_USER_INFO entity by its id.
func (c *EMPLOYER_USER_INFOClient) Get(ctx context.Context, id int) (*EMPLOYER_USER_INFO, error) {
	return c.Query().Where(employer_user_info.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EMPLOYER_USER_INFOClient) GetX(ctx context.Context, id int) *EMPLOYER_USER_INFO {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWorkUnder queries the work_under edge of a EMPLOYER_USER_INFO.
func (c *EMPLOYER_USER_INFOClient) QueryWorkUnder(eui *EMPLOYER_USER_INFO) *EMPLOYEEQuery {
	query := &EMPLOYEEQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := eui.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(employer_user_info.Table, employer_user_info.FieldID, id),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, employer_user_info.WorkUnderTable, employer_user_info.WorkUnderColumn),
		)
		fromV = sqlgraph.Neighbors(eui.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EMPLOYER_USER_INFOClient) Hooks() []Hook {
	return c.hooks.EMPLOYER_USER_INFO
}

// EMPLOY_TYPEClient is a client for the EMPLOY_TYPE schema.
type EMPLOY_TYPEClient struct {
	config
}

// NewEMPLOY_TYPEClient returns a client for the EMPLOY_TYPE from the given config.
func NewEMPLOY_TYPEClient(c config) *EMPLOY_TYPEClient {
	return &EMPLOY_TYPEClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `employ_type.Hooks(f(g(h())))`.
func (c *EMPLOY_TYPEClient) Use(hooks ...Hook) {
	c.hooks.EMPLOY_TYPE = append(c.hooks.EMPLOY_TYPE, hooks...)
}

// Create returns a builder for creating a EMPLOY_TYPE entity.
func (c *EMPLOY_TYPEClient) Create() *EMPLOYTYPECreate {
	mutation := newEMPLOYTYPEMutation(c.config, OpCreate)
	return &EMPLOYTYPECreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EMPLOY_TYPE entities.
func (c *EMPLOY_TYPEClient) CreateBulk(builders ...*EMPLOYTYPECreate) *EMPLOYTYPECreateBulk {
	return &EMPLOYTYPECreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EMPLOY_TYPE.
func (c *EMPLOY_TYPEClient) Update() *EMPLOYTYPEUpdate {
	mutation := newEMPLOYTYPEMutation(c.config, OpUpdate)
	return &EMPLOYTYPEUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EMPLOY_TYPEClient) UpdateOne(et *EMPLOY_TYPE) *EMPLOYTYPEUpdateOne {
	mutation := newEMPLOYTYPEMutation(c.config, OpUpdateOne, withEMPLOY_TYPE(et))
	return &EMPLOYTYPEUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EMPLOY_TYPEClient) UpdateOneID(id int) *EMPLOYTYPEUpdateOne {
	mutation := newEMPLOYTYPEMutation(c.config, OpUpdateOne, withEMPLOY_TYPEID(id))
	return &EMPLOYTYPEUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EMPLOY_TYPE.
func (c *EMPLOY_TYPEClient) Delete() *EMPLOYTYPEDelete {
	mutation := newEMPLOYTYPEMutation(c.config, OpDelete)
	return &EMPLOYTYPEDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EMPLOY_TYPEClient) DeleteOne(et *EMPLOY_TYPE) *EMPLOYTYPEDeleteOne {
	return c.DeleteOneID(et.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *EMPLOY_TYPEClient) DeleteOneID(id int) *EMPLOYTYPEDeleteOne {
	builder := c.Delete().Where(employ_type.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EMPLOYTYPEDeleteOne{builder}
}

// Query returns a query builder for EMPLOY_TYPE.
func (c *EMPLOY_TYPEClient) Query() *EMPLOYTYPEQuery {
	return &EMPLOYTYPEQuery{
		config: c.config,
	}
}

// Get returns a EMPLOY_TYPE entity by its id.
func (c *EMPLOY_TYPEClient) Get(ctx context.Context, id int) (*EMPLOY_TYPE, error) {
	return c.Query().Where(employ_type.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EMPLOY_TYPEClient) GetX(ctx context.Context, id int) *EMPLOY_TYPE {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEmployeeTypeTo queries the employee_type_to edge of a EMPLOY_TYPE.
func (c *EMPLOY_TYPEClient) QueryEmployeeTypeTo(et *EMPLOY_TYPE) *EMPLOYEEQuery {
	query := &EMPLOYEEQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := et.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(employ_type.Table, employ_type.FieldID, id),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, employ_type.EmployeeTypeToTable, employ_type.EmployeeTypeToColumn),
		)
		fromV = sqlgraph.Neighbors(et.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EMPLOY_TYPEClient) Hooks() []Hook {
	return c.hooks.EMPLOY_TYPE
}

// PAYMENT_HISTORYClient is a client for the PAYMENT_HISTORY schema.
type PAYMENT_HISTORYClient struct {
	config
}

// NewPAYMENT_HISTORYClient returns a client for the PAYMENT_HISTORY from the given config.
func NewPAYMENT_HISTORYClient(c config) *PAYMENT_HISTORYClient {
	return &PAYMENT_HISTORYClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `payment_history.Hooks(f(g(h())))`.
func (c *PAYMENT_HISTORYClient) Use(hooks ...Hook) {
	c.hooks.PAYMENT_HISTORY = append(c.hooks.PAYMENT_HISTORY, hooks...)
}

// Create returns a builder for creating a PAYMENT_HISTORY entity.
func (c *PAYMENT_HISTORYClient) Create() *PAYMENTHISTORYCreate {
	mutation := newPAYMENTHISTORYMutation(c.config, OpCreate)
	return &PAYMENTHISTORYCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PAYMENT_HISTORY entities.
func (c *PAYMENT_HISTORYClient) CreateBulk(builders ...*PAYMENTHISTORYCreate) *PAYMENTHISTORYCreateBulk {
	return &PAYMENTHISTORYCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PAYMENT_HISTORY.
func (c *PAYMENT_HISTORYClient) Update() *PAYMENTHISTORYUpdate {
	mutation := newPAYMENTHISTORYMutation(c.config, OpUpdate)
	return &PAYMENTHISTORYUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PAYMENT_HISTORYClient) UpdateOne(ph *PAYMENT_HISTORY) *PAYMENTHISTORYUpdateOne {
	mutation := newPAYMENTHISTORYMutation(c.config, OpUpdateOne, withPAYMENT_HISTORY(ph))
	return &PAYMENTHISTORYUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PAYMENT_HISTORYClient) UpdateOneID(id int) *PAYMENTHISTORYUpdateOne {
	mutation := newPAYMENTHISTORYMutation(c.config, OpUpdateOne, withPAYMENT_HISTORYID(id))
	return &PAYMENTHISTORYUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PAYMENT_HISTORY.
func (c *PAYMENT_HISTORYClient) Delete() *PAYMENTHISTORYDelete {
	mutation := newPAYMENTHISTORYMutation(c.config, OpDelete)
	return &PAYMENTHISTORYDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PAYMENT_HISTORYClient) DeleteOne(ph *PAYMENT_HISTORY) *PAYMENTHISTORYDeleteOne {
	return c.DeleteOneID(ph.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PAYMENT_HISTORYClient) DeleteOneID(id int) *PAYMENTHISTORYDeleteOne {
	builder := c.Delete().Where(payment_history.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PAYMENTHISTORYDeleteOne{builder}
}

// Query returns a query builder for PAYMENT_HISTORY.
func (c *PAYMENT_HISTORYClient) Query() *PAYMENTHISTORYQuery {
	return &PAYMENTHISTORYQuery{
		config: c.config,
	}
}

// Get returns a PAYMENT_HISTORY entity by its id.
func (c *PAYMENT_HISTORYClient) Get(ctx context.Context, id int) (*PAYMENT_HISTORY, error) {
	return c.Query().Where(payment_history.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PAYMENT_HISTORYClient) GetX(ctx context.Context, id int) *PAYMENT_HISTORY {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPaymentHistoryRec queries the payment_history_rec edge of a PAYMENT_HISTORY.
func (c *PAYMENT_HISTORYClient) QueryPaymentHistoryRec(ph *PAYMENT_HISTORY) *EMPLOYEEQuery {
	query := &EMPLOYEEQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ph.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(payment_history.Table, payment_history.FieldID, id),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, payment_history.PaymentHistoryRecTable, payment_history.PaymentHistoryRecColumn),
		)
		fromV = sqlgraph.Neighbors(ph.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PAYMENT_HISTORYClient) Hooks() []Hook {
	return c.hooks.PAYMENT_HISTORY
}

// Tr_logClient is a client for the Tr_log schema.
type Tr_logClient struct {
	config
}

// NewTr_logClient returns a client for the Tr_log from the given config.
func NewTr_logClient(c config) *Tr_logClient {
	return &Tr_logClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tr_log.Hooks(f(g(h())))`.
func (c *Tr_logClient) Use(hooks ...Hook) {
	c.hooks.Tr_log = append(c.hooks.Tr_log, hooks...)
}

// Create returns a builder for creating a Tr_log entity.
func (c *Tr_logClient) Create() *TrLogCreate {
	mutation := newTrLogMutation(c.config, OpCreate)
	return &TrLogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tr_log entities.
func (c *Tr_logClient) CreateBulk(builders ...*TrLogCreate) *TrLogCreateBulk {
	return &TrLogCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tr_log.
func (c *Tr_logClient) Update() *TrLogUpdate {
	mutation := newTrLogMutation(c.config, OpUpdate)
	return &TrLogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *Tr_logClient) UpdateOne(tl *Tr_log) *TrLogUpdateOne {
	mutation := newTrLogMutation(c.config, OpUpdateOne, withTr_log(tl))
	return &TrLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *Tr_logClient) UpdateOneID(id int) *TrLogUpdateOne {
	mutation := newTrLogMutation(c.config, OpUpdateOne, withTr_logID(id))
	return &TrLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tr_log.
func (c *Tr_logClient) Delete() *TrLogDelete {
	mutation := newTrLogMutation(c.config, OpDelete)
	return &TrLogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *Tr_logClient) DeleteOne(tl *Tr_log) *TrLogDeleteOne {
	return c.DeleteOneID(tl.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *Tr_logClient) DeleteOneID(id int) *TrLogDeleteOne {
	builder := c.Delete().Where(tr_log.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TrLogDeleteOne{builder}
}

// Query returns a query builder for Tr_log.
func (c *Tr_logClient) Query() *TrLogQuery {
	return &TrLogQuery{
		config: c.config,
	}
}

// Get returns a Tr_log entity by its id.
func (c *Tr_logClient) Get(ctx context.Context, id int) (*Tr_log, error) {
	return c.Query().Where(tr_log.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *Tr_logClient) GetX(ctx context.Context, id int) *Tr_log {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *Tr_logClient) Hooks() []Hook {
	return c.hooks.Tr_log
}
