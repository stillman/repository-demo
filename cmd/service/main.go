package main

import (
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	user2 "github.com/stillman/repository-demo/internal/domain/user"
	userSpec "github.com/stillman/repository-demo/internal/domain/user/specification"
	"github.com/stillman/repository-demo/internal/domain/user_by_stats"
	"github.com/stillman/repository-demo/internal/domain/user_by_stats/specification"
	"github.com/stillman/repository-demo/internal/repository"
)

var (
	app *App
)

type App struct {
	DB *sqlx.DB
}

func main() {
	db, err := sqlx.Open("sqlite3", "file:db.db")
	if err != nil {
		log.Fatal(err)
	}

	app = &App{
		DB: db.Unsafe(),
	}

	{
		// Save entity
		r := repository.New[user2.User](app.DB)
		user := &user2.User{}
		user.SetStatus(user2.StatusActive)
		user.SetName("some cool user")
		_ = r.Save(user) // only changed fields are updated ("name" and "status" are updated in this case)

		// PK is updated
		printJson(user.ID) // 1

		user.SetStatus(user2.StatusDisabled)
		_ = r.Save(user) // only changed fields are updated ("status" is the only updated field in this case)

		rowsDeleted, _ := r.Delete(userSpec.ByID(user.ID))
		printJson(rowsDeleted) // 1
	}

	{
		// Find single entity
		r := repository.New[user2.User](app.DB)
		user, _ := r.Find(userSpec.ByID(1))
		printJson(user) // {"ID":1,"Name":"user","Status":"active"}
	}

	{
		// Find by spec
		r := repository.New[user2.User](app.DB)
		users, _ := r.FindAll(userSpec.NewByStatus("active"))
		printJson(users) // [{"ID":2,"Name":"test user","Status":"active"},{"ID":1,"Name":"user","Status":"active"}]
	}

	{
		// Aggregation
		r := repository.New[user_by_stats.UserStatsGrouped](app.DB)
		users, _ := r.FindAll(specification.Default{})
		printJson(users) // [{"status":"active","cnt":2}]
	}
}

func printJson(el any) {
	jsonStr, _ := json.Marshal(el)
	println(string(jsonStr))
}
