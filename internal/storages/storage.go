package storages

import (
	"context"
)

//go:generate mockgen -source=storage.go -destination=gomock/storage.go

type TaskStorage interface {
	RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
}

type UserStorage interface {
	GetUser(ctx context.Context, userID string) (*User, error)
}
