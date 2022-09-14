// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer_user_info"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EMPLOYERUSERINFOCreate is the builder for creating a EMPLOYER_USER_INFO entity.
type EMPLOYERUSERINFOCreate struct {
	config
	mutation *EMPLOYERUSERINFOMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (ec *EMPLOYERUSERINFOCreate) SetUsername(s string) *EMPLOYERUSERINFOCreate {
	ec.mutation.SetUsername(s)
	return ec
}

// SetPassword sets the "password" field.
func (ec *EMPLOYERUSERINFOCreate) SetPassword(s string) *EMPLOYERUSERINFOCreate {
	ec.mutation.SetPassword(s)
	return ec
}

// SetGid sets the "gid" field.
func (ec *EMPLOYERUSERINFOCreate) SetGid(s string) *EMPLOYERUSERINFOCreate {
	ec.mutation.SetGid(s)
	return ec
}

// SetCorpName sets the "corp_name" field.
func (ec *EMPLOYERUSERINFOCreate) SetCorpName(s string) *EMPLOYERUSERINFOCreate {
	ec.mutation.SetCorpName(s)
	return ec
}

// SetCorpEmail sets the "corp_email" field.
func (ec *EMPLOYERUSERINFOCreate) SetCorpEmail(s string) *EMPLOYERUSERINFOCreate {
	ec.mutation.SetCorpEmail(s)
	return ec
}

// SetWallet sets the "wallet" field.
func (ec *EMPLOYERUSERINFOCreate) SetWallet(s string) *EMPLOYERUSERINFOCreate {
	ec.mutation.SetWallet(s)
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EMPLOYERUSERINFOCreate) SetCreatedAt(t time.Time) *EMPLOYERUSERINFOCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetCreatedBy sets the "created_by" field.
func (ec *EMPLOYERUSERINFOCreate) SetCreatedBy(s string) *EMPLOYERUSERINFOCreate {
	ec.mutation.SetCreatedBy(s)
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EMPLOYERUSERINFOCreate) SetUpdatedAt(t time.Time) *EMPLOYERUSERINFOCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetUpdatedBy sets the "updated_by" field.
func (ec *EMPLOYERUSERINFOCreate) SetUpdatedBy(s string) *EMPLOYERUSERINFOCreate {
	ec.mutation.SetUpdatedBy(s)
	return ec
}

// SetID sets the "id" field.
func (ec *EMPLOYERUSERINFOCreate) SetID(i int) *EMPLOYERUSERINFOCreate {
	ec.mutation.SetID(i)
	return ec
}

// AddWorkUnderIDs adds the "work_under" edge to the EMPLOYEE entity by IDs.
func (ec *EMPLOYERUSERINFOCreate) AddWorkUnderIDs(ids ...int) *EMPLOYERUSERINFOCreate {
	ec.mutation.AddWorkUnderIDs(ids...)
	return ec
}

// AddWorkUnder adds the "work_under" edges to the EMPLOYEE entity.
func (ec *EMPLOYERUSERINFOCreate) AddWorkUnder(e ...*EMPLOYEE) *EMPLOYERUSERINFOCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddWorkUnderIDs(ids...)
}

// Mutation returns the EMPLOYERUSERINFOMutation object of the builder.
func (ec *EMPLOYERUSERINFOCreate) Mutation() *EMPLOYERUSERINFOMutation {
	return ec.mutation
}

// Save creates the EMPLOYER_USER_INFO in the database.
func (ec *EMPLOYERUSERINFOCreate) Save(ctx context.Context) (*EMPLOYER_USER_INFO, error) {
	var (
		err  error
		node *EMPLOYER_USER_INFO
	)
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EMPLOYERUSERINFOMutation)
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
		nv, ok := v.(*EMPLOYER_USER_INFO)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EMPLOYERUSERINFOMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EMPLOYERUSERINFOCreate) SaveX(ctx context.Context) *EMPLOYER_USER_INFO {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EMPLOYERUSERINFOCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EMPLOYERUSERINFOCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EMPLOYERUSERINFOCreate) check() error {
	if _, ok := ec.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "EMPLOYER_USER_INFO.username"`)}
	}
	if _, ok := ec.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "EMPLOYER_USER_INFO.password"`)}
	}
	if _, ok := ec.mutation.Gid(); !ok {
		return &ValidationError{Name: "gid", err: errors.New(`ent: missing required field "EMPLOYER_USER_INFO.gid"`)}
	}
	if _, ok := ec.mutation.CorpName(); !ok {
		return &ValidationError{Name: "corp_name", err: errors.New(`ent: missing required field "EMPLOYER_USER_INFO.corp_name"`)}
	}
	if _, ok := ec.mutation.CorpEmail(); !ok {
		return &ValidationError{Name: "corp_email", err: errors.New(`ent: missing required field "EMPLOYER_USER_INFO.corp_email"`)}
	}
	if _, ok := ec.mutation.Wallet(); !ok {
		return &ValidationError{Name: "wallet", err: errors.New(`ent: missing required field "EMPLOYER_USER_INFO.wallet"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EMPLOYER_USER_INFO.created_at"`)}
	}
	if _, ok := ec.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "EMPLOYER_USER_INFO.created_by"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "EMPLOYER_USER_INFO.updated_at"`)}
	}
	if _, ok := ec.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "EMPLOYER_USER_INFO.updated_by"`)}
	}
	return nil
}

func (ec *EMPLOYERUSERINFOCreate) sqlSave(ctx context.Context) (*EMPLOYER_USER_INFO, error) {
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

func (ec *EMPLOYERUSERINFOCreate) createSpec() (*EMPLOYER_USER_INFO, *sqlgraph.CreateSpec) {
	var (
		_node = &EMPLOYER_USER_INFO{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: employer_user_info.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employer_user_info.FieldID,
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
			Column: employer_user_info.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := ec.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer_user_info.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := ec.mutation.Gid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer_user_info.FieldGid,
		})
		_node.Gid = value
	}
	if value, ok := ec.mutation.CorpName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer_user_info.FieldCorpName,
		})
		_node.CorpName = value
	}
	if value, ok := ec.mutation.CorpEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer_user_info.FieldCorpEmail,
		})
		_node.CorpEmail = value
	}
	if value, ok := ec.mutation.Wallet(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer_user_info.FieldWallet,
		})
		_node.Wallet = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: employer_user_info.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer_user_info.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: employer_user_info.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employer_user_info.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if nodes := ec.mutation.WorkUnderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employer_user_info.WorkUnderTable,
			Columns: []string{employer_user_info.WorkUnderColumn},
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
	return _node, _spec
}

// EMPLOYERUSERINFOCreateBulk is the builder for creating many EMPLOYER_USER_INFO entities in bulk.
type EMPLOYERUSERINFOCreateBulk struct {
	config
	builders []*EMPLOYERUSERINFOCreate
}

// Save creates the EMPLOYER_USER_INFO entities in the database.
func (ecb *EMPLOYERUSERINFOCreateBulk) Save(ctx context.Context) ([]*EMPLOYER_USER_INFO, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*EMPLOYER_USER_INFO, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EMPLOYERUSERINFOMutation)
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
func (ecb *EMPLOYERUSERINFOCreateBulk) SaveX(ctx context.Context) []*EMPLOYER_USER_INFO {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EMPLOYERUSERINFOCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EMPLOYERUSERINFOCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
