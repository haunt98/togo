// +build integration

package integration

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/haunt98/togo/internal/services/usecases"
	"github.com/haunt98/togo/internal/storages/postgres"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type usecasesTestSuite struct {
	suite.Suite

	db          *sql.DB
	migration   *migrate.Migrate
	taskUseCase *usecases.TaskUseCase
	userUseCase *usecases.UserUseCase
}

func TestUsecasesTestSuite(t *testing.T) {
	suite.Run(t, &usecasesTestSuite{})
}

func (s *usecasesTestSuite) SetupSuite() {
	// Init configs
	viper.SetConfigName("integration")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	s.NoError(err)

	dialect := viper.GetString("database.dialect")
	s.NotEqual(dialect, "")

	connectionStr := viper.GetString("database.connection")
	s.NotEqual(connectionStr, "")

	db, err := sql.Open(dialect, connectionStr)
	s.NoError(err)

	log.Println(connectionStr)
	migration, err := migrate.New("file://migrations", connectionStr)
	log.Println(err)
	s.NoError(err)
	s.migration = migration

	taskStorage := postgres.NewPostgresDB(db)
	userStorage := postgres.NewPostgresDB(db)

	s.taskUseCase = usecases.NewTaskUseCase(taskStorage, userStorage,
		func() string { return "taskid" },
		func() time.Time { return time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC) },
	)
	s.userUseCase = usecases.NewUserUseCase(userStorage)
}

func (s *usecasesTestSuite) SetupTest() {
	err := s.migration.Up()
	s.NoError(err)
}

func (s *usecasesTestSuite) TearDownTest() {
	err := s.migration.Down()
	s.NoError(err)
}

func (s *usecasesTestSuite) TestUserValidate() {
	valid, err := s.userUseCase.Validate(context.Background(), "firstUser", "example")
	s.NoError(err)
	s.True(valid)

	valid, err = s.userUseCase.Validate(context.Background(), "firstUser", "bla")
	s.NoError(err)
	s.False(valid)

	valid, err = s.userUseCase.Validate(context.Background(), "bla", "bla")
	s.NoError(err)
	s.False(valid)
}
