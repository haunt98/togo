package storages

import (
	"context"
	"database/sql"
)

type Storage interface {
	// Task
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
	// User
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}
