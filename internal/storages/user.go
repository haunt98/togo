package storages

import (
	"context"
	"database/sql"
)

type UserStorage interface {
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}
