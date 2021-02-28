package usecases

import (
	"context"
	"fmt"

	"github.com/haunt98/togo/internal/storages"
)

type TaskUseCase struct {
	taskStorage storages.TaskStorage
}

func (u *TaskUseCase) List(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	tasks, err := u.taskStorage.RetrieveTasks(ctx, userID, createdDate)
	if err != nil {
		return nil, fmt.Errorf("task storage failed to retrieve tasks of userid %s createdDate %s: %w", userID, createdDate)
	}

	return tasks, nil
}

func (u *TaskUseCase) Add(ctx context.Context, task *storages.Task) (*storages.Task, error) {
	// TODO: generate uuid

	// TODO: generate created date

	if err := u.taskStorage.AddTask(ctx, task); err != nil {
		return nil, fmt.Errorf("task storage failed to add: %w", err)
	}

	return task, nil
}
