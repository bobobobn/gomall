package service

import (
	"context"

	checkout "gomall/app/frontend/hertz_gen/frontend/checkout"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	rpccart "gomall.local/rpc_gen/kitex_gen/cart"
	rpcproduct "gomall.local/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.Empty) (resp map[string]any, err error) {
	cartResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: utils.GetUserIdFromCtx(h.Context),
	})
	if err != nil {
		return nil, err
	}
	var total float32 = 0.0
	var items []map[string]any
	for _, item := range cartResp.Items {
		p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: int64(item.ProductId),
		})
		if err != nil {
			return nil, err
		}
		total += float32(item.Quantity) * p.Product.Price
		items = append(items, map[string]any{
			"Name":    p.Product.Name,
			"Qty":     item.Quantity,
			"Price":   p.Product.Price,
			"Picture": p.Product.Picture,
		})
	}

	resp = map[string]any{
		"total": total,
		"items": items,
	}

	return
}
