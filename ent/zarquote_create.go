// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/zarquote"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ZARQuoteCreate is the builder for creating a ZARQuote entity.
type ZARQuoteCreate struct {
	config
	mutation *ZARQuoteMutation
	hooks    []Hook
}

// SetPrice sets the "price" field.
func (zqc *ZARQuoteCreate) SetPrice(f float64) *ZARQuoteCreate {
	zqc.mutation.SetPrice(f)
	return zqc
}

// SetTimestamp sets the "Timestamp" field.
func (zqc *ZARQuoteCreate) SetTimestamp(t time.Time) *ZARQuoteCreate {
	zqc.mutation.SetTimestamp(t)
	return zqc
}

// SetID sets the "id" field.
func (zqc *ZARQuoteCreate) SetID(i int) *ZARQuoteCreate {
	zqc.mutation.SetID(i)
	return zqc
}

// Mutation returns the ZARQuoteMutation object of the builder.
func (zqc *ZARQuoteCreate) Mutation() *ZARQuoteMutation {
	return zqc.mutation
}

// Save creates the ZARQuote in the database.
func (zqc *ZARQuoteCreate) Save(ctx context.Context) (*ZARQuote, error) {
	var (
		err  error
		node *ZARQuote
	)
	if len(zqc.hooks) == 0 {
		if err = zqc.check(); err != nil {
			return nil, err
		}
		node, err = zqc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ZARQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = zqc.check(); err != nil {
				return nil, err
			}
			zqc.mutation = mutation
			if node, err = zqc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(zqc.hooks) - 1; i >= 0; i-- {
			if zqc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = zqc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, zqc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (zqc *ZARQuoteCreate) SaveX(ctx context.Context) *ZARQuote {
	v, err := zqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (zqc *ZARQuoteCreate) Exec(ctx context.Context) error {
	_, err := zqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (zqc *ZARQuoteCreate) ExecX(ctx context.Context) {
	if err := zqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (zqc *ZARQuoteCreate) check() error {
	if _, ok := zqc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "ZARQuote.price"`)}
	}
	if _, ok := zqc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "Timestamp", err: errors.New(`ent: missing required field "ZARQuote.Timestamp"`)}
	}
	return nil
}

func (zqc *ZARQuoteCreate) sqlSave(ctx context.Context) (*ZARQuote, error) {
	_node, _spec := zqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, zqc.driver, _spec); err != nil {
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

func (zqc *ZARQuoteCreate) createSpec() (*ZARQuote, *sqlgraph.CreateSpec) {
	var (
		_node = &ZARQuote{config: zqc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: zarquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: zarquote.FieldID,
			},
		}
	)
	if id, ok := zqc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := zqc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: zarquote.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := zqc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: zarquote.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	return _node, _spec
}

// ZARQuoteCreateBulk is the builder for creating many ZARQuote entities in bulk.
type ZARQuoteCreateBulk struct {
	config
	builders []*ZARQuoteCreate
}

// Save creates the ZARQuote entities in the database.
func (zqcb *ZARQuoteCreateBulk) Save(ctx context.Context) ([]*ZARQuote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(zqcb.builders))
	nodes := make([]*ZARQuote, len(zqcb.builders))
	mutators := make([]Mutator, len(zqcb.builders))
	for i := range zqcb.builders {
		func(i int, root context.Context) {
			builder := zqcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ZARQuoteMutation)
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
					_, err = mutators[i+1].Mutate(root, zqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, zqcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, zqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (zqcb *ZARQuoteCreateBulk) SaveX(ctx context.Context) []*ZARQuote {
	v, err := zqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (zqcb *ZARQuoteCreateBulk) Exec(ctx context.Context) error {
	_, err := zqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (zqcb *ZARQuoteCreateBulk) ExecX(ctx context.Context) {
	if err := zqcb.Exec(ctx); err != nil {
		panic(err)
	}
}
