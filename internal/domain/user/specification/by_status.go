package userSpec

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/stillman/repository-demo/internal/domain/user"
	"github.com/stillman/repository-demo/internal/repository"
)

func NewByStatus(status string) ByStatus {
	return ByStatus{status: status}
}

type ByStatus struct {
	repository.DefaultSelectSpec
	status string
}

func (b ByStatus) Condition() sq.Sqlizer {
	return sq.Eq{user.FieldStatus: b.status}
}

func (b ByStatus) OrderBy() string {
	return "id DESC"
}
