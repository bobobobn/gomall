package service

import (
	"context"
	"gomall/rpc_gen/biz/dal/mysql"
	user "gomall/rpc_gen/kitex_gen/user"
	"gomall/rpc_gen/model"

	"github.com/pingcap/errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}
	// Save user to database.
	err = model.CreateUser(mysql.DB, u)
	if err != nil {
		return nil, err
	}
	resp = new(user.RegisterResp)
	resp.UserId = int32(u.ID)
	return
}
