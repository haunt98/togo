package storages

import (
	"context"
	"database/sql"
)

type TaskStorage interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
}