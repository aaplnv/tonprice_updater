// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/cnyquote"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CNYQuoteCreate is the builder for creating a CNYQuote entity.
type CNYQuoteCreate struct {
	config
	mutation *CNYQuoteMutation
	hooks    []Hook
}

// SetPrice sets the "price" field.
func (cqc *CNYQuoteCreate) SetPrice(f float64) *CNYQuoteCreate {
	cqc.mutation.SetPrice(f)
	return cqc
}

// SetTimestamp sets the "Timestamp" field.
func (cqc *CNYQuoteCreate) SetTimestamp(t time.Time) *CNYQuoteCreate {
	cqc.mutation.SetTimestamp(t)
	return cqc
}

// SetID sets the "id" field.
func (cqc *CNYQuoteCreate) SetID(i int) *CNYQuoteCreate {
	cqc.mutation.SetID(i)
	return cqc
}

// Mutation returns the CNYQuoteMutation object of the builder.
func (cqc *CNYQuoteCreate) Mutation() *CNYQuoteMutation {
	return cqc.mutation
}

// Save creates the CNYQuote in the database.
func (cqc *CNYQuoteCreate) Save(ctx context.Context) (*CNYQuote, error) {
	var (
		err  error
		node *CNYQuote
	)
	if len(cqc.hooks) == 0 {
		if err = cqc.check(); err != nil {
			return nil, err
		}
		node, err = cqc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CNYQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cqc.check(); err != nil {
				return nil, err
			}
			cqc.mutation = mutation
			if node, err = cqc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cqc.hooks) - 1; i >= 0; i-- {
			if cqc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cqc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cqc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cqc *CNYQuoteCreate) SaveX(ctx context.Context) *CNYQuote {
	v, err := cqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cqc *CNYQuoteCreate) Exec(ctx context.Context) error {
	_, err := cqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cqc *CNYQuoteCreate) ExecX(ctx context.Context) {
	if err := cqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cqc *CNYQuoteCreate) check() error {
	if _, ok := cqc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "CNYQuote.price"`)}
	}
	if _, ok := cqc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "Timestamp", err: errors.New(`ent: missing required field "CNYQuote.Timestamp"`)}
	}
	return nil
}

func (cqc *CNYQuoteCreate) sqlSave(ctx context.Context) (*CNYQuote, error) {
	_node, _spec := cqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cqc.driver, _spec); err != nil {
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

func (cqc *CNYQuoteCreate) createSpec() (*CNYQuote, *sqlgraph.CreateSpec) {
	var (
		_node = &CNYQuote{config: cqc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cnyquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cnyquote.FieldID,
			},
		}
	)
	if id, ok := cqc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cqc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: cnyquote.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := cqc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cnyquote.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	return _node, _spec
}

// CNYQuoteCreateBulk is the builder for creating many CNYQuote entities in bulk.
type CNYQuoteCreateBulk struct {
	config
	builders []*CNYQuoteCreate
}

// Save creates the CNYQuote entities in the database.
func (cqcb *CNYQuoteCreateBulk) Save(ctx context.Context) ([]*CNYQuote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cqcb.builders))
	nodes := make([]*CNYQuote, len(cqcb.builders))
	mutators := make([]Mutator, len(cqcb.builders))
	for i := range cqcb.builders {
		func(i int, root context.Context) {
			builder := cqcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CNYQuoteMutation)
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
					_, err = mutators[i+1].Mutate(root, cqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cqcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cqcb *CNYQuoteCreateBulk) SaveX(ctx context.Context) []*CNYQuote {
	v, err := cqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cqcb *CNYQuoteCreateBulk) Exec(ctx context.Context) error {
	_, err := cqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cqcb *CNYQuoteCreateBulk) ExecX(ctx context.Context) {
	if err := cqcb.Exec(ctx); err != nil {
		panic(err)
	}
}
