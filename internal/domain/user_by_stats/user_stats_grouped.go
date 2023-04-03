package user_by_stats

import (
	"github.com/stillman/repository-demo/internal/domain/user"
)

type UserStatsGrouped struct {
	user.Model
	Status string `db:"status"`
	Count  int    `db:"cnt"`
}

func (u UserStatsGrouped) SelectFields() []string {
	return []string{"status", "COUNT(*) AS cnt"}
}
