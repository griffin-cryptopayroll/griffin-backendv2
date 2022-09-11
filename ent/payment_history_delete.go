// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"griffin-dao/ent/payment_history"
	"griffin-dao/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PAYMENTHISTORYDelete is the builder for deleting a PAYMENT_HISTORY entity.
type PAYMENTHISTORYDelete struct {
	config
	hooks    []Hook
	mutation *PAYMENTHISTORYMutation
}

// Where appends a list predicates to the PAYMENTHISTORYDelete builder.
func (pd *PAYMENTHISTORYDelete) Where(ps ...predicate.PAYMENT_HISTORY) *PAYMENTHISTORYDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PAYMENTHISTORYDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pd.hooks) == 0 {
		affected, err = pd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PAYMENTHISTORYMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pd.mutation = mutation
			affected, err = pd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pd.hooks) - 1; i >= 0; i-- {
			if pd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PAYMENTHISTORYDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PAYMENTHISTORYDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: payment_history.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payment_history.FieldID,
			},
		},
	}
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// PAYMENTHISTORYDeleteOne is the builder for deleting a single PAYMENT_HISTORY entity.
type PAYMENTHISTORYDeleteOne struct {
	pd *PAYMENTHISTORYDelete
}

// Exec executes the deletion query.
func (pdo *PAYMENTHISTORYDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{payment_history.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PAYMENTHISTORYDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}
