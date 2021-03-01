package usecases

import (
	"context"
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
	user, err := u.userStorage.GetUser(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("user storage failed to get user: %w", err)
	}

	if pwd != user.Password {
		return false, nil
	}

	return true, nil
}
