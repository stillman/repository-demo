package repository

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

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
