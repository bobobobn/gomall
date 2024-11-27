package service

import (
	"context"
	"gomall/app/product/biz/dal/mysql"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
	product "gomall.local/rpc_gen/kitex_gen/product"
)

func TestListProducts_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewListProductsService(ctx)
	// init req and assert value

	req := &product.ListProductsReq{CategoryName: "t-shirt"}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestListProductsService_Run(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		req *product.ListProductsReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *product.ListProductsResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ListProductsService{
				ctx: tt.fields.ctx,
			}
			gotResp, err := s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListProductsService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ListProductsService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
