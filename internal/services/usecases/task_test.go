package usecases

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/haunt98/togo/internal/pkg/clock"
	"github.com/haunt98/togo/internal/pkg/uuid"
	"github.com/haunt98/togo/internal/storages"
	mock_storages "github.com/haunt98/togo/internal/storages/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTaskUseCaseListTasks(t *testing.T) {
	type mockRetriveTasks struct {
		userID      string
		createdDate string
		mockTasks   []*storages.Task
		mockErr     error
	}

	tests := []struct {
		name             string
		userID           string
		createdDate      string
		mockRetriveTasks mockRetriveTasks
		wantResult       []*storages.Task
		wantErr          error
	}{
		{
			name:        "ok",
			userID:      "abc",
			createdDate: "2020-03-01",
			mockRetriveTasks: mockRetriveTasks{
				userID:      "abc",
				createdDate: "2020-03-01",
				mockTasks: []*storages.Task{
					{
						Content: "abc",
					},
					{
						Content: "def",
					},
				},
			},
			wantResult: []*storages.Task{
				{
					Content: "abc",
				},
				{
					Content: "def",
				},
			},
		},
		{
			name:        "database failed",
			userID:      "abc",
			createdDate: "2020-03-01",
			mockRetriveTasks: mockRetriveTasks{
				userID:      "abc",
				createdDate: "2020-03-01",
				mockErr:     errors.New("database failed"),
			},
			wantErr: errors.New("task storage failed to retrieve tasks of userid abc createdDate 2020-03-01: database failed"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			taskStorage := mock_storages.NewMockTaskStorage(ctrl)

			taskStorage.EXPECT().RetrieveTasks(gomock.Any(), tc.mockRetriveTasks.userID, tc.mockRetriveTasks.createdDate).
				Return(tc.mockRetriveTasks.mockTasks, tc.mockRetriveTasks.mockErr)

			taskUseCase := NewTaskUseCase(taskStorage, nil, nil, nil)
			gotResult, gotErr := taskUseCase.ListTasks(context.Background(), tc.userID, tc.createdDate)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr.Error(), gotErr.Error())
				return
			}
			assert.Equal(t, tc.wantResult, gotResult)
		})
	}
}

func TestTaskUseCaseAddTask(t *testing.T) {
	type mockRetriveTasks struct {
		userID      string
		createdDate string
		mockTasks   []*storages.Task
		mockErr     error
	}

	type mockGetUser struct {
		enable   bool
		userID   string
		mockUser *storages.User
		mockErr  error
	}

	type mockAddTask struct {
		enable  bool
		task    *storages.Task
		mockErr error
	}

	tests := []struct {
		name               string
		userID             string
		task               *storages.Task
		mockRetriveTasks   mockRetriveTasks
		mockGetUser        mockGetUser
		mockAddTask        mockAddTask
		mockNowFn          clock.NowFn
		mockUUIDGenerateFn uuid.GenerateFn
		wantResult         *storages.Task
		wantErr            error
	}{
		{
			name:   "ok",
			userID: "abc",
			task: &storages.Task{
				Content: "content",
			},
			mockRetriveTasks: mockRetriveTasks{
				userID:      "abc",
				createdDate: "2020-03-02",
				mockTasks:   []*storages.Task{},
			},
			mockGetUser: mockGetUser{
				enable: true,
				userID: "abc",
				mockUser: &storages.User{
					ID:      "abc",
					MaxTodo: 2,
				},
			},
			mockAddTask: mockAddTask{
				enable: true,
				task: &storages.Task{
					ID:          "taskid",
					Content:     "content",
					UserID:      "abc",
					CreatedDate: "2020-03-02",
				},
			},
			mockNowFn:          func() time.Time { return time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC) },
			mockUUIDGenerateFn: func() string { return "taskid" },
			wantResult: &storages.Task{
				ID:          "taskid",
				Content:     "content",
				UserID:      "abc",
				CreatedDate: "2020-03-02",
			},
		},
		{
			name:   "database failed",
			userID: "abc",
			task: &storages.Task{
				Content: "content",
			},
			mockRetriveTasks: mockRetriveTasks{
				userID:      "abc",
				createdDate: "2020-03-02",
				mockErr:     errors.New("database failed"),
			},
			mockGetUser: mockGetUser{
				enable: false,
			},
			mockAddTask: mockAddTask{
				enable: false,
			},
			mockNowFn:          func() time.Time { return time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC) },
			mockUUIDGenerateFn: func() string { return "taskid" },
			wantErr:            errors.New("task storage failed to retrieve tasks of userid abc createdDate 2020-03-02: database failed"),
		},
		{
			name:   "reach limit",
			userID: "abc",
			task: &storages.Task{
				Content: "content",
			},
			mockRetriveTasks: mockRetriveTasks{
				userID:      "abc",
				createdDate: "2020-03-02",
				mockTasks: []*storages.Task{
					{
						Content: "content 1",
					},
					{
						Content: "content 2",
					},
				},
			},
			mockGetUser: mockGetUser{
				enable: true,
				userID: "abc",
				mockUser: &storages.User{
					ID:      "abc",
					MaxTodo: 2,
				},
			},
			mockAddTask: mockAddTask{
				enable: false,
			},
			mockNowFn:          func() time.Time { return time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC) },
			mockUUIDGenerateFn: func() string { return "taskid" },
			wantErr:            UserReachTaskLimitError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			taskStorage := mock_storages.NewMockTaskStorage(ctrl)
			userStorage := mock_storages.NewMockUserStorage(ctrl)

			taskStorage.EXPECT().RetrieveTasks(gomock.Any(), tc.mockRetriveTasks.userID, tc.mockRetriveTasks.createdDate).
				Return(tc.mockRetriveTasks.mockTasks, tc.mockRetriveTasks.mockErr)

			if tc.mockGetUser.enable {
				userStorage.EXPECT().GetUser(gomock.Any(), tc.mockGetUser.userID).
					Return(tc.mockGetUser.mockUser, tc.mockGetUser.mockErr)
			}

			if tc.mockAddTask.enable {
				taskStorage.EXPECT().AddTask(gomock.Any(), tc.mockAddTask.task).
					Return(tc.mockAddTask.mockErr)
			}

			taskUseCase := NewTaskUseCase(taskStorage, userStorage, tc.mockUUIDGenerateFn, tc.mockNowFn)
			gotResult, gotErr := taskUseCase.AddTask(context.Background(), tc.userID, tc.task)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr.Error(), gotErr.Error())
			}
			assert.Equal(t, tc.wantResult, gotResult)
		})
	}
}
