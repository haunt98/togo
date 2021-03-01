package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/haunt98/togo/internal/pkg/clock"
	"github.com/haunt98/togo/internal/pkg/uuid"
	"github.com/haunt98/togo/internal/services/transports"
	"github.com/haunt98/togo/internal/services/usecases"
	"github.com/haunt98/togo/internal/storages"
	"github.com/haunt98/togo/internal/token/jwt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	// Storage layer
	taskStorage := storages.NewTaskDB(db)
	userStorage := storages.NewUserDB(db)

	// Use case layer
	taskUseCase := usecases.NewTaskUseCase(taskStorage, uuid.Generate, clock.Now)
	userUseCase := usecases.NewUserUseCase(userStorage)

	// Transport layer
	taskTransport := transports.NewTaskTransport(taskUseCase)
	userTransport := transports.NewUserTransport(userUseCase, jwt.NewGenerator("wqGyEBBfPK9w3Lxw"))
	transport := transports.NewTransport(taskTransport, userTransport)

	http.ListenAndServe(":8080", transport)
}
