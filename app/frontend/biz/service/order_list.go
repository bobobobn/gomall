package service

import (
	"context"

	order "gomall/app/frontend/hertz_gen/frontend/order"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	rpcorder "gomall.local/rpc_gen/kitex_gen/order"
	rpcproduct "gomall.local/rpc_gen/kitex_gen/product"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *order.Empty) (resp map[string]any, err error) {
	listOrderResp, err := rpc.OrderClient.ListOrder(h.Context, &rpcorder.ListOrderReq{
		UserId: utils.GetUserIdFromCtx(h.Context),
	})
	if err != nil {
		return nil, err
	}
	orders := make([]map[string]any, 0)
	for _, order := range listOrderResp.Orders {
		items := make([]map[string]any, 0)
		for _, item := range order.Items {
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
				Id: int64(item.Item.ProductId),
			})
			if err != nil {
				return nil, err
			}
			items = append(items, map[string]any{
				"Picture":     productResp.Product.Picture,
				"ProductName": productResp.Product.Name,
				"Qty":         item.Item.Quantity,
				"Cost":        item.Cost,
			})
		}
		orders = append(orders, map[string]any{
			"OrderId":     order.OrderId,
			"CreatedDate": order.CreatedAt,
			"Items":       items,
		})
	}
	return map[string]any{
		"orders": orders,
	}, nil
}
