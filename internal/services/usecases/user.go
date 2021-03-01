package usecases

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/haunt98/togo/internal/storages"
)

type UserUseCase struct {
	userStorage storages.UserStorage
}

func NewUserUseCase(
	userStorage storages.UserStorage,
) *UserUseCase {
	return &UserUseCase{
		userStorage: userStorage,
	}
}

func (u *UserUseCase) Validate(ctx context.Context, userID, pwd string) (bool, error) {
	userIDSql := sql.NullString{
		String: userID,
		Valid:  true,
	}

	user, err := u.userStorage.GetUser(ctx, userIDSql)
	if err != nil {
		return false, fmt.Errorf("user storage failed to get user: %w", err)
	}

	if pwd != user.Password {
		return false, nil
	}

	return true, nil
}
