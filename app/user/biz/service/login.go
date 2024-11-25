package service

import (
	"context"
	"gomall/rpc_gen/biz/dal/mysql"
	user "gomall/rpc_gen/kitex_gen/user"
	"gomall/rpc_gen/model"

	"github.com/pingcap/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	u, err := model.GetUserByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHashed), []byte(req.Password)) != nil {
		return nil, errors.New("password is incorrect")
	}
	resp = &user.LoginResp{
		UserId: int32(u.ID),
	}
	return resp, nil
}
