package service

import (
	"context"
	"gomall/app/cart/biz/dal/mysql"
	"gomall/app/cart/biz/model"

	cart "gomall.local/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	r, err := model.GetCartByUserID(mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	carts := []*cart.CartItem{}
	for _, item := range r {
		carts = append(carts, &cart.CartItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		})
	}
	resp = &cart.GetCartResp{
		Items: carts,
	}

	return
}
