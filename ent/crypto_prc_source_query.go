// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"griffin-dao/ent/crypto_prc_source"
	"griffin-dao/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CRYPTOPRCSOURCEQuery is the builder for querying CRYPTO_PRC_SOURCE entities.
type CRYPTOPRCSOURCEQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CRYPTO_PRC_SOURCE
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CRYPTOPRCSOURCEQuery builder.
func (cq *CRYPTOPRCSOURCEQuery) Where(ps ...predicate.CRYPTO_PRC_SOURCE) *CRYPTOPRCSOURCEQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CRYPTOPRCSOURCEQuery) Limit(limit int) *CRYPTOPRCSOURCEQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CRYPTOPRCSOURCEQuery) Offset(offset int) *CRYPTOPRCSOURCEQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CRYPTOPRCSOURCEQuery) Unique(unique bool) *CRYPTOPRCSOURCEQuery {
	cq.unique = &unique
	return cq
}

// Order adds an order step to the query.
func (cq *CRYPTOPRCSOURCEQuery) Order(o ...OrderFunc) *CRYPTOPRCSOURCEQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// First returns the first CRYPTO_PRC_SOURCE entity from the query.
// Returns a *NotFoundError when no CRYPTO_PRC_SOURCE was found.
func (cq *CRYPTOPRCSOURCEQuery) First(ctx context.Context) (*CRYPTO_PRC_SOURCE, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{crypto_prc_source.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CRYPTOPRCSOURCEQuery) FirstX(ctx context.Context) *CRYPTO_PRC_SOURCE {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CRYPTO_PRC_SOURCE ID from the query.
// Returns a *NotFoundError when no CRYPTO_PRC_SOURCE ID was found.
func (cq *CRYPTOPRCSOURCEQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{crypto_prc_source.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CRYPTOPRCSOURCEQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CRYPTO_PRC_SOURCE entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CRYPTO_PRC_SOURCE entity is found.
// Returns a *NotFoundError when no CRYPTO_PRC_SOURCE entities are found.
func (cq *CRYPTOPRCSOURCEQuery) Only(ctx context.Context) (*CRYPTO_PRC_SOURCE, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{crypto_prc_source.Label}
	default:
		return nil, &NotSingularError{crypto_prc_source.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CRYPTOPRCSOURCEQuery) OnlyX(ctx context.Context) *CRYPTO_PRC_SOURCE {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CRYPTO_PRC_SOURCE ID in the query.
// Returns a *NotSingularError when more than one CRYPTO_PRC_SOURCE ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CRYPTOPRCSOURCEQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{crypto_prc_source.Label}
	default:
		err = &NotSingularError{crypto_prc_source.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CRYPTOPRCSOURCEQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CRYPTO_PRC_SOURCEs.
func (cq *CRYPTOPRCSOURCEQuery) All(ctx context.Context) ([]*CRYPTO_PRC_SOURCE, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CRYPTOPRCSOURCEQuery) AllX(ctx context.Context) []*CRYPTO_PRC_SOURCE {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CRYPTO_PRC_SOURCE IDs.
func (cq *CRYPTOPRCSOURCEQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cq.Select(crypto_prc_source.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CRYPTOPRCSOURCEQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CRYPTOPRCSOURCEQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CRYPTOPRCSOURCEQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CRYPTOPRCSOURCEQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CRYPTOPRCSOURCEQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CRYPTOPRCSOURCEQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CRYPTOPRCSOURCEQuery) Clone() *CRYPTOPRCSOURCEQuery {
	if cq == nil {
		return nil
	}
	return &CRYPTOPRCSOURCEQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]OrderFunc{}, cq.order...),
		predicates: append([]predicate.CRYPTO_PRC_SOURCE{}, cq.predicates...),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CRYPTOPRCSOURCE.Query().
//		GroupBy(crypto_prc_source.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *CRYPTOPRCSOURCEQuery) GroupBy(field string, fields ...string) *CRYPTOPRCSOURCEGroupBy {
	grbuild := &CRYPTOPRCSOURCEGroupBy{config: cq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(ctx), nil
	}
	grbuild.label = crypto_prc_source.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.CRYPTOPRCSOURCE.Query().
//		Select(crypto_prc_source.FieldName).
//		Scan(ctx, &v)
//
func (cq *CRYPTOPRCSOURCEQuery) Select(fields ...string) *CRYPTOPRCSOURCESelect {
	cq.fields = append(cq.fields, fields...)
	selbuild := &CRYPTOPRCSOURCESelect{CRYPTOPRCSOURCEQuery: cq}
	selbuild.label = crypto_prc_source.Label
	selbuild.flds, selbuild.scan = &cq.fields, selbuild.Scan
	return selbuild
}

func (cq *CRYPTOPRCSOURCEQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !crypto_prc_source.ValidColumn(f) {
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

func (cq *CRYPTOPRCSOURCEQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CRYPTO_PRC_SOURCE, error) {
	var (
		nodes   = []*CRYPTO_PRC_SOURCE{}
		withFKs = cq.withFKs
		_spec   = cq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, crypto_prc_source.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CRYPTO_PRC_SOURCE).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CRYPTO_PRC_SOURCE{config: cq.config}
		nodes = append(nodes, node)
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
	return nodes, nil
}

func (cq *CRYPTOPRCSOURCEQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CRYPTOPRCSOURCEQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cq *CRYPTOPRCSOURCEQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   crypto_prc_source.Table,
			Columns: crypto_prc_source.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: crypto_prc_source.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, crypto_prc_source.FieldID)
		for i := range fields {
			if fields[i] != crypto_prc_source.FieldID {
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

func (cq *CRYPTOPRCSOURCEQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(crypto_prc_source.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = crypto_prc_source.Columns
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

// CRYPTOPRCSOURCEGroupBy is the group-by builder for CRYPTO_PRC_SOURCE entities.
type CRYPTOPRCSOURCEGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CRYPTOPRCSOURCEGroupBy) Aggregate(fns ...AggregateFunc) *CRYPTOPRCSOURCEGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *CRYPTOPRCSOURCEGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

func (cgb *CRYPTOPRCSOURCEGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cgb.fields {
		if !crypto_prc_source.ValidColumn(f) {
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

func (cgb *CRYPTOPRCSOURCEGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql.Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
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

// CRYPTOPRCSOURCESelect is the builder for selecting fields of CRYPTOPRCSOURCE entities.
type CRYPTOPRCSOURCESelect struct {
	*CRYPTOPRCSOURCEQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CRYPTOPRCSOURCESelect) Scan(ctx context.Context, v interface{}) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.CRYPTOPRCSOURCEQuery.sqlQuery(ctx)
	return cs.sqlScan(ctx, v)
}

func (cs *CRYPTOPRCSOURCESelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cs.sql.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
