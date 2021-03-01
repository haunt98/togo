package usecases

import (
	"context"
	"database/sql"

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

func (u *UserUseCase) Validate(ctx context.Context, userID, pwd string) bool {
	userIDSql := sql.NullString{
		String: userID,
		Valid:  true,
	}
	pwdSql := sql.NullString{
		String: pwd,
		Valid:  true,
	}

	return u.userStorage.ValidateUser(ctx, userIDSql, pwdSql)
}
