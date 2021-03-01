package storages

import (
	"context"
	"database/sql"
)

type TaskStorage interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
}

type UserStorage interface {
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) (bool, error)
}
