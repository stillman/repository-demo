package user_by_stats

import (
	"github.com/stillman/repository-demo/internal/domain/user"
)

type UserStatsGrouped struct {
	user.Model
	Status string `json:"status" db:"status"`
	Cnt    int    `json:"cnt" db:"cnt"`
}

func (u UserStatsGrouped) SelectFields() []string {
	return []string{"status", "COUNT(*) AS cnt"}
}

func (u UserStatsGrouped) Table() string {
	return "user"
}
