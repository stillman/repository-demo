package repository

import (
	sq "github.com/Masterminds/squirrel"
)

// sqler is used for SELECT and DELETE queries
type sqler interface {
	Table() string
	SelectFields() []string
}

// sqlerEntity is used for INSERT and UPDATE queries
type sqlerEntity interface {
	sqler
	DirtyValues() map[string]any
	ClearDirtyValues()
	SetID(int)
	PK() string
	PKValue() any
}

// SelectSpec is used for select specifications
type SelectSpec interface {
	Condition() sq.Sqlizer
	OrderBy() string
	AlterQuery(sq.SelectBuilder) sq.SelectBuilder
}

// DefaultSelectSpec can be embedded into other specifications
type DefaultSelectSpec struct {
}

func (d DefaultSelectSpec) Condition() sq.Sqlizer {
	return nil
}

func (d DefaultSelectSpec) OrderBy() string {
	return ""
}

func (d DefaultSelectSpec) AlterQuery(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder
}
