package service

import (
	"context"
	"gomall/rpc_gen/biz/dal/mysql"
	user "gomall/rpc_gen/kitex_gen/user"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

func TestLogin_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		Email:    "test@te.com",
		Password: "12346",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestLoginService_Run(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		req *user.LoginReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *user.LoginResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LoginService{
				ctx: tt.fields.ctx,
			}
			gotResp, err := s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("LoginService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
