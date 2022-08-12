// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"github.com/starryrbs/kfan/app/history/service/internal/data/ent/history"
	"github.com/starryrbs/kfan/app/history/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HistoryDelete is the builder for deleting a History entity.
type HistoryDelete struct {
	config
	hooks    []Hook
	mutation *HistoryMutation
}

// Where appends a list predicates to the HistoryDelete builder.
func (hd *HistoryDelete) Where(ps ...predicate.History) *HistoryDelete {
	hd.mutation.Where(ps...)
	return hd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hd *HistoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hd.hooks) == 0 {
		affected, err = hd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hd.mutation = mutation
			affected, err = hd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hd.hooks) - 1; i >= 0; i-- {
			if hd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hd *HistoryDelete) ExecX(ctx context.Context) int {
	n, err := hd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hd *HistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: history.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: history.FieldID,
			},
		},
	}
	if ps := hd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, hd.driver, _spec)
}

// HistoryDeleteOne is the builder for deleting a single History entity.
type HistoryDeleteOne struct {
	hd *HistoryDelete
}

// Exec executes the deletion query.
func (hdo *HistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := hdo.hd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{history.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hdo *HistoryDeleteOne) ExecX(ctx context.Context) {
	hdo.hd.ExecX(ctx)
}
