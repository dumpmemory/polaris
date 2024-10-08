// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/importlist"
	"polaris/ent/schema"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImportListCreate is the builder for creating a ImportList entity.
type ImportListCreate struct {
	config
	mutation *ImportListMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ilc *ImportListCreate) SetName(s string) *ImportListCreate {
	ilc.mutation.SetName(s)
	return ilc
}

// SetType sets the "type" field.
func (ilc *ImportListCreate) SetType(i importlist.Type) *ImportListCreate {
	ilc.mutation.SetType(i)
	return ilc
}

// SetURL sets the "url" field.
func (ilc *ImportListCreate) SetURL(s string) *ImportListCreate {
	ilc.mutation.SetURL(s)
	return ilc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (ilc *ImportListCreate) SetNillableURL(s *string) *ImportListCreate {
	if s != nil {
		ilc.SetURL(*s)
	}
	return ilc
}

// SetQulity sets the "qulity" field.
func (ilc *ImportListCreate) SetQulity(s string) *ImportListCreate {
	ilc.mutation.SetQulity(s)
	return ilc
}

// SetStorageID sets the "storage_id" field.
func (ilc *ImportListCreate) SetStorageID(i int) *ImportListCreate {
	ilc.mutation.SetStorageID(i)
	return ilc
}

// SetSettings sets the "settings" field.
func (ilc *ImportListCreate) SetSettings(sls schema.ImportListSettings) *ImportListCreate {
	ilc.mutation.SetSettings(sls)
	return ilc
}

// SetNillableSettings sets the "settings" field if the given value is not nil.
func (ilc *ImportListCreate) SetNillableSettings(sls *schema.ImportListSettings) *ImportListCreate {
	if sls != nil {
		ilc.SetSettings(*sls)
	}
	return ilc
}

// Mutation returns the ImportListMutation object of the builder.
func (ilc *ImportListCreate) Mutation() *ImportListMutation {
	return ilc.mutation
}

// Save creates the ImportList in the database.
func (ilc *ImportListCreate) Save(ctx context.Context) (*ImportList, error) {
	return withHooks(ctx, ilc.sqlSave, ilc.mutation, ilc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ilc *ImportListCreate) SaveX(ctx context.Context) *ImportList {
	v, err := ilc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ilc *ImportListCreate) Exec(ctx context.Context) error {
	_, err := ilc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ilc *ImportListCreate) ExecX(ctx context.Context) {
	if err := ilc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ilc *ImportListCreate) check() error {
	if _, ok := ilc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ImportList.name"`)}
	}
	if _, ok := ilc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "ImportList.type"`)}
	}
	if v, ok := ilc.mutation.GetType(); ok {
		if err := importlist.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ImportList.type": %w`, err)}
		}
	}
	if _, ok := ilc.mutation.Qulity(); !ok {
		return &ValidationError{Name: "qulity", err: errors.New(`ent: missing required field "ImportList.qulity"`)}
	}
	if _, ok := ilc.mutation.StorageID(); !ok {
		return &ValidationError{Name: "storage_id", err: errors.New(`ent: missing required field "ImportList.storage_id"`)}
	}
	return nil
}

func (ilc *ImportListCreate) sqlSave(ctx context.Context) (*ImportList, error) {
	if err := ilc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ilc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ilc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ilc.mutation.id = &_node.ID
	ilc.mutation.done = true
	return _node, nil
}

func (ilc *ImportListCreate) createSpec() (*ImportList, *sqlgraph.CreateSpec) {
	var (
		_node = &ImportList{config: ilc.config}
		_spec = sqlgraph.NewCreateSpec(importlist.Table, sqlgraph.NewFieldSpec(importlist.FieldID, field.TypeInt))
	)
	if value, ok := ilc.mutation.Name(); ok {
		_spec.SetField(importlist.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ilc.mutation.GetType(); ok {
		_spec.SetField(importlist.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ilc.mutation.URL(); ok {
		_spec.SetField(importlist.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ilc.mutation.Qulity(); ok {
		_spec.SetField(importlist.FieldQulity, field.TypeString, value)
		_node.Qulity = value
	}
	if value, ok := ilc.mutation.StorageID(); ok {
		_spec.SetField(importlist.FieldStorageID, field.TypeInt, value)
		_node.StorageID = value
	}
	if value, ok := ilc.mutation.Settings(); ok {
		_spec.SetField(importlist.FieldSettings, field.TypeJSON, value)
		_node.Settings = value
	}
	return _node, _spec
}

// ImportListCreateBulk is the builder for creating many ImportList entities in bulk.
type ImportListCreateBulk struct {
	config
	err      error
	builders []*ImportListCreate
}

// Save creates the ImportList entities in the database.
func (ilcb *ImportListCreateBulk) Save(ctx context.Context) ([]*ImportList, error) {
	if ilcb.err != nil {
		return nil, ilcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ilcb.builders))
	nodes := make([]*ImportList, len(ilcb.builders))
	mutators := make([]Mutator, len(ilcb.builders))
	for i := range ilcb.builders {
		func(i int, root context.Context) {
			builder := ilcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImportListMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ilcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ilcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, ilcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ilcb *ImportListCreateBulk) SaveX(ctx context.Context) []*ImportList {
	v, err := ilcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ilcb *ImportListCreateBulk) Exec(ctx context.Context) error {
	_, err := ilcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ilcb *ImportListCreateBulk) ExecX(ctx context.Context) {
	if err := ilcb.Exec(ctx); err != nil {
		panic(err)
	}
}
