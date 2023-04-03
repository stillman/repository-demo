package repository

import (
	sq "github.com/Masterminds/squirrel"
)

type Spec interface {
	Condition() sq.Sqlizer
	OrderBy() string
	AlterQuery(sq.SelectBuilder) sq.SelectBuilder
}

type DefaultSpec struct {
}

func (d DefaultSpec) Condition() sq.Sqlizer {
	return nil
}

func (d DefaultSpec) OrderBy() string {
	return ""
}

func (d DefaultSpec) AlterQuery(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder
}
