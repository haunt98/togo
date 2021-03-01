package usecases

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/haunt98/togo/internal/storages"
	mock_storages "github.com/haunt98/togo/internal/storages/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserUseCaseValidate(t *testing.T) {
	type mockGetUser struct {
		userIDSql sql.NullString
		mockUser  *storages.User
		mockErr   error
	}

	tests := []struct {
		name        string
		userID      string
		pwd         string
		mockGetUser mockGetUser
		wantResult  bool
		wantErr     error
	}{
		{
			name:   "valid",
			userID: "abc",
			pwd:    "123",
			mockGetUser: mockGetUser{
				userIDSql: sql.NullString{
					String: "abc",
					Valid:  true,
				},
				mockUser: &storages.User{
					ID:       "abc",
					Password: "123",
					MaxTodo:  5,
				},
			},
			wantResult: true,
		},
		{
			name:   "invalid",
			userID: "abc",
			pwd:    "123",
			mockGetUser: mockGetUser{
				userIDSql: sql.NullString{
					String: "abc",
					Valid:  true,
				},
				mockUser: &storages.User{
					ID:       "abc",
					Password: "456",
					MaxTodo:  5,
				},
			},
			wantResult: false,
		},
		{
			name:   "storage failed",
			userID: "abc",
			pwd:    "123",
			mockGetUser: mockGetUser{
				userIDSql: sql.NullString{
					String: "abc",
					Valid:  true,
				},
				mockErr: errors.New("storage failed"),
			},
			wantResult: false,
			wantErr:    errors.New("user storage failed to get user: storage failed"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userStorage := mock_storages.NewMockUserStorage(ctrl)
			userStorage.EXPECT().GetUser(gomock.Any(), tc.mockGetUser.userIDSql).
				Return(tc.mockGetUser.mockUser, tc.mockGetUser.mockErr)

			userUseCase := NewUserUseCase(userStorage)
			gotResult, gotErr := userUseCase.Validate(context.Background(), tc.userID, tc.pwd)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr.Error(), gotErr.Error())
				return
			}
			assert.Equal(t, tc.wantResult, gotResult)
		})
	}
}
