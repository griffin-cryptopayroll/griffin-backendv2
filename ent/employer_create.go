// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer"
	"griffin-dao/ent/payment"
	"griffin-dao/ent/payment_history"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EMPLOYERCreate is the builder for creating a EMPLOYER entity.
type EMPLOYERCreate struct {
	config
	mutation *EMPLOYERMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (ec *EMPLOYERCreate) SetUsername(s string) *EMPLOYERCreate {
	ec.mutation.SetUsername(s)
	return ec
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (ec *EMPLOYERCreate) SetNillableUsername(s *string) *EMPLOYERCreate {
	if s != nil {
		ec.SetUsername(*s)
	}
	return ec
}

// SetPassword sets the "password" field.
func (ec *EMPLOYERCreate) SetPassword(s string) *EMPLOYERCreate {
	ec.mutation.SetPassword(s)
	return ec
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (ec *EMPLOYERCreate) SetNillablePassword(s *string) *EMPLOYERCreate {
	if s != nil {
		ec.SetPassword(*s)
	}
	return ec
}

// SetGid sets the "gid" field.
func (ec *EMPLOYERCreate) SetGid(s string) *EMPLOYERCreate {
	ec.mutation.SetGid(s)
	return ec
}

// SetCorpName sets the "corp_name" field.
func (ec *EMPLOYERCreate) SetCorpName(s string) *EMPLOYERCreate {
	ec.mutation.SetCorpName(s)
	return ec
}

// SetNillableCorpName sets the "corp_name" field if the given value is not nil.
func (ec *EMPLOYERCreate) SetNillableCorpName(s *string) *EMPLOYERCreate {
	if s != nil {
		ec.SetCorpName(*s)
	}
	return ec
}

// SetCorpEmail sets the "corp_email" field.
func (ec *EMPLOYERCreate) SetCorpEmail(s string) *EMPLOYERCreate {
	ec.mutation.SetCorpEmail(s)
	return ec
}

// SetNillableCorpEmail sets the "corp_email" field if the given value is not nil.
func (ec *EMPLOYERCreate) SetNillableCorpEmail(s *string) *EMPLOYERCreate {
	if s != nil {
		ec.SetCorpEmail(*s)
	}
	return ec
}

// SetWallet sets the "wallet" field.
func (ec *EMPLOYERCreate) SetWallet(s string) *EMPLOYERCreate {
	ec.mutation.SetWallet(s)
	return ec
}

// SetWalletAztec sets the "wallet_aztec" field.
func (ec *EMPLOYERCreate) SetWalletAztec(s string) *EMPLOYERCreate {
	ec.mutation.SetWalletAztec(s)
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EMPLOYERCreate) SetCreatedAt(t time.Time) *EMPLOYERCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetCreatedBy sets the "created_by" field.
func (ec *EMPLOYERCreate) SetCreatedBy(s string) *EMPLOYERCreate {
	ec.mutation.SetCreatedBy(s)
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EMPLOYERCreate) SetUpdatedAt(t time.Time) *EMPLOYERCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetUpdatedBy sets the "updated_by" field.
func (ec *EMPLOYERCreate) SetUpdatedBy(s string) *EMPLOYERCreate {
	ec.mutation.SetUpdatedBy(s)
	return ec
}

// SetID sets the "id" field.
func (ec *EMPLOYERCreate) SetID(i int) *EMPLOYERCreate {
	ec.mutation.SetID(i)
	return ec
}

// AddEmployerOfEmployeeIDs adds the "employer_of_employee" edge to the EMPLOYEE entity by IDs.
func (ec *EMPLOYERCreate) AddEmployerOfEmployeeIDs(ids ...int) *EMPLOYERCreate {
	ec.mutation.AddEmployerOfEmployeeIDs(ids...)
	return ec
}

// AddEmployerOfEmployee adds the "employer_of_employee" edges to the EMPLOYEE entity.
func (ec *EMPLOYERCreate) AddEmployerOfEmployee(e ...*EMPLOYEE) *EMPLOYERCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddEmployerOfEmployeeIDs(ids...)
}

// AddEmployerOfPaymentHistoryIDs adds the "employer_of_payment_history" edge to the PAYMENT_HISTORY entity by IDs.
func (ec *EMPLOYERCreate) AddEmployerOfPaymentHistoryIDs(ids ...int) *EMPLOYERCreate {
	ec.mutation.AddEmployerOfPaymentHistoryIDs(ids...)
	return ec
}

// AddEmployerOfPaymentHistory adds the "employer_of_payment_history" edges to the PAYMENT_HISTORY entity.
func (ec *EMPLOYERCreate) AddEmployerOfPaymentHistory(p ...*PAYMENT_HISTORY) *EMPLOYERCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddEmployerOfPaymentHistoryIDs(ids...)
}

// AddEmployerOfPaymentIDs adds the "employer_of_payment" edge to the PAYMENT entity by IDs.
func (ec *EMPLOYERCreate) AddEmployerOfPaymentIDs(ids ...int) *EMPLOYERCreate {
	ec.mutation.AddEmployerOfPaymentIDs(ids...)
	return ec
}

// AddEmployerOfPayment adds the "employer_of_payment" edges to the PAYMENT entity.
func (ec *EMPLOYERCreate) AddEmployerOfPayment(p ...*PAYMENT) *EMPLOYERCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddEmployerOfPaymentIDs(ids...)
}

