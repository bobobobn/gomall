package service

import (
	"context"

	checkout "gomall/app/frontend/hertz_gen/frontend/checkout"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	rpccheckout "gomall.local/rpc_gen/kitex_gen/checkout"
	"gomall.local/rpc_gen/kitex_gen/payment"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}
func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	_, err = rpc.CheckoutCLient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    utils.GetUserIdFromCtx(h.Context),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Address: &rpccheckout.Address{
			StreetAddress: req.Street,
			City:          req.City,
			State:         req.Province,
			Country:       req.Country,
			ZipCode:       req.Zipcode,
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
		},
	})
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"title":    "waiting",
		"redirect": "result",
	}, nil
}
