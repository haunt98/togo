package usecases

import (
	"context"
	"database/sql"

	"github.com/haunt98/togo/internal/storages"
)

type UserUseCase struct {
	userStorage storages.UserStorage
}

func (u *UserUseCase) Validate(ctx context.Context, userID, pwd string) bool {
	userIDSql := sql.NullString{}
	pwdSql := sql.NullString{}

	return u.userStorage.ValidateUser(ctx, userIDSql, pwdSql)
}

func (u *UserUseCase) CreateToken() (string, error) {
	return "", nil
}
