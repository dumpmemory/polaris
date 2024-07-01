// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/epidodes"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EpidodesCreate is the builder for creating a Epidodes entity.
type EpidodesCreate struct {
	config
	mutation *EpidodesMutation
	hooks    []Hook
}

// SetSeriesID sets the "series_id" field.
func (ec *EpidodesCreate) SetSeriesID(i int) *EpidodesCreate {
	ec.mutation.SetSeriesID(i)
	return ec
}

// SetSeasonNumber sets the "season_number" field.
func (ec *EpidodesCreate) SetSeasonNumber(i int) *EpidodesCreate {
	ec.mutation.SetSeasonNumber(i)
	return ec
}

// SetEpisodeNumber sets the "episode_number" field.
func (ec *EpidodesCreate) SetEpisodeNumber(i int) *EpidodesCreate {
	ec.mutation.SetEpisodeNumber(i)
	return ec
}

// SetTitle sets the "title" field.
func (ec *EpidodesCreate) SetTitle(s string) *EpidodesCreate {
	ec.mutation.SetTitle(s)
	return ec
}

// SetOverview sets the "overview" field.
func (ec *EpidodesCreate) SetOverview(s string) *EpidodesCreate {
	ec.mutation.SetOverview(s)
	return ec
}

// SetAirDate sets the "air_date" field.
func (ec *EpidodesCreate) SetAirDate(s string) *EpidodesCreate {
	ec.mutation.SetAirDate(s)
	return ec
}

// Mutation returns the EpidodesMutation object of the builder.
func (ec *EpidodesCreate) Mutation() *EpidodesMutation {
	return ec.mutation
}

// Save creates the Epidodes in the database.
func (ec *EpidodesCreate) Save(ctx context.Context) (*Epidodes, error) {
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EpidodesCreate) SaveX(ctx context.Context) *Epidodes {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EpidodesCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EpidodesCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EpidodesCreate) check() error {
	if _, ok := ec.mutation.SeriesID(); !ok {
		return &ValidationError{Name: "series_id", err: errors.New(`ent: missing required field "Epidodes.series_id"`)}
	}
	if _, ok := ec.mutation.SeasonNumber(); !ok {
		return &ValidationError{Name: "season_number", err: errors.New(`ent: missing required field "Epidodes.season_number"`)}
	}
	if _, ok := ec.mutation.EpisodeNumber(); !ok {
		return &ValidationError{Name: "episode_number", err: errors.New(`ent: missing required field "Epidodes.episode_number"`)}
	}
	if _, ok := ec.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Epidodes.title"`)}
	}
	if _, ok := ec.mutation.Overview(); !ok {
		return &ValidationError{Name: "overview", err: errors.New(`ent: missing required field "Epidodes.overview"`)}
	}
	if _, ok := ec.mutation.AirDate(); !ok {
		return &ValidationError{Name: "air_date", err: errors.New(`ent: missing required field "Epidodes.air_date"`)}
	}
	return nil
}

func (ec *EpidodesCreate) sqlSave(ctx context.Context) (*Epidodes, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EpidodesCreate) createSpec() (*Epidodes, *sqlgraph.CreateSpec) {
	var (
		_node = &Epidodes{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(epidodes.Table, sqlgraph.NewFieldSpec(epidodes.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.SeriesID(); ok {
		_spec.SetField(epidodes.FieldSeriesID, field.TypeInt, value)
		_node.SeriesID = value
	}
	if value, ok := ec.mutation.SeasonNumber(); ok {
		_spec.SetField(epidodes.FieldSeasonNumber, field.TypeInt, value)
		_node.SeasonNumber = value
	}
	if value, ok := ec.mutation.EpisodeNumber(); ok {
		_spec.SetField(epidodes.FieldEpisodeNumber, field.TypeInt, value)
		_node.EpisodeNumber = value
	}
	if value, ok := ec.mutation.Title(); ok {
		_spec.SetField(epidodes.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ec.mutation.Overview(); ok {
		_spec.SetField(epidodes.FieldOverview, field.TypeString, value)
		_node.Overview = value
	}
	if value, ok := ec.mutation.AirDate(); ok {
		_spec.SetField(epidodes.FieldAirDate, field.TypeString, value)
		_node.AirDate = value
	}
	return _node, _spec
}

// EpidodesCreateBulk is the builder for creating many Epidodes entities in bulk.
type EpidodesCreateBulk struct {
	config
	err      error
	builders []*EpidodesCreate
}

// Save creates the Epidodes entities in the database.
func (ecb *EpidodesCreateBulk) Save(ctx context.Context) ([]*Epidodes, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Epidodes, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EpidodesMutation)
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EpidodesCreateBulk) SaveX(ctx context.Context) []*Epidodes {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EpidodesCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EpidodesCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
