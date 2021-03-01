package usecases

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			taskStorage := mock_storages.NewMockTaskStorage(ctrl)
			taskStorage.EXPECT().RetrieveTasks(gomock.Any(), tc.mockRetriveTasks.userID, tc.mockRetriveTasks.createdDate).Return(tc.mockRetriveTasks.mockTasks, tc.mockRetriveTasks.mockErr)

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
