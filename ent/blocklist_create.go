// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/blocklist"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlocklistCreate is the builder for creating a Blocklist entity.
type BlocklistCreate struct {
	config
	mutation *BlocklistMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (bc *BlocklistCreate) SetType(b blocklist.Type) *BlocklistCreate {
	bc.mutation.SetType(b)
	return bc
}

// SetValue sets the "value" field.
func (bc *BlocklistCreate) SetValue(s string) *BlocklistCreate {
	bc.mutation.SetValue(s)
	return bc
}

// Mutation returns the BlocklistMutation object of the builder.
func (bc *BlocklistCreate) Mutation() *BlocklistMutation {
	return bc.mutation
}

// Save creates the Blocklist in the database.
func (bc *BlocklistCreate) Save(ctx context.Context) (*Blocklist, error) {
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlocklistCreate) SaveX(ctx context.Context) *Blocklist {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlocklistCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlocklistCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlocklistCreate) check() error {
	if _, ok := bc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Blocklist.type"`)}
	}
	if v, ok := bc.mutation.GetType(); ok {
		if err := blocklist.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Blocklist.type": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Blocklist.value"`)}
	}
	return nil
}

func (bc *BlocklistCreate) sqlSave(ctx context.Context) (*Blocklist, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BlocklistCreate) createSpec() (*Blocklist, *sqlgraph.CreateSpec) {
	var (
		_node = &Blocklist{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(blocklist.Table, sqlgraph.NewFieldSpec(blocklist.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.GetType(); ok {
		_spec.SetField(blocklist.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := bc.mutation.Value(); ok {
		_spec.SetField(blocklist.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	return _node, _spec
}

// BlocklistCreateBulk is the builder for creating many Blocklist entities in bulk.
type BlocklistCreateBulk struct {
	config
	err      error
	builders []*BlocklistCreate
}

// Save creates the Blocklist entities in the database.
func (bcb *BlocklistCreateBulk) Save(ctx context.Context) ([]*Blocklist, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blocklist, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlocklistMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlocklistCreateBulk) SaveX(ctx context.Context) []*Blocklist {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlocklistCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlocklistCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
