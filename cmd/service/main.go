package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/stillman/repository-demo/internal/domain/user"
	userSpec "github.com/stillman/repository-demo/internal/domain/user/specification"
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

	r := repository.New[user.User](app.DB)
	u, err := r.Limit(10).FindAll(userSpec.NewByStatus(user.StatusActive))
	if err != nil {
		log.Fatal(err)
	}
	println(u[0].Name)
}
