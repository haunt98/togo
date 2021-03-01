package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/haunt98/togo/internal/services"
	sqllite "github.com/haunt98/togo/internal/storages/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	// Storage layer

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	})
}
