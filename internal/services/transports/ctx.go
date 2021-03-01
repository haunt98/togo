package transports

import (
	"context"
	"errors"
)

var (
	getUserIDFromCtxError = errors.New("failed to get userID from ctx")
)

func getUserIDFromCtx(ctx context.Context) (string, error) {
	v := ctx.Value(userIDField)
	id, ok := v.(string)
	if !ok {
		return "", getUserIDFromCtxError
	}
	return id, nil
}
