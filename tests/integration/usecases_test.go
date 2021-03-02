// +build integration

package integration

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/haunt98/togo/internal/pkg/uuid"
	"github.com/haunt98/togo/internal/services/usecases"
	"github.com/haunt98/togo/internal/storages"
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

	migration, err := migrate.New("file://migrations", connectionStr)
	s.NoError(err)
	s.migration = migration

	taskStorage := postgres.NewPostgresDB(db)
	userStorage := postgres.NewPostgresDB(db)

	s.taskUseCase = usecases.NewTaskUseCase(taskStorage, userStorage,
		uuid.Generate,
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

	_, err = s.userUseCase.Validate(context.Background(), "bla", "bla")
	s.Error(err)
}

func (s *usecasesTestSuite) TestListTasks() {
	tasks, err := s.taskUseCase.ListTasks(context.Background(), "firstUser", "2020-03-02")
	s.NoError(err)
	s.Equal(len(tasks), 0)

	tasks, err = s.taskUseCase.ListTasks(context.Background(), "abc", "2020-03-02")
	s.NoError(err)
	s.Equal(len(tasks), 1)
}

func (s *usecasesTestSuite) TestAddTask() {
	tasks, err := s.taskUseCase.ListTasks(context.Background(), "firstUser", "2020-03-02")
	s.NoError(err)
	s.Equal(len(tasks), 0)

	task, err := s.taskUseCase.AddTask(context.Background(), "firstUser", &storages.Task{
		Content: "content",
	})
	s.NoError(err)
	s.Equal(task.UserID, "firstUser")
	s.Equal(task.Content, "content")
	s.Equal(task.CreatedDate, "2020-03-02")

	tasks, err = s.taskUseCase.ListTasks(context.Background(), "firstUser", "2020-03-02")
	s.NoError(err)
	s.Equal(len(tasks), 1)
}

func (s *usecasesTestSuite) TestAddTaskReachLimit() {
	task, err := s.taskUseCase.AddTask(context.Background(), "abc", &storages.Task{
		Content: "content",
	})
	s.Equal(err, usecases.UserReachTaskLimitError)
	s.Nil(task)
}
