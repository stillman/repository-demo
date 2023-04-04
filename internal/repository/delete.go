package repository

import (
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository[T]) Delete(conditions sq.Sqlizer) (int64, error) {
	query, args, err := sq.StatementBuilder.
		Delete(r.entity.Table()).
		Where(conditions).
		ToSql()

	if err != nil {
		return 0, err
	}

	res, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := res.RowsAffected()

	return rowsAffected, nil
}
