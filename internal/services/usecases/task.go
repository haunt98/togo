package usecases

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/haunt98/togo/internal/pkg/clock"
	"github.com/haunt98/togo/internal/pkg/uuid"
	"github.com/haunt98/togo/internal/storages"
)

type TaskUseCase struct {
	taskStorage    storages.TaskStorage
	uuidGenerateFn uuid.GenerateFn
	nowFn          clock.NowFn
}

func NewTaskUseCase(
	taskStorage storages.TaskStorage,
	uuidGenerateFn uuid.GenerateFn,
	nowFn clock.NowFn,
) *TaskUseCase {
	return &TaskUseCase{
		taskStorage:    taskStorage,
		uuidGenerateFn: uuid.Generate,
		nowFn:          clock.Now,
	}
}

func (u *TaskUseCase) ListTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	userIDSql := sql.NullString{
		String: userID,
		Valid:  true,
	}
	createdDateSql := sql.NullString{
		String: createdDate,
		Valid:  true,
	}

	tasks, err := u.taskStorage.RetrieveTasks(ctx, userIDSql, createdDateSql)
	if err != nil {
		return nil, fmt.Errorf("task storage failed to retrieve tasks of userid %s createdDate %s: %w", userID, createdDate, err)
	}

	return tasks, nil
}

func (u *TaskUseCase) AddTask(ctx context.Context, userID string, task *storages.Task) (*storages.Task, error) {
	// Update task
	// Skip userID in task
	task.UserID = userID
	task.ID = u.uuidGenerateFn()
	task.CreatedDate = u.nowFn().Format(clock.DateFormat)

	if err := u.taskStorage.AddTask(ctx, task); err != nil {
		return nil, fmt.Errorf("task storage failed to add: %w", err)
	}

	return task, nil
}
