// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/predicate"
	"main/ent/tryquote"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TRYQuoteUpdate is the builder for updating TRYQuote entities.
type TRYQuoteUpdate struct {
	config
	hooks    []Hook
	mutation *TRYQuoteMutation
}

// Where appends a list predicates to the TRYQuoteUpdate builder.
func (tqu *TRYQuoteUpdate) Where(ps ...predicate.TRYQuote) *TRYQuoteUpdate {
	tqu.mutation.Where(ps...)
	return tqu
}

// SetPrice sets the "price" field.
func (tqu *TRYQuoteUpdate) SetPrice(f float64) *TRYQuoteUpdate {
	tqu.mutation.ResetPrice()
	tqu.mutation.SetPrice(f)
	return tqu
}

// AddPrice adds f to the "price" field.
func (tqu *TRYQuoteUpdate) AddPrice(f float64) *TRYQuoteUpdate {
	tqu.mutation.AddPrice(f)
	return tqu
}

// SetTimestamp sets the "Timestamp" field.
func (tqu *TRYQuoteUpdate) SetTimestamp(t time.Time) *TRYQuoteUpdate {
	tqu.mutation.SetTimestamp(t)
	return tqu
}

// Mutation returns the TRYQuoteMutation object of the builder.
func (tqu *TRYQuoteUpdate) Mutation() *TRYQuoteMutation {
	return tqu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tqu *TRYQuoteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tqu.hooks) == 0 {
		affected, err = tqu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TRYQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tqu.mutation = mutation
			affected, err = tqu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tqu.hooks) - 1; i >= 0; i-- {
			if tqu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tqu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tqu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tqu *TRYQuoteUpdate) SaveX(ctx context.Context) int {
	affected, err := tqu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tqu *TRYQuoteUpdate) Exec(ctx context.Context) error {
	_, err := tqu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqu *TRYQuoteUpdate) ExecX(ctx context.Context) {
	if err := tqu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tqu *TRYQuoteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tryquote.Table,
			Columns: tryquote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tryquote.FieldID,
			},
		},
	}
	if ps := tqu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tqu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: tryquote.FieldPrice,
		})
	}
	if value, ok := tqu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: tryquote.FieldPrice,
		})
	}
	if value, ok := tqu.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tryquote.FieldTimestamp,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tqu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tryquote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TRYQuoteUpdateOne is the builder for updating a single TRYQuote entity.
type TRYQuoteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TRYQuoteMutation
}

// SetPrice sets the "price" field.
func (tquo *TRYQuoteUpdateOne) SetPrice(f float64) *TRYQuoteUpdateOne {
	tquo.mutation.ResetPrice()
	tquo.mutation.SetPrice(f)
	return tquo
}

// AddPrice adds f to the "price" field.
func (tquo *TRYQuoteUpdateOne) AddPrice(f float64) *TRYQuoteUpdateOne {
	tquo.mutation.AddPrice(f)
	return tquo
}

// SetTimestamp sets the "Timestamp" field.
func (tquo *TRYQuoteUpdateOne) SetTimestamp(t time.Time) *TRYQuoteUpdateOne {
	tquo.mutation.SetTimestamp(t)
	return tquo
}

// Mutation returns the TRYQuoteMutation object of the builder.
func (tquo *TRYQuoteUpdateOne) Mutation() *TRYQuoteMutation {
	return tquo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tquo *TRYQuoteUpdateOne) Select(field string, fields ...string) *TRYQuoteUpdateOne {
	tquo.fields = append([]string{field}, fields...)
	return tquo
}

// Save executes the query and returns the updated TRYQuote entity.
func (tquo *TRYQuoteUpdateOne) Save(ctx context.Context) (*TRYQuote, error) {
	var (
		err  error
		node *TRYQuote
	)
	if len(tquo.hooks) == 0 {
		node, err = tquo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TRYQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tquo.mutation = mutation
			node, err = tquo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tquo.hooks) - 1; i >= 0; i-- {
			if tquo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tquo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tquo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tquo *TRYQuoteUpdateOne) SaveX(ctx context.Context) *TRYQuote {
	node, err := tquo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tquo *TRYQuoteUpdateOne) Exec(ctx context.Context) error {
	_, err := tquo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tquo *TRYQuoteUpdateOne) ExecX(ctx context.Context) {
	if err := tquo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tquo *TRYQuoteUpdateOne) sqlSave(ctx context.Context) (_node *TRYQuote, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tryquote.Table,
			Columns: tryquote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tryquote.FieldID,
			},
		},
	}
	id, ok := tquo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TRYQuote.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tquo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tryquote.FieldID)
		for _, f := range fields {
			if !tryquote.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tryquote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tquo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tquo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: tryquote.FieldPrice,
		})
	}
	if value, ok := tquo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: tryquote.FieldPrice,
		})
	}
	if value, ok := tquo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tryquote.FieldTimestamp,
		})
	}
	_node = &TRYQuote{config: tquo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tquo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tryquote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
