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
	"github.com/haunt98/togo/internal/token/jwt"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func main() {
	// Init configs
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("viper failed to read config", err)
	}

	// Storage layer
	db := initDatabase()
	taskStorage := storages.NewTaskDB(db)
	userStorage := storages.NewUserDB(db)

	// Use case layer
	taskUseCase := usecases.NewTaskUseCase(taskStorage, uuid.Generate, clock.Now)
	userUseCase := usecases.NewUserUseCase(userStorage)

	// Init JWT
	jwtKey := viper.GetString("jwt.key")
	if jwtKey == "" {
		log.Fatal("invalid jwt.key")
	}
	jwtGenerator := jwt.NewGenerator(jwtKey)

	// Transport layer
	taskTransport := transports.NewTaskTransport(taskUseCase)
	userTransport := transports.NewUserTransport(userUseCase, jwtGenerator)
	transport := transports.NewTransport(taskTransport, userTransport)

	port := viper.GetInt("service.port")
	if port == 0 {
		log.Fatal("invalid service.port")
	}
	log.Printf("running with port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), transport)
}

func initDatabase() *sql.DB {
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
		log.Fatal("failed to open database", err)
	}

	return db
}
