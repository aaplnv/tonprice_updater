// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/jpyquote"
	"main/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JPYQuoteQuery is the builder for querying JPYQuote entities.
type JPYQuoteQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.JPYQuote
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JPYQuoteQuery builder.
func (jqq *JPYQuoteQuery) Where(ps ...predicate.JPYQuote) *JPYQuoteQuery {
	jqq.predicates = append(jqq.predicates, ps...)
	return jqq
}

// Limit adds a limit step to the query.
func (jqq *JPYQuoteQuery) Limit(limit int) *JPYQuoteQuery {
	jqq.limit = &limit
	return jqq
}

// Offset adds an offset step to the query.
func (jqq *JPYQuoteQuery) Offset(offset int) *JPYQuoteQuery {
	jqq.offset = &offset
	return jqq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jqq *JPYQuoteQuery) Unique(unique bool) *JPYQuoteQuery {
	jqq.unique = &unique
	return jqq
}

// Order adds an order step to the query.
func (jqq *JPYQuoteQuery) Order(o ...OrderFunc) *JPYQuoteQuery {
	jqq.order = append(jqq.order, o...)
	return jqq
}

// First returns the first JPYQuote entity from the query.
// Returns a *NotFoundError when no JPYQuote was found.
func (jqq *JPYQuoteQuery) First(ctx context.Context) (*JPYQuote, error) {
	nodes, err := jqq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{jpyquote.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jqq *JPYQuoteQuery) FirstX(ctx context.Context) *JPYQuote {
	node, err := jqq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first JPYQuote ID from the query.
// Returns a *NotFoundError when no JPYQuote ID was found.
func (jqq *JPYQuoteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jqq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{jpyquote.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jqq *JPYQuoteQuery) FirstIDX(ctx context.Context) int {
	id, err := jqq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single JPYQuote entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one JPYQuote entity is found.
// Returns a *NotFoundError when no JPYQuote entities are found.
func (jqq *JPYQuoteQuery) Only(ctx context.Context) (*JPYQuote, error) {
	nodes, err := jqq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{jpyquote.Label}
	default:
		return nil, &NotSingularError{jpyquote.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jqq *JPYQuoteQuery) OnlyX(ctx context.Context) *JPYQuote {
	node, err := jqq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only JPYQuote ID in the query.
// Returns a *NotSingularError when more than one JPYQuote ID is found.
// Returns a *NotFoundError when no entities are found.
func (jqq *JPYQuoteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jqq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{jpyquote.Label}
	default:
		err = &NotSingularError{jpyquote.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jqq *JPYQuoteQuery) OnlyIDX(ctx context.Context) int {
	id, err := jqq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JPYQuotes.
func (jqq *JPYQuoteQuery) All(ctx context.Context) ([]*JPYQuote, error) {
	if err := jqq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return jqq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (jqq *JPYQuoteQuery) AllX(ctx context.Context) []*JPYQuote {
	nodes, err := jqq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of JPYQuote IDs.
func (jqq *JPYQuoteQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := jqq.Select(jpyquote.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jqq *JPYQuoteQuery) IDsX(ctx context.Context) []int {
	ids, err := jqq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jqq *JPYQuoteQuery) Count(ctx context.Context) (int, error) {
	if err := jqq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return jqq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (jqq *JPYQuoteQuery) CountX(ctx context.Context) int {
	count, err := jqq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jqq *JPYQuoteQuery) Exist(ctx context.Context) (bool, error) {
	if err := jqq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return jqq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (jqq *JPYQuoteQuery) ExistX(ctx context.Context) bool {
	exist, err := jqq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JPYQuoteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jqq *JPYQuoteQuery) Clone() *JPYQuoteQuery {
	if jqq == nil {
		return nil
	}
	return &JPYQuoteQuery{
		config:     jqq.config,
		limit:      jqq.limit,
		offset:     jqq.offset,
		order:      append([]OrderFunc{}, jqq.order...),
		predicates: append([]predicate.JPYQuote{}, jqq.predicates...),
		// clone intermediate query.
		sql:    jqq.sql.Clone(),
		path:   jqq.path,
		unique: jqq.unique,
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
//	client.JPYQuote.Query().
//		GroupBy(jpyquote.FieldPrice).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (jqq *JPYQuoteQuery) GroupBy(field string, fields ...string) *JPYQuoteGroupBy {
	group := &JPYQuoteGroupBy{config: jqq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := jqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return jqq.sqlQuery(ctx), nil
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
//	client.JPYQuote.Query().
//		Select(jpyquote.FieldPrice).
//		Scan(ctx, &v)
//
func (jqq *JPYQuoteQuery) Select(fields ...string) *JPYQuoteSelect {
	jqq.fields = append(jqq.fields, fields...)
	return &JPYQuoteSelect{JPYQuoteQuery: jqq}
}

func (jqq *JPYQuoteQuery) prepareQuery(ctx context.Context) error {
	for _, f := range jqq.fields {
		if !jpyquote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if jqq.path != nil {
		prev, err := jqq.path(ctx)
		if err != nil {
			return err
		}
		jqq.sql = prev
	}
	return nil
}

func (jqq *JPYQuoteQuery) sqlAll(ctx context.Context) ([]*JPYQuote, error) {
	var (
		nodes = []*JPYQuote{}
		_spec = jqq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &JPYQuote{config: jqq.config}
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
	if err := sqlgraph.QueryNodes(ctx, jqq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (jqq *JPYQuoteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jqq.querySpec()
	_spec.Node.Columns = jqq.fields
	if len(jqq.fields) > 0 {
		_spec.Unique = jqq.unique != nil && *jqq.unique
	}
	return sqlgraph.CountNodes(ctx, jqq.driver, _spec)
}

func (jqq *JPYQuoteQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := jqq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (jqq *JPYQuoteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   jpyquote.Table,
			Columns: jpyquote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: jpyquote.FieldID,
			},
		},
		From:   jqq.sql,
		Unique: true,
	}
	if unique := jqq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := jqq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jpyquote.FieldID)
		for i := range fields {
			if fields[i] != jpyquote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := jqq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jqq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jqq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jqq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jqq *JPYQuoteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jqq.driver.Dialect())
	t1 := builder.Table(jpyquote.Table)
	columns := jqq.fields
	if len(columns) == 0 {
		columns = jpyquote.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if jqq.sql != nil {
		selector = jqq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if jqq.unique != nil && *jqq.unique {
		selector.Distinct()
	}
	for _, p := range jqq.predicates {
		p(selector)
	}
	for _, p := range jqq.order {
		p(selector)
	}
	if offset := jqq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jqq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// JPYQuoteGroupBy is the group-by builder for JPYQuote entities.
type JPYQuoteGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jqgb *JPYQuoteGroupBy) Aggregate(fns ...AggregateFunc) *JPYQuoteGroupBy {
	jqgb.fns = append(jqgb.fns, fns...)
	return jqgb
}

// Scan applies the group-by query and scans the result into the given value.
func (jqgb *JPYQuoteGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := jqgb.path(ctx)
	if err != nil {
		return err
	}
	jqgb.sql = query
	return jqgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (jqgb *JPYQuoteGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := jqgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (jqgb *JPYQuoteGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(jqgb.fields) > 1 {
		return nil, errors.New("ent: JPYQuoteGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := jqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (jqgb *JPYQuoteGroupBy) StringsX(ctx context.Context) []string {
	v, err := jqgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jqgb *JPYQuoteGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = jqgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jpyquote.Label}
	default:
		err = fmt.Errorf("ent: JPYQuoteGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (jqgb *JPYQuoteGroupBy) StringX(ctx context.Context) string {
	v, err := jqgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (jqgb *JPYQuoteGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(jqgb.fields) > 1 {
		return nil, errors.New("ent: JPYQuoteGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := jqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (jqgb *JPYQuoteGroupBy) IntsX(ctx context.Context) []int {
	v, err := jqgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jqgb *JPYQuoteGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = jqgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jpyquote.Label}
	default:
		err = fmt.Errorf("ent: JPYQuoteGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (jqgb *JPYQuoteGroupBy) IntX(ctx context.Context) int {
	v, err := jqgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (jqgb *JPYQuoteGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(jqgb.fields) > 1 {
		return nil, errors.New("ent: JPYQuoteGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := jqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (jqgb *JPYQuoteGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := jqgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jqgb *JPYQuoteGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = jqgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jpyquote.Label}
	default:
		err = fmt.Errorf("ent: JPYQuoteGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (jqgb *JPYQuoteGroupBy) Float64X(ctx context.Context) float64 {
	v, err := jqgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (jqgb *JPYQuoteGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(jqgb.fields) > 1 {
		return nil, errors.New("ent: JPYQuoteGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := jqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (jqgb *JPYQuoteGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := jqgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jqgb *JPYQuoteGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = jqgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jpyquote.Label}
	default:
		err = fmt.Errorf("ent: JPYQuoteGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (jqgb *JPYQuoteGroupBy) BoolX(ctx context.Context) bool {
	v, err := jqgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jqgb *JPYQuoteGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range jqgb.fields {
		if !jpyquote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := jqgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jqgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (jqgb *JPYQuoteGroupBy) sqlQuery() *sql.Selector {
	selector := jqgb.sql.Select()
	aggregation := make([]string, 0, len(jqgb.fns))
	for _, fn := range jqgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(jqgb.fields)+len(jqgb.fns))
		for _, f := range jqgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(jqgb.fields...)...)
}

// JPYQuoteSelect is the builder for selecting fields of JPYQuote entities.
type JPYQuoteSelect struct {
	*JPYQuoteQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (jqs *JPYQuoteSelect) Scan(ctx context.Context, v interface{}) error {
	if err := jqs.prepareQuery(ctx); err != nil {
		return err
	}
	jqs.sql = jqs.JPYQuoteQuery.sqlQuery(ctx)
	return jqs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (jqs *JPYQuoteSelect) ScanX(ctx context.Context, v interface{}) {
	if err := jqs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (jqs *JPYQuoteSelect) Strings(ctx context.Context) ([]string, error) {
	if len(jqs.fields) > 1 {
		return nil, errors.New("ent: JPYQuoteSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := jqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (jqs *JPYQuoteSelect) StringsX(ctx context.Context) []string {
	v, err := jqs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (jqs *JPYQuoteSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = jqs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jpyquote.Label}
	default:
		err = fmt.Errorf("ent: JPYQuoteSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (jqs *JPYQuoteSelect) StringX(ctx context.Context) string {
	v, err := jqs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (jqs *JPYQuoteSelect) Ints(ctx context.Context) ([]int, error) {
	if len(jqs.fields) > 1 {
		return nil, errors.New("ent: JPYQuoteSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := jqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (jqs *JPYQuoteSelect) IntsX(ctx context.Context) []int {
	v, err := jqs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (jqs *JPYQuoteSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = jqs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jpyquote.Label}
	default:
		err = fmt.Errorf("ent: JPYQuoteSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (jqs *JPYQuoteSelect) IntX(ctx context.Context) int {
	v, err := jqs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (jqs *JPYQuoteSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(jqs.fields) > 1 {
		return nil, errors.New("ent: JPYQuoteSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := jqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (jqs *JPYQuoteSelect) Float64sX(ctx context.Context) []float64 {
	v, err := jqs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (jqs *JPYQuoteSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = jqs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jpyquote.Label}
	default:
		err = fmt.Errorf("ent: JPYQuoteSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (jqs *JPYQuoteSelect) Float64X(ctx context.Context) float64 {
	v, err := jqs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (jqs *JPYQuoteSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(jqs.fields) > 1 {
		return nil, errors.New("ent: JPYQuoteSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := jqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (jqs *JPYQuoteSelect) BoolsX(ctx context.Context) []bool {
	v, err := jqs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (jqs *JPYQuoteSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = jqs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jpyquote.Label}
	default:
		err = fmt.Errorf("ent: JPYQuoteSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (jqs *JPYQuoteSelect) BoolX(ctx context.Context) bool {
	v, err := jqs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jqs *JPYQuoteSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := jqs.sql.Query()
	if err := jqs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
