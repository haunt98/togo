package usecases

import (
	"context"

	"github.com/haunt98/togo/internal/storages"
)

type UserUseCase struct {
	userStorage storages.UserStorage
}

func (u *UserUseCase) Validate(ctx context.Context, userID, pwd string) bool {
	if userID == "" || pwd == "" {
		return false
	}

	return u.userStorage.ValidateUser(ctx, userID, pwd)
}

func (u *UserUseCase) CreateToken() (string, error) {
	return "", nil
}
