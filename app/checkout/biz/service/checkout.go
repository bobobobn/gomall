package service

import (
	"context"
	"gomall/app/checkout/infra/rpc"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"gomall.local/rpc_gen/kitex_gen/cart"
	checkout "gomall.local/rpc_gen/kitex_gen/checkout"
	"gomall.local/rpc_gen/kitex_gen/order"
	"gomall.local/rpc_gen/kitex_gen/payment"
	"gomall.local/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(3005001, err.Error())
	}
	var total float32 = 0.0
	var orderItems []*order.OrderItem
	for _, item := range cartResp.Items {
		p, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: int64(item.ProductId),
		})
		if err != nil {
			return nil, kerrors.NewGRPCBizStatusError(2005001, err.Error())
		}
		total += float32(item.Quantity) * p.Product.Price
		orderItems = append(orderItems, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
			},
			Cost: float32(item.Quantity) * p.Product.Price,
		})
	}

	//todo order service
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
		Email: req.Email,
		Items: orderItems,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(6005002, err.Error())
	}
	order_id := orderResp.Order.OrderId
	r, err := rpc.PaymentClient.Charge(s.ctx, &payment.ChargeReq{
		UserId:     req.UserId,
		Amount:     total,
		OrderId:    order_id,
		CreditCard: req.CreditCard,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}
	resp = &checkout.CheckoutResp{
		OrderId:       order_id,
		TransactionId: r.TransactionId,
	}
	return
}
