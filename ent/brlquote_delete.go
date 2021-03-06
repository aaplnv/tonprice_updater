// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"main/ent/brlquote"
	"main/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BRLQuoteDelete is the builder for deleting a BRLQuote entity.
type BRLQuoteDelete struct {
	config
	hooks    []Hook
	mutation *BRLQuoteMutation
}

// Where appends a list predicates to the BRLQuoteDelete builder.
func (bqd *BRLQuoteDelete) Where(ps ...predicate.BRLQuote) *BRLQuoteDelete {
	bqd.mutation.Where(ps...)
	return bqd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bqd *BRLQuoteDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bqd.hooks) == 0 {
		affected, err = bqd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BRLQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bqd.mutation = mutation
			affected, err = bqd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bqd.hooks) - 1; i >= 0; i-- {
			if bqd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bqd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bqd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bqd *BRLQuoteDelete) ExecX(ctx context.Context) int {
	n, err := bqd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bqd *BRLQuoteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: brlquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: brlquote.FieldID,
			},
		},
	}
	if ps := bqd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bqd.driver, _spec)
}

// BRLQuoteDeleteOne is the builder for deleting a single BRLQuote entity.
type BRLQuoteDeleteOne struct {
	bqd *BRLQuoteDelete
}

// Exec executes the deletion query.
func (bqdo *BRLQuoteDeleteOne) Exec(ctx context.Context) error {
	n, err := bqdo.bqd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{brlquote.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bqdo *BRLQuoteDeleteOne) ExecX(ctx context.Context) {
	bqdo.bqd.ExecX(ctx)
}
