package specification

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/stillman/repository-demo/internal/domain/user"
	"github.com/stillman/repository-demo/internal/repository"
)

type Default struct {
	repository.DefaultSelectSpec
}

func (d Default) AlterQuery(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.
		RemoveColumns().
		Columns(user.FieldStatus, "COUNT(*) cnt").
		GroupBy(user.FieldStatus)
}
