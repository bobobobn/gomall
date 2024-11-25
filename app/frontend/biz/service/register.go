package service

import (
	"context"
	"errors"

	auth "gomall/app/frontend/hertz_gen/frontend/auth"
	"gomall/app/frontend/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"gomall.local/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.SignUpReq) (resp *auth.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	if req.Email == "" || req.Password == "" || req.PasswordConfirm != req.Password {
		return nil, errors.New("invalid input")
	}
	registerReq := user.RegisterReq{
		Email:    req.Email,
		Password: req.Password,
	}
	r, err := rpc.UserClient.Register(context.Background(), &registerReq)
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
