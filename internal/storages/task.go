package storages

import (
	"context"
)

type TaskStorage interface {
	RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
}
