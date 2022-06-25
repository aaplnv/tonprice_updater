// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/predicate"
	"main/ent/rubchart"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RUBChartUpdate is the builder for updating RUBChart entities.
type RUBChartUpdate struct {
	config
	hooks    []Hook
	mutation *RUBChartMutation
}

// Where appends a list predicates to the RUBChartUpdate builder.
func (rcu *RUBChartUpdate) Where(ps ...predicate.RUBChart) *RUBChartUpdate {
	rcu.mutation.Where(ps...)
	return rcu
}

// SetPrice sets the "price" field.
func (rcu *RUBChartUpdate) SetPrice(f float64) *RUBChartUpdate {
	rcu.mutation.ResetPrice()
	rcu.mutation.SetPrice(f)
	return rcu
}

// AddPrice adds f to the "price" field.
func (rcu *RUBChartUpdate) AddPrice(f float64) *RUBChartUpdate {
	rcu.mutation.AddPrice(f)
	return rcu
}

// Mutation returns the RUBChartMutation object of the builder.
func (rcu *RUBChartUpdate) Mutation() *RUBChartMutation {
	return rcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rcu *RUBChartUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rcu.hooks) == 0 {
		affected, err = rcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RUBChartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rcu.mutation = mutation
			affected, err = rcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rcu.hooks) - 1; i >= 0; i-- {
			if rcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rcu *RUBChartUpdate) SaveX(ctx context.Context) int {
	affected, err := rcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rcu *RUBChartUpdate) Exec(ctx context.Context) error {
	_, err := rcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcu *RUBChartUpdate) ExecX(ctx context.Context) {
	if err := rcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rcu *RUBChartUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rubchart.Table,
			Columns: rubchart.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rubchart.FieldID,
			},
		},
	}
	if ps := rcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: rubchart.FieldPrice,
		})
	}
	if value, ok := rcu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: rubchart.FieldPrice,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rubchart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RUBChartUpdateOne is the builder for updating a single RUBChart entity.
type RUBChartUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RUBChartMutation
}

// SetPrice sets the "price" field.
func (rcuo *RUBChartUpdateOne) SetPrice(f float64) *RUBChartUpdateOne {
	rcuo.mutation.ResetPrice()
	rcuo.mutation.SetPrice(f)
	return rcuo
}

// AddPrice adds f to the "price" field.
func (rcuo *RUBChartUpdateOne) AddPrice(f float64) *RUBChartUpdateOne {
	rcuo.mutation.AddPrice(f)
	return rcuo
}

// Mutation returns the RUBChartMutation object of the builder.
func (rcuo *RUBChartUpdateOne) Mutation() *RUBChartMutation {
	return rcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rcuo *RUBChartUpdateOne) Select(field string, fields ...string) *RUBChartUpdateOne {
	rcuo.fields = append([]string{field}, fields...)
	return rcuo
}

// Save executes the query and returns the updated RUBChart entity.
func (rcuo *RUBChartUpdateOne) Save(ctx context.Context) (*RUBChart, error) {
	var (
		err  error
		node *RUBChart
	)
	if len(rcuo.hooks) == 0 {
		node, err = rcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RUBChartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rcuo.mutation = mutation
			node, err = rcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rcuo.hooks) - 1; i >= 0; i-- {
			if rcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rcuo *RUBChartUpdateOne) SaveX(ctx context.Context) *RUBChart {
	node, err := rcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rcuo *RUBChartUpdateOne) Exec(ctx context.Context) error {
	_, err := rcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcuo *RUBChartUpdateOne) ExecX(ctx context.Context) {
	if err := rcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rcuo *RUBChartUpdateOne) sqlSave(ctx context.Context) (_node *RUBChart, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rubchart.Table,
			Columns: rubchart.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rubchart.FieldID,
			},
		},
	}
	id, ok := rcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RUBChart.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rubchart.FieldID)
		for _, f := range fields {
			if !rubchart.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rubchart.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: rubchart.FieldPrice,
		})
	}
	if value, ok := rcuo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: rubchart.FieldPrice,
		})
	}
	_node = &RUBChart{config: rcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rubchart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
