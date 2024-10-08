// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"griffin-dao/ent/predicate"
	"griffin-dao/ent/tr_log"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TrLogUpdate is the builder for updating Tr_log entities.
type TrLogUpdate struct {
	config
	hooks    []Hook
	mutation *TrLogMutation
}

// Where appends a list predicates to the TrLogUpdate builder.
func (tlu *TrLogUpdate) Where(ps ...predicate.Tr_log) *TrLogUpdate {
	tlu.mutation.Where(ps...)
	return tlu
}

// SetTrType sets the "tr_type" field.
func (tlu *TrLogUpdate) SetTrType(s string) *TrLogUpdate {
	tlu.mutation.SetTrType(s)
	return tlu
}

// SetCreatedAt sets the "created_at" field.
func (tlu *TrLogUpdate) SetCreatedAt(t time.Time) *TrLogUpdate {
	tlu.mutation.SetCreatedAt(t)
	return tlu
}

// Mutation returns the TrLogMutation object of the builder.
func (tlu *TrLogUpdate) Mutation() *TrLogMutation {
	return tlu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tlu *TrLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tlu.hooks) == 0 {
		affected, err = tlu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tlu.mutation = mutation
			affected, err = tlu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tlu.hooks) - 1; i >= 0; i-- {
			if tlu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tlu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tlu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tlu *TrLogUpdate) SaveX(ctx context.Context) int {
	affected, err := tlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tlu *TrLogUpdate) Exec(ctx context.Context) error {
	_, err := tlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlu *TrLogUpdate) ExecX(ctx context.Context) {
	if err := tlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tlu *TrLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tr_log.Table,
			Columns: tr_log.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tr_log.FieldID,
			},
		},
	}
	if ps := tlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tlu.mutation.TrType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tr_log.FieldTrType,
		})
	}
	if value, ok := tlu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tr_log.FieldCreatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tr_log.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TrLogUpdateOne is the builder for updating a single Tr_log entity.
type TrLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TrLogMutation
}

// SetTrType sets the "tr_type" field.
func (tluo *TrLogUpdateOne) SetTrType(s string) *TrLogUpdateOne {
	tluo.mutation.SetTrType(s)
	return tluo
}

// SetCreatedAt sets the "created_at" field.
func (tluo *TrLogUpdateOne) SetCreatedAt(t time.Time) *TrLogUpdateOne {
	tluo.mutation.SetCreatedAt(t)
	return tluo
}

// Mutation returns the TrLogMutation object of the builder.
func (tluo *TrLogUpdateOne) Mutation() *TrLogMutation {
	return tluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tluo *TrLogUpdateOne) Select(field string, fields ...string) *TrLogUpdateOne {
	tluo.fields = append([]string{field}, fields...)
	return tluo
}

// Save executes the query and returns the updated Tr_log entity.
func (tluo *TrLogUpdateOne) Save(ctx context.Context) (*Tr_log, error) {
	var (
		err  error
		node *Tr_log
	)
	if len(tluo.hooks) == 0 {
		node, err = tluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tluo.mutation = mutation
			node, err = tluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tluo.hooks) - 1; i >= 0; i-- {
			if tluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tluo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tluo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Tr_log)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TrLogMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tluo *TrLogUpdateOne) SaveX(ctx context.Context) *Tr_log {
	node, err := tluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tluo *TrLogUpdateOne) Exec(ctx context.Context) error {
	_, err := tluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tluo *TrLogUpdateOne) ExecX(ctx context.Context) {
	if err := tluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tluo *TrLogUpdateOne) sqlSave(ctx context.Context) (_node *Tr_log, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tr_log.Table,
			Columns: tr_log.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tr_log.FieldID,
			},
		},
	}
	id, ok := tluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tr_log.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tr_log.FieldID)
		for _, f := range fields {
			if !tr_log.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tr_log.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tluo.mutation.TrType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tr_log.FieldTrType,
		})
	}
	if value, ok := tluo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tr_log.FieldCreatedAt,
		})
	}
	_node = &Tr_log{config: tluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tr_log.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
