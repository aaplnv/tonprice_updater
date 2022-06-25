// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/predicate"
	"main/ent/zarquote"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ZARQuoteQuery is the builder for querying ZARQuote entities.
type ZARQuoteQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ZARQuote
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ZARQuoteQuery builder.
func (zqq *ZARQuoteQuery) Where(ps ...predicate.ZARQuote) *ZARQuoteQuery {
	zqq.predicates = append(zqq.predicates, ps...)
	return zqq
}

// Limit adds a limit step to the query.
func (zqq *ZARQuoteQuery) Limit(limit int) *ZARQuoteQuery {
	zqq.limit = &limit
	return zqq
}

// Offset adds an offset step to the query.
func (zqq *ZARQuoteQuery) Offset(offset int) *ZARQuoteQuery {
	zqq.offset = &offset
	return zqq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (zqq *ZARQuoteQuery) Unique(unique bool) *ZARQuoteQuery {
	zqq.unique = &unique
	return zqq
}

// Order adds an order step to the query.
func (zqq *ZARQuoteQuery) Order(o ...OrderFunc) *ZARQuoteQuery {
	zqq.order = append(zqq.order, o...)
	return zqq
}

// First returns the first ZARQuote entity from the query.
// Returns a *NotFoundError when no ZARQuote was found.
func (zqq *ZARQuoteQuery) First(ctx context.Context) (*ZARQuote, error) {
	nodes, err := zqq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{zarquote.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (zqq *ZARQuoteQuery) FirstX(ctx context.Context) *ZARQuote {
	node, err := zqq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ZARQuote ID from the query.
// Returns a *NotFoundError when no ZARQuote ID was found.
func (zqq *ZARQuoteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = zqq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{zarquote.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (zqq *ZARQuoteQuery) FirstIDX(ctx context.Context) int {
	id, err := zqq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ZARQuote entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ZARQuote entity is found.
// Returns a *NotFoundError when no ZARQuote entities are found.
func (zqq *ZARQuoteQuery) Only(ctx context.Context) (*ZARQuote, error) {
	nodes, err := zqq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{zarquote.Label}
	default:
		return nil, &NotSingularError{zarquote.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (zqq *ZARQuoteQuery) OnlyX(ctx context.Context) *ZARQuote {
	node, err := zqq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ZARQuote ID in the query.
// Returns a *NotSingularError when more than one ZARQuote ID is found.
// Returns a *NotFoundError when no entities are found.
func (zqq *ZARQuoteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = zqq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{zarquote.Label}
	default:
		err = &NotSingularError{zarquote.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (zqq *ZARQuoteQuery) OnlyIDX(ctx context.Context) int {
	id, err := zqq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ZARQuotes.
func (zqq *ZARQuoteQuery) All(ctx context.Context) ([]*ZARQuote, error) {
	if err := zqq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return zqq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (zqq *ZARQuoteQuery) AllX(ctx context.Context) []*ZARQuote {
	nodes, err := zqq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ZARQuote IDs.
func (zqq *ZARQuoteQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := zqq.Select(zarquote.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (zqq *ZARQuoteQuery) IDsX(ctx context.Context) []int {
	ids, err := zqq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (zqq *ZARQuoteQuery) Count(ctx context.Context) (int, error) {
	if err := zqq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return zqq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (zqq *ZARQuoteQuery) CountX(ctx context.Context) int {
	count, err := zqq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (zqq *ZARQuoteQuery) Exist(ctx context.Context) (bool, error) {
	if err := zqq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return zqq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (zqq *ZARQuoteQuery) ExistX(ctx context.Context) bool {
	exist, err := zqq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ZARQuoteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (zqq *ZARQuoteQuery) Clone() *ZARQuoteQuery {
	if zqq == nil {
		return nil
	}
	return &ZARQuoteQuery{
		config:     zqq.config,
		limit:      zqq.limit,
		offset:     zqq.offset,
		order:      append([]OrderFunc{}, zqq.order...),
		predicates: append([]predicate.ZARQuote{}, zqq.predicates...),
		// clone intermediate query.
		sql:    zqq.sql.Clone(),
		path:   zqq.path,
		unique: zqq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Price float64 `json:"price,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ZARQuote.Query().
//		GroupBy(zarquote.FieldPrice).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (zqq *ZARQuoteQuery) GroupBy(field string, fields ...string) *ZARQuoteGroupBy {
	group := &ZARQuoteGroupBy{config: zqq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := zqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return zqq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Price float64 `json:"price,omitempty"`
//	}
//
//	client.ZARQuote.Query().
//		Select(zarquote.FieldPrice).
//		Scan(ctx, &v)
//
func (zqq *ZARQuoteQuery) Select(fields ...string) *ZARQuoteSelect {
	zqq.fields = append(zqq.fields, fields...)
	return &ZARQuoteSelect{ZARQuoteQuery: zqq}
}

func (zqq *ZARQuoteQuery) prepareQuery(ctx context.Context) error {
	for _, f := range zqq.fields {
		if !zarquote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if zqq.path != nil {
		prev, err := zqq.path(ctx)
		if err != nil {
			return err
		}
		zqq.sql = prev
	}
	return nil
}

func (zqq *ZARQuoteQuery) sqlAll(ctx context.Context) ([]*ZARQuote, error) {
	var (
		nodes = []*ZARQuote{}
		_spec = zqq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ZARQuote{config: zqq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, zqq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (zqq *ZARQuoteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := zqq.querySpec()
	_spec.Node.Columns = zqq.fields
	if len(zqq.fields) > 0 {
		_spec.Unique = zqq.unique != nil && *zqq.unique
	}
	return sqlgraph.CountNodes(ctx, zqq.driver, _spec)
}

func (zqq *ZARQuoteQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := zqq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (zqq *ZARQuoteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   zarquote.Table,
			Columns: zarquote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: zarquote.FieldID,
			},
		},
		From:   zqq.sql,
		Unique: true,
	}
	if unique := zqq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := zqq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, zarquote.FieldID)
		for i := range fields {
			if fields[i] != zarquote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := zqq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := zqq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := zqq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := zqq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (zqq *ZARQuoteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(zqq.driver.Dialect())
	t1 := builder.Table(zarquote.Table)
	columns := zqq.fields
	if len(columns) == 0 {
		columns = zarquote.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if zqq.sql != nil {
		selector = zqq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if zqq.unique != nil && *zqq.unique {
		selector.Distinct()
	}
	for _, p := range zqq.predicates {
		p(selector)
	}
	for _, p := range zqq.order {
		p(selector)
	}
	if offset := zqq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := zqq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ZARQuoteGroupBy is the group-by builder for ZARQuote entities.
type ZARQuoteGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (zqgb *ZARQuoteGroupBy) Aggregate(fns ...AggregateFunc) *ZARQuoteGroupBy {
	zqgb.fns = append(zqgb.fns, fns...)
	return zqgb
}

// Scan applies the group-by query and scans the result into the given value.
func (zqgb *ZARQuoteGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := zqgb.path(ctx)
	if err != nil {
		return err
	}
	zqgb.sql = query
	return zqgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (zqgb *ZARQuoteGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := zqgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (zqgb *ZARQuoteGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(zqgb.fields) > 1 {
		return nil, errors.New("ent: ZARQuoteGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := zqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (zqgb *ZARQuoteGroupBy) StringsX(ctx context.Context) []string {
	v, err := zqgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (zqgb *ZARQuoteGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = zqgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zarquote.Label}
	default:
		err = fmt.Errorf("ent: ZARQuoteGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (zqgb *ZARQuoteGroupBy) StringX(ctx context.Context) string {
	v, err := zqgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (zqgb *ZARQuoteGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(zqgb.fields) > 1 {
		return nil, errors.New("ent: ZARQuoteGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := zqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (zqgb *ZARQuoteGroupBy) IntsX(ctx context.Context) []int {
	v, err := zqgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (zqgb *ZARQuoteGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = zqgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zarquote.Label}
	default:
		err = fmt.Errorf("ent: ZARQuoteGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (zqgb *ZARQuoteGroupBy) IntX(ctx context.Context) int {
	v, err := zqgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (zqgb *ZARQuoteGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(zqgb.fields) > 1 {
		return nil, errors.New("ent: ZARQuoteGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := zqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (zqgb *ZARQuoteGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := zqgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (zqgb *ZARQuoteGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = zqgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zarquote.Label}
	default:
		err = fmt.Errorf("ent: ZARQuoteGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (zqgb *ZARQuoteGroupBy) Float64X(ctx context.Context) float64 {
	v, err := zqgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (zqgb *ZARQuoteGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(zqgb.fields) > 1 {
		return nil, errors.New("ent: ZARQuoteGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := zqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (zqgb *ZARQuoteGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := zqgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (zqgb *ZARQuoteGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = zqgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zarquote.Label}
	default:
		err = fmt.Errorf("ent: ZARQuoteGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (zqgb *ZARQuoteGroupBy) BoolX(ctx context.Context) bool {
	v, err := zqgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (zqgb *ZARQuoteGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range zqgb.fields {
		if !zarquote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := zqgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := zqgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (zqgb *ZARQuoteGroupBy) sqlQuery() *sql.Selector {
	selector := zqgb.sql.Select()
	aggregation := make([]string, 0, len(zqgb.fns))
	for _, fn := range zqgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(zqgb.fields)+len(zqgb.fns))
		for _, f := range zqgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(zqgb.fields...)...)
}

// ZARQuoteSelect is the builder for selecting fields of ZARQuote entities.
type ZARQuoteSelect struct {
	*ZARQuoteQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (zqs *ZARQuoteSelect) Scan(ctx context.Context, v interface{}) error {
	if err := zqs.prepareQuery(ctx); err != nil {
		return err
	}
	zqs.sql = zqs.ZARQuoteQuery.sqlQuery(ctx)
	return zqs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (zqs *ZARQuoteSelect) ScanX(ctx context.Context, v interface{}) {
	if err := zqs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (zqs *ZARQuoteSelect) Strings(ctx context.Context) ([]string, error) {
	if len(zqs.fields) > 1 {
		return nil, errors.New("ent: ZARQuoteSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := zqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (zqs *ZARQuoteSelect) StringsX(ctx context.Context) []string {
	v, err := zqs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (zqs *ZARQuoteSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = zqs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zarquote.Label}
	default:
		err = fmt.Errorf("ent: ZARQuoteSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (zqs *ZARQuoteSelect) StringX(ctx context.Context) string {
	v, err := zqs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (zqs *ZARQuoteSelect) Ints(ctx context.Context) ([]int, error) {
	if len(zqs.fields) > 1 {
		return nil, errors.New("ent: ZARQuoteSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := zqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (zqs *ZARQuoteSelect) IntsX(ctx context.Context) []int {
	v, err := zqs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (zqs *ZARQuoteSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = zqs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zarquote.Label}
	default:
		err = fmt.Errorf("ent: ZARQuoteSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (zqs *ZARQuoteSelect) IntX(ctx context.Context) int {
	v, err := zqs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (zqs *ZARQuoteSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(zqs.fields) > 1 {
		return nil, errors.New("ent: ZARQuoteSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := zqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (zqs *ZARQuoteSelect) Float64sX(ctx context.Context) []float64 {
	v, err := zqs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (zqs *ZARQuoteSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = zqs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zarquote.Label}
	default:
		err = fmt.Errorf("ent: ZARQuoteSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (zqs *ZARQuoteSelect) Float64X(ctx context.Context) float64 {
	v, err := zqs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (zqs *ZARQuoteSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(zqs.fields) > 1 {
		return nil, errors.New("ent: ZARQuoteSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := zqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (zqs *ZARQuoteSelect) BoolsX(ctx context.Context) []bool {
	v, err := zqs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (zqs *ZARQuoteSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = zqs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zarquote.Label}
	default:
		err = fmt.Errorf("ent: ZARQuoteSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (zqs *ZARQuoteSelect) BoolX(ctx context.Context) bool {
	v, err := zqs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (zqs *ZARQuoteSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := zqs.sql.Query()
	if err := zqs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
