// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"griffin-dao/ent/crypto_currency"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer"
	"griffin-dao/ent/payment"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PAYMENTCreate is the builder for creating a PAYMENT entity.
type PAYMENTCreate struct {
	config
	mutation *PAYMENTMutation
	hooks    []Hook
}

// SetEmployeeID sets the "employee_id" field.
func (pc *PAYMENTCreate) SetEmployeeID(i int) *PAYMENTCreate {
	pc.mutation.SetEmployeeID(i)
	return pc
}

// SetNillableEmployeeID sets the "employee_id" field if the given value is not nil.
func (pc *PAYMENTCreate) SetNillableEmployeeID(i *int) *PAYMENTCreate {
	if i != nil {
		pc.SetEmployeeID(*i)
	}
	return pc
}

// SetEmployerID sets the "employer_id" field.
func (pc *PAYMENTCreate) SetEmployerID(i int) *PAYMENTCreate {
	pc.mutation.SetEmployerID(i)
	return pc
}

// SetNillableEmployerID sets the "employer_id" field if the given value is not nil.
func (pc *PAYMENTCreate) SetNillableEmployerID(i *int) *PAYMENTCreate {
	if i != nil {
		pc.SetEmployerID(*i)
	}
	return pc
}

// SetPaymentScheduled sets the "payment_scheduled" field.
func (pc *PAYMENTCreate) SetPaymentScheduled(t time.Time) *PAYMENTCreate {
	pc.mutation.SetPaymentScheduled(t)
	return pc
}

// SetNillablePaymentScheduled sets the "payment_scheduled" field if the given value is not nil.
func (pc *PAYMENTCreate) SetNillablePaymentScheduled(t *time.Time) *PAYMENTCreate {
	if t != nil {
		pc.SetPaymentScheduled(*t)
	}
	return pc
}

// SetPaymentExecuted sets the "payment_executed" field.
func (pc *PAYMENTCreate) SetPaymentExecuted(t time.Time) *PAYMENTCreate {
	pc.mutation.SetPaymentExecuted(t)
	return pc
}

// SetNillablePaymentExecuted sets the "payment_executed" field if the given value is not nil.
func (pc *PAYMENTCreate) SetNillablePaymentExecuted(t *time.Time) *PAYMENTCreate {
	if t != nil {
		pc.SetPaymentExecuted(*t)
	}
	return pc
}

// SetPaymentAmount sets the "payment_amount" field.
func (pc *PAYMENTCreate) SetPaymentAmount(f float64) *PAYMENTCreate {
	pc.mutation.SetPaymentAmount(f)
	return pc
}

// SetCryptoCurrencyID sets the "crypto_currency_id" field.
func (pc *PAYMENTCreate) SetCryptoCurrencyID(i int) *PAYMENTCreate {
	pc.mutation.SetCryptoCurrencyID(i)
	return pc
}

