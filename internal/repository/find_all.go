package repository

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r *Repository[T]) FindAll(spec SelectSpec) ([]*T, error) {
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
