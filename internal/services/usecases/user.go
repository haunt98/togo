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
	pwdSql := sql.NullString{
		String: pwd,
		Valid:  true,
	}

	valid, err := u.userStorage.ValidateUser(ctx, userIDSql, pwdSql)
	if err != nil {
		return false, fmt.Errorf("user storage failed to validate: %w", err)
	}

	return valid, nil
}