// SetNillableCryptoCurrencyID sets the "crypto_currency_id" field if the given value is not nil.
func (pc *PAYMENTCreate) SetNillableCryptoCurrencyID(i *int) *PAYMENTCreate {
	if i != nil {
		pc.SetCryptoCurrencyID(*i)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PAYMENTCreate) SetID(i int) *PAYMENTCreate {
	pc.mutation.SetID(i)
	return pc
}

// SetPaymentFromEmployerID sets the "payment_from_employer" edge to the EMPLOYER entity by ID.
func (pc *PAYMENTCreate) SetPaymentFromEmployerID(id int) *PAYMENTCreate {
	pc.mutation.SetPaymentFromEmployerID(id)
	return pc
}

// SetNillablePaymentFromEmployerID sets the "payment_from_employer" edge to the EMPLOYER entity by ID if the given value is not nil.
func (pc *PAYMENTCreate) SetNillablePaymentFromEmployerID(id *int) *PAYMENTCreate {
	if id != nil {
		pc = pc.SetPaymentFromEmployerID(*id)
	}
	return pc
}

// SetPaymentFromEmployer sets the "payment_from_employer" edge to the EMPLOYER entity.
func (pc *PAYMENTCreate) SetPaymentFromEmployer(e *EMPLOYER) *PAYMENTCreate {
	return pc.SetPaymentFromEmployerID(e.ID)
}

// SetPaymentFromEmployeeID sets the "payment_from_employee" edge to the EMPLOYEE entity by ID.
func (pc *PAYMENTCreate) SetPaymentFromEmployeeID(id int) *PAYMENTCreate {
	pc.mutation.SetPaymentFromEmployeeID(id)
	return pc
}

// SetNillablePaymentFromEmployeeID sets the "payment_from_employee" edge to the EMPLOYEE entity by ID if the given value is not nil.
func (pc *PAYMENTCreate) SetNillablePaymentFromEmployeeID(id *int) *PAYMENTCreate {
	if id != nil {
		pc = pc.SetPaymentFromEmployeeID(*id)
	}
	return pc
}

// SetPaymentFromEmployee sets the "payment_from_employee" edge to the EMPLOYEE entity.
func (pc *PAYMENTCreate) SetPaymentFromEmployee(e *EMPLOYEE) *PAYMENTCreate {
	return pc.SetPaymentFromEmployeeID(e.ID)
}

// SetPaymentFromCurrencyID sets the "payment_from_currency" edge to the CRYPTO_CURRENCY entity by ID.
func (pc *PAYMENTCreate) SetPaymentFromCurrencyID(id int) *PAYMENTCreate {
	pc.mutation.SetPaymentFromCurrencyID(id)
	return pc
}

// SetNillablePaymentFromCurrencyID sets the "payment_from_currency" edge to the CRYPTO_CURRENCY entity by ID if the given value is not nil.
func (pc *PAYMENTCreate) SetNillablePaymentFromCurrencyID(id *int) *PAYMENTCreate {
	if id != nil {
		pc = pc.SetPaymentFromCurrencyID(*id)
	}
	return pc
}

// SetPaymentFromCurrency sets the "payment_from_currency" edge to the CRYPTO_CURRENCY entity.
func (pc *PAYMENTCreate) SetPaymentFromCurrency(c *CRYPTO_CURRENCY) *PAYMENTCreate {
	return pc.SetPaymentFromCurrencyID(c.ID)
}

// Mutation returns the PAYMENTMutation object of the builder.
func (pc *PAYMENTCreate) Mutation() *PAYMENTMutation {
	return pc.mutation
}

// Save creates the PAYMENT in the database.
func (pc *PAYMENTCreate) Save(ctx context.Context) (*PAYMENT, error) {
	var (
		err  error
		node *PAYMENT
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PAYMENTMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PAYMENT)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PAYMENTMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PAYMENTCreate) SaveX(ctx context.Context) *PAYMENT {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PAYMENTCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PAYMENTCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PAYMENTCreate) check() error {
	if _, ok := pc.mutation.PaymentAmount(); !ok {
		return &ValidationError{Name: "payment_amount", err: errors.New(`ent: missing required field "PAYMENT.payment_amount"`)}
	}
	return nil
}

func (pc *PAYMENTCreate) sqlSave(ctx context.Context) (*PAYMENT, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (pc *PAYMENTCreate) createSpec() (*PAYMENT, *sqlgraph.CreateSpec) {
	var (
		_node = &PAYMENT{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: payment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payment.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.PaymentScheduled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payment.FieldPaymentScheduled,
		})
		_node.PaymentScheduled = value
	}
	if value, ok := pc.mutation.PaymentExecuted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payment.FieldPaymentExecuted,
		})
		_node.PaymentExecuted = value
	}
	if value, ok := pc.mutation.PaymentAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: payment.FieldPaymentAmount,
		})
		_node.PaymentAmount = value
	}
	if nodes := pc.mutation.PaymentFromEmployerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.PaymentFromEmployerTable,
			Columns: []string{payment.PaymentFromEmployerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EmployerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PaymentFromEmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.PaymentFromEmployeeTable,
			Columns: []string{payment.PaymentFromEmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EmployeeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PaymentFromCurrencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.PaymentFromCurrencyTable,
			Columns: []string{payment.PaymentFromCurrencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: crypto_currency.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CryptoCurrencyID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PAYMENTCreateBulk is the builder for creating many PAYMENT entities in bulk.
type PAYMENTCreateBulk struct {
	config
	builders []*PAYMENTCreate
}

// Save creates the PAYMENT entities in the database.
func (pcb *PAYMENTCreateBulk) Save(ctx context.Context) ([]*PAYMENT, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*PAYMENT, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PAYMENTMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PAYMENTCreateBulk) SaveX(ctx context.Context) []*PAYMENT {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PAYMENTCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PAYMENTCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
