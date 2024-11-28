package service

import (
	"context"
	"gomall/app/payment/biz/dal/model"
	"gomall/app/payment/biz/dal/mysql"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	payment "gomall.local/rpc_gen/kitex_gen/payment"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.FormatInt(int64(req.CreditCard.CreditCardCvv), 10),
		Month:  strconv.FormatInt(int64(req.CreditCard.CreditCardExpirationMonth), 10),
		Year:   strconv.FormatInt(int64(req.CreditCard.CreditCardExpirationYear), 10),
	}
	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	}
	transac_id := uuid.New().String()
	err = model.CreatePaymentLog(mysql.DB, &model.PaymentLog{
		TransactionId: transac_id,
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}
	resp = &payment.ChargeResp{
		TransactionId: transac_id,
	}
	return
}
