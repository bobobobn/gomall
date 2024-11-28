package utils

import "context"

func GetUserIdFromCtx(ctx context.Context) uint32 {
	val := ctx.Value(UserIDKey)
	if val == nil {
		return 0
	}
	return uint32(val.(int32))
}
