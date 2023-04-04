package repository

import (
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository[T]) Save(entity sqlerEntity) error {
	if entity.PKValue() == 0 {
		return r.Insert(entity)
	}

	return r.Update(entity)
}

func (r *Repository[T]) Insert(entity sqlerEntity) error {
	query, args, err := sq.StatementBuilder.
		Insert(entity.Table()).
		SetMap(entity.DirtyValues()).
		ToSql()

	if err != nil {
		return err
	}

	res, err := r.db.Exec(query, args...)

	if err != nil {
		return err
	}

	insertID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	entity.SetID(int(insertID))
	entity.ClearDirtyValues()

	return nil
}

func (r *Repository[T]) Update(entity sqlerEntity) error {
	query, args, err := sq.StatementBuilder.
		Update(entity.Table()).
		SetMap(entity.DirtyValues()).
		Where(sq.Eq{entity.PK(): entity.PKValue()}).
		ToSql()

	if err != nil {
		return err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		return err
	}

	entity.ClearDirtyValues()

	return nil
}
