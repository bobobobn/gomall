package middleware

import (
	"context"
	"gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		user_id := session.Get("user_id")
		if user_id == nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIDKey, user_id.(int32))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		user_id := session.Get("user_id")
		if user_id == nil {
			c.Redirect(302, []byte("/login?next="+c.FullPath()))
			c.Abort()
			return
		}
		ctx = context.WithValue(ctx, utils.UserIDKey, user_id.(int32))
		c.Next(ctx)
	}
}
