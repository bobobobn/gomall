package utils

import "context"

func GetUserIdFromCtx(ctx context.Context) int {
	val := ctx.Value(UserIDKey)
	if val == nil {
		return 0
	}
	return int(val.(int32))
}
