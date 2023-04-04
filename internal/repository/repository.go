package repository

import (
	"github.com/jmoiron/sqlx"
)

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
