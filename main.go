package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/haunt98/togo/internal/pkg/clock"
	"github.com/haunt98/togo/internal/pkg/uuid"
	"github.com/haunt98/togo/internal/services/transports"
	"github.com/haunt98/togo/internal/services/usecases"
	"github.com/haunt98/togo/internal/storages"
	"github.com/haunt98/togo/internal/storages/postgres"
	"github.com/haunt98/togo/internal/token/jwt"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

const (
	postgresDialect = "postgres"
)

func main() {
	// Init configs
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("viper failed to read config", err)
	}

	// Storage -> Use case -> Transport
	taskStorage, userStorage := initStorage()
	taskUseCase, userUseCase := initUseCase(taskStorage, userStorage)
	transport := initTransport(taskUseCase, userUseCase)

	port := viper.GetInt("service.port")
	if port == 0 {
		log.Fatal("invalid service.port")
	}
	log.Printf("running with port %d", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), transport); err != nil {
		log.Fatalf("failed to start service: %s", err)
	}
}

func initStorage() (storages.TaskStorage, storages.UserStorage) {
	dialect := viper.GetString("database.dialect")
	if dialect == "" {
		log.Fatal("invalid database.dialect")
	}

	connectionStr := viper.GetString("database.connection")
	if connectionStr == "" {
		log.Fatal("invalid database.connection")
	}

	db, err := sql.Open(dialect, connectionStr)
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %s", err)
	}

	var taskStorage storages.TaskStorage
	var userStorage storages.UserStorage

	switch dialect {
	case postgresDialect:
		taskStorage = postgres.NewPostgresDB(db)
		userStorage = postgres.NewPostgresDB(db)
	default:
		log.Fatalf("unsupport dialect %s", dialect)
	}

	return taskStorage, userStorage
}

func initUseCase(
	taskStorage storages.TaskStorage,
	userStorage storages.UserStorage,
) (*usecases.TaskUseCase, *usecases.UserUseCase) {
	taskUseCase := usecases.NewTaskUseCase(taskStorage, uuid.Generate, clock.Now)
	userUseCase := usecases.NewUserUseCase(userStorage)
	return taskUseCase, userUseCase
}

func initTransport(
	taskUseCase *usecases.TaskUseCase,
	userUseCase *usecases.UserUseCase,
) *transports.Transport {
	jwtKey := viper.GetString("jwt.key")
	if jwtKey == "" {
		log.Fatal("invalid jwt.key")
	}
	jwtGenerator := jwt.NewGenerator(jwtKey)

	taskTransport := transports.NewTaskTransport(taskUseCase)
	userTransport := transports.NewUserTransport(userUseCase, jwtGenerator)
	transport := transports.NewTransport(taskTransport, userTransport)
	return transport
}
