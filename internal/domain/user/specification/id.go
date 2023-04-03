package userSpec

import (
	sq "github.com/Masterminds/squirrel"
)

func ByID(id int) sq.Sqlizer {
	return sq.Eq{"id": id}
}
