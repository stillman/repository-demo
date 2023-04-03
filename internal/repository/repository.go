package repository

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type sqler interface {
	Table() string
	PK() string
	PKValue() any
	SelectFields() []string
}

type Repository[T sqler] struct {
	entity T
	db     *sqlx.DB
	offset uint64
	limit  uint64
}

func New[T sqler](db *sqlx.DB) *Repository[T] {
	return &Repository[T]{db: db}
}

func (r *Repository[T]) Limit(limit uint64) *Repository[T] {
	r.limit = limit
	return r
}

func (r *Repository[T]) Offset(offset uint64) *Repository[T] {
	r.offset = offset
	return r
}

func (r *Repository[T]) Find(sqlizer sq.Sqlizer) (*T, error) {
	query, args, _ := sq.StatementBuilder.
		Select(r.entity.SelectFields()...).
		From(r.entity.Table()).
		Where(sqlizer).
		Limit(1).
		ToSql()

	var result T

	if err := r.db.Get(&result, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("wat: %w", err)
	}

	return &result, nil
}

func (r *Repository[T]) FindAll(spec Spec) ([]*T, error) {
	q := sq.StatementBuilder.
		Select(r.entity.SelectFields()...).
		From(r.entity.Table())

	if spec != nil {
		q = q.Where(spec.Condition())

		if spec.OrderBy() != "" {
			q = q.OrderBy(spec.OrderBy())
		}
	}

	if r.limit > 0 {
		q = q.Offset(r.offset).Limit(r.limit)
	}

	query, args, _ := spec.AlterQuery(q).ToSql()

	var result []*T

	if err := r.db.Select(&result, query, args...); err != nil {
		return nil, fmt.Errorf("wat2: %w", err)
	}

	return result, nil
}