// Mutation returns the EMPLOYERMutation object of the builder.
func (ec *EMPLOYERCreate) Mutation() *EMPLOYERMutation {
	return ec.mutation
}

// Save creates the EMPLOYER in the database.
func (ec *EMPLOYERCreate) Save(ctx context.Context) (*EMPLOYER, error) {
	var (
		err  error
		node *EMPLOYER
	)
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EMPLOYERMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EMPLOYER)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EMPLOYERMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EMPLOYERCreate) SaveX(ctx context.Context) *EMPLOYER {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EMPLOYERCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EMPLOYERCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EMPLOYERCreate) check() error {
	if _, ok := ec.mutation.Gid(); !ok {
		return &ValidationError{Name: "gid", err: errors.New(`ent: missing required field "EMPLOYER.gid"`)}
	}
	if _, ok := ec.mutation.Wallet(); !ok {
		return &ValidationError{Name: "wallet", err: errors.New(`ent: missing required field "EMPLOYER.wallet"`)}
	}
	if _, ok := ec.mutation.WalletAztec(); !ok {
		return &ValidationError{Name: "wallet_aztec", err: errors.New(`ent: missing required field "EMPLOYER.wallet_aztec"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EMPLOYER.created_at"`)}
	}
	if _, ok := ec.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "EMPLOYER.created_by"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "EMPLOYER.updated_at"`)}
	}
	if _, ok := ec.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "EMPLOYER.updated_by"`)}
	}
	return nil
}

func (ec *EMPLOYERCreate) sqlSave(ctx context.Context) (*EMPLOYER, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
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

func (ec *EMPLOYERCreate) createSpec() (*EMPLOYER, *sqlgraph.CreateSpec) {
	var (
		_node = &EMPLOYER{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: employer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employer.FieldID,
			},
		}
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := ec.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := ec.mutation.Gid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer.FieldGid,
		})
		_node.Gid = value
	}
	if value, ok := ec.mutation.CorpName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer.FieldCorpName,
		})
		_node.CorpName = value
	}
	if value, ok := ec.mutation.CorpEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer.FieldCorpEmail,
		})
		_node.CorpEmail = value
	}
	if value, ok := ec.mutation.Wallet(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer.FieldWallet,
		})
		_node.Wallet = value
	}
	if value, ok := ec.mutation.WalletAztec(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer.FieldWalletAztec,
		})
		_node.WalletAztec = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: employer.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: employer.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if nodes := ec.mutation.EmployerOfEmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employer.EmployerOfEmployeeTable,
			Columns: []string{employer.EmployerOfEmployeeColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EmployerOfPaymentHistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employer.EmployerOfPaymentHistoryTable,
			Columns: []string{employer.EmployerOfPaymentHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment_history.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EmployerOfPaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employer.EmployerOfPaymentTable,
			Columns: []string{employer.EmployerOfPaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EMPLOYERCreateBulk is the builder for creating many EMPLOYER entities in bulk.
type EMPLOYERCreateBulk struct {
	config
	builders []*EMPLOYERCreate
}

// Save creates the EMPLOYER entities in the database.
func (ecb *EMPLOYERCreateBulk) Save(ctx context.Context) ([]*EMPLOYER, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*EMPLOYER, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EMPLOYERMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EMPLOYERCreateBulk) SaveX(ctx context.Context) []*EMPLOYER {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EMPLOYERCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EMPLOYERCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
