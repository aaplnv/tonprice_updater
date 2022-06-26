// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/arsquote"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ARSQuoteCreate is the builder for creating a ARSQuote entity.
type ARSQuoteCreate struct {
	config
	mutation *ARSQuoteMutation
	hooks    []Hook
}

// SetPrice sets the "price" field.
func (aqc *ARSQuoteCreate) SetPrice(f float64) *ARSQuoteCreate {
	aqc.mutation.SetPrice(f)
	return aqc
}

// SetTimestamp sets the "Timestamp" field.
func (aqc *ARSQuoteCreate) SetTimestamp(t time.Time) *ARSQuoteCreate {
	aqc.mutation.SetTimestamp(t)
	return aqc
}

// SetID sets the "id" field.
func (aqc *ARSQuoteCreate) SetID(i int) *ARSQuoteCreate {
	aqc.mutation.SetID(i)
	return aqc
}

// Mutation returns the ARSQuoteMutation object of the builder.
func (aqc *ARSQuoteCreate) Mutation() *ARSQuoteMutation {
	return aqc.mutation
}

// Save creates the ARSQuote in the database.
func (aqc *ARSQuoteCreate) Save(ctx context.Context) (*ARSQuote, error) {
	var (
		err  error
		node *ARSQuote
	)
	if len(aqc.hooks) == 0 {
		if err = aqc.check(); err != nil {
			return nil, err
		}
		node, err = aqc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ARSQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aqc.check(); err != nil {
				return nil, err
			}
			aqc.mutation = mutation
			if node, err = aqc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(aqc.hooks) - 1; i >= 0; i-- {
			if aqc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aqc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aqc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aqc *ARSQuoteCreate) SaveX(ctx context.Context) *ARSQuote {
	v, err := aqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aqc *ARSQuoteCreate) Exec(ctx context.Context) error {
	_, err := aqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aqc *ARSQuoteCreate) ExecX(ctx context.Context) {
	if err := aqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aqc *ARSQuoteCreate) check() error {
	if _, ok := aqc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "ARSQuote.price"`)}
	}
	if _, ok := aqc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "Timestamp", err: errors.New(`ent: missing required field "ARSQuote.Timestamp"`)}
	}
	return nil
}

func (aqc *ARSQuoteCreate) sqlSave(ctx context.Context) (*ARSQuote, error) {
	_node, _spec := aqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aqc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (aqc *ARSQuoteCreate) createSpec() (*ARSQuote, *sqlgraph.CreateSpec) {
	var (
		_node = &ARSQuote{config: aqc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: arsquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: arsquote.FieldID,
			},
		}
	)
	if id, ok := aqc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := aqc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: arsquote.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := aqc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: arsquote.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	return _node, _spec
}

// ARSQuoteCreateBulk is the builder for creating many ARSQuote entities in bulk.
type ARSQuoteCreateBulk struct {
	config
	builders []*ARSQuoteCreate
}

// Save creates the ARSQuote entities in the database.
func (aqcb *ARSQuoteCreateBulk) Save(ctx context.Context) ([]*ARSQuote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aqcb.builders))
	nodes := make([]*ARSQuote, len(aqcb.builders))
	mutators := make([]Mutator, len(aqcb.builders))
	for i := range aqcb.builders {
		func(i int, root context.Context) {
			builder := aqcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ARSQuoteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aqcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, aqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aqcb *ARSQuoteCreateBulk) SaveX(ctx context.Context) []*ARSQuote {
	v, err := aqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aqcb *ARSQuoteCreateBulk) Exec(ctx context.Context) error {
	_, err := aqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aqcb *ARSQuoteCreateBulk) ExecX(ctx context.Context) {
	if err := aqcb.Exec(ctx); err != nil {
		panic(err)
	}
}
