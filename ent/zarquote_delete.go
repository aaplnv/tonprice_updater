// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"main/ent/predicate"
	"main/ent/zarquote"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ZARQuoteDelete is the builder for deleting a ZARQuote entity.
type ZARQuoteDelete struct {
	config
	hooks    []Hook
	mutation *ZARQuoteMutation
}

// Where appends a list predicates to the ZARQuoteDelete builder.
func (zqd *ZARQuoteDelete) Where(ps ...predicate.ZARQuote) *ZARQuoteDelete {
	zqd.mutation.Where(ps...)
	return zqd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (zqd *ZARQuoteDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(zqd.hooks) == 0 {
		affected, err = zqd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ZARQuoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			zqd.mutation = mutation
			affected, err = zqd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(zqd.hooks) - 1; i >= 0; i-- {
			if zqd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = zqd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, zqd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (zqd *ZARQuoteDelete) ExecX(ctx context.Context) int {
	n, err := zqd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (zqd *ZARQuoteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: zarquote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: zarquote.FieldID,
			},
		},
	}
	if ps := zqd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, zqd.driver, _spec)
}

// ZARQuoteDeleteOne is the builder for deleting a single ZARQuote entity.
type ZARQuoteDeleteOne struct {
	zqd *ZARQuoteDelete
}

// Exec executes the deletion query.
func (zqdo *ZARQuoteDeleteOne) Exec(ctx context.Context) error {
	n, err := zqdo.zqd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{zarquote.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (zqdo *ZARQuoteDeleteOne) ExecX(ctx context.Context) {
	zqdo.zqd.ExecX(ctx)
}
