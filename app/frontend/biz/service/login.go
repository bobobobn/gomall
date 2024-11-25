package service

import (
	"context"

	auth "gomall/app/frontend/hertz_gen/frontend/auth"
	"gomall/app/frontend/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	user "gomall.local/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp *auth.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	loginReq := &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	}
	r, err := rpc.UserClient.Login(context.Background(), loginReq)
	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", r.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
