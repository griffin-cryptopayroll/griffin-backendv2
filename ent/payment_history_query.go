// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/payment_history"
	"griffin-dao/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PAYMENTHISTORYQuery is the builder for querying PAYMENT_HISTORY entities.
type PAYMENTHISTORYQuery struct {
	config
	limit                          *int
	offset                         *int
	unique                         *bool
	order                          []OrderFunc
	fields                         []string
	predicates                     []predicate.PAYMENT_HISTORY
	withPaymentHistoryFromEmployee *EMPLOYEEQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PAYMENTHISTORYQuery builder.
func (pq *PAYMENTHISTORYQuery) Where(ps ...predicate.PAYMENT_HISTORY) *PAYMENTHISTORYQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PAYMENTHISTORYQuery) Limit(limit int) *PAYMENTHISTORYQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PAYMENTHISTORYQuery) Offset(offset int) *PAYMENTHISTORYQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PAYMENTHISTORYQuery) Unique(unique bool) *PAYMENTHISTORYQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *PAYMENTHISTORYQuery) Order(o ...OrderFunc) *PAYMENTHISTORYQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryPaymentHistoryFromEmployee chains the current query on the "payment_history_from_employee" edge.
func (pq *PAYMENTHISTORYQuery) QueryPaymentHistoryFromEmployee() *EMPLOYEEQuery {
	query := &EMPLOYEEQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(payment_history.Table, payment_history.FieldID, selector),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, payment_history.PaymentHistoryFromEmployeeTable, payment_history.PaymentHistoryFromEmployeeColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PAYMENT_HISTORY entity from the query.
// Returns a *NotFoundError when no PAYMENT_HISTORY was found.
func (pq *PAYMENTHISTORYQuery) First(ctx context.Context) (*PAYMENT_HISTORY, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{payment_history.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PAYMENTHISTORYQuery) FirstX(ctx context.Context) *PAYMENT_HISTORY {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PAYMENT_HISTORY ID from the query.
// Returns a *NotFoundError when no PAYMENT_HISTORY ID was found.
func (pq *PAYMENTHISTORYQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{payment_history.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PAYMENTHISTORYQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PAYMENT_HISTORY entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PAYMENT_HISTORY entity is found.
// Returns a *NotFoundError when no PAYMENT_HISTORY entities are found.
func (pq *PAYMENTHISTORYQuery) Only(ctx context.Context) (*PAYMENT_HISTORY, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{payment_history.Label}
	default:
		return nil, &NotSingularError{payment_history.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PAYMENTHISTORYQuery) OnlyX(ctx context.Context) *PAYMENT_HISTORY {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PAYMENT_HISTORY ID in the query.
// Returns a *NotSingularError when more than one PAYMENT_HISTORY ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PAYMENTHISTORYQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{payment_history.Label}
	default:
		err = &NotSingularError{payment_history.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PAYMENTHISTORYQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PAYMENT_HISTORYs.
func (pq *PAYMENTHISTORYQuery) All(ctx context.Context) ([]*PAYMENT_HISTORY, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PAYMENTHISTORYQuery) AllX(ctx context.Context) []*PAYMENT_HISTORY {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PAYMENT_HISTORY IDs.
func (pq *PAYMENTHISTORYQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(payment_history.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PAYMENTHISTORYQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PAYMENTHISTORYQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PAYMENTHISTORYQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PAYMENTHISTORYQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PAYMENTHISTORYQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PAYMENTHISTORYQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PAYMENTHISTORYQuery) Clone() *PAYMENTHISTORYQuery {
	if pq == nil {
		return nil
	}
	return &PAYMENTHISTORYQuery{
		config:                         pq.config,
		limit:                          pq.limit,
		offset:                         pq.offset,
		order:                          append([]OrderFunc{}, pq.order...),
		predicates:                     append([]predicate.PAYMENT_HISTORY{}, pq.predicates...),
		withPaymentHistoryFromEmployee: pq.withPaymentHistoryFromEmployee.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithPaymentHistoryFromEmployee tells the query-builder to eager-load the nodes that are connected to
// the "payment_history_from_employee" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PAYMENTHISTORYQuery) WithPaymentHistoryFromEmployee(opts ...func(*EMPLOYEEQuery)) *PAYMENTHISTORYQuery {
	query := &EMPLOYEEQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withPaymentHistoryFromEmployee = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EmployeeID int `json:"employee_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PAYMENTHISTORY.Query().
//		GroupBy(payment_history.FieldEmployeeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PAYMENTHISTORYQuery) GroupBy(field string, fields ...string) *PAYMENTHISTORYGroupBy {
	grbuild := &PAYMENTHISTORYGroupBy{config: pq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	grbuild.label = payment_history.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EmployeeID int `json:"employee_id,omitempty"`
//	}
//
//	client.PAYMENTHISTORY.Query().
//		Select(payment_history.FieldEmployeeID).
//		Scan(ctx, &v)
func (pq *PAYMENTHISTORYQuery) Select(fields ...string) *PAYMENTHISTORYSelect {
	pq.fields = append(pq.fields, fields...)
	selbuild := &PAYMENTHISTORYSelect{PAYMENTHISTORYQuery: pq}
	selbuild.label = payment_history.Label
	selbuild.flds, selbuild.scan = &pq.fields, selbuild.Scan
	return selbuild
}

func (pq *PAYMENTHISTORYQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !payment_history.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PAYMENTHISTORYQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PAYMENT_HISTORY, error) {
	var (
		nodes       = []*PAYMENT_HISTORY{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withPaymentHistoryFromEmployee != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PAYMENT_HISTORY).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PAYMENT_HISTORY{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withPaymentHistoryFromEmployee; query != nil {
		if err := pq.loadPaymentHistoryFromEmployee(ctx, query, nodes, nil,
			func(n *PAYMENT_HISTORY, e *EMPLOYEE) { n.Edges.PaymentHistoryFromEmployee = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PAYMENTHISTORYQuery) loadPaymentHistoryFromEmployee(ctx context.Context, query *EMPLOYEEQuery, nodes []*PAYMENT_HISTORY, init func(*PAYMENT_HISTORY), assign func(*PAYMENT_HISTORY, *EMPLOYEE)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PAYMENT_HISTORY)
	for i := range nodes {
		fk := nodes[i].EmployeeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(employee.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "employee_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PAYMENTHISTORYQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PAYMENTHISTORYQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pq *PAYMENTHISTORYQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payment_history.Table,
			Columns: payment_history.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payment_history.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment_history.FieldID)
		for i := range fields {
			if fields[i] != payment_history.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PAYMENTHISTORYQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(payment_history.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = payment_history.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PAYMENTHISTORYGroupBy is the group-by builder for PAYMENT_HISTORY entities.
type PAYMENTHISTORYGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PAYMENTHISTORYGroupBy) Aggregate(fns ...AggregateFunc) *PAYMENTHISTORYGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *PAYMENTHISTORYGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

func (pgb *PAYMENTHISTORYGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pgb.fields {
		if !payment_history.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PAYMENTHISTORYGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// PAYMENTHISTORYSelect is the builder for selecting fields of PAYMENTHISTORY entities.
type PAYMENTHISTORYSelect struct {
	*PAYMENTHISTORYQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PAYMENTHISTORYSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.PAYMENTHISTORYQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

func (ps *PAYMENTHISTORYSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
