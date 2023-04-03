package specification

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/stillman/repository-demo/internal/domain/user"
	"github.com/stillman/repository-demo/internal/repository"
)

type Default struct {
	repository.DefaultSpec
}

func (d Default) AlterQuery(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.
		GroupBy(user.FieldStatus).
		Having("cnt > 3")
}
