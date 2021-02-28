package storages

import (
	"context"
)

type UserStorage interface {
	ValidateUser(ctx context.Context, userID, pwd string) bool
}
