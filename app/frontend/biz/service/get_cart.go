package service

import (
	"context"
	"strconv"

	cart "gomall/app/frontend/hertz_gen/frontend/cart"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	rpccart "gomall.local/rpc_gen/kitex_gen/cart"
	rpcproduct "gomall.local/rpc_gen/kitex_gen/product"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

type itemsCart struct {
	Id       int64
	Name     string
	Price    string
	Quantity int32
	Picture  string
}

func (h *GetCartService) Run(req *cart.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	r, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint32(utils.GetUserIdFromCtx(h.Context))})
	var items []itemsCart
	var totalPrice float64 = 0.0
	for _, item := range r.Items {
		p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: int64(item.ProductId)})
		if err != nil {
			return nil, err
		}

		items = append(items, itemsCart{
			Id:       int64(item.ProductId),
			Name:     p.Product.Name,
			Price:    strconv.FormatFloat(float64(p.Product.Price), 'f', 2, 64),
			Quantity: int32(item.Quantity),
			Picture:  p.Product.Picture,
		})
		totalPrice += float64(p.Product.Price) * float64(item.Quantity)
	}
	resp = map[string]any{
		"items": items,
		"total": strconv.FormatFloat(totalPrice, 'f', 2, 64),
	}
	return
}
