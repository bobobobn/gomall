package model

import (
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	OrderId       string    `json:"order_id"`
	Amount        float32   `json:"amount"`
	UserId        uint32    `json:"user_id"`
	TransactionId string    `json:"transaction_id"`
	PayAt         time.Time `json:"pay_at"`
}

func (PaymentLog) TableName() string {
	return "payment_logs"
}

func CreatePaymentLog(db *gorm.DB, paymentLog *PaymentLog) error {
	return db.Model(&PaymentLog{}).Create(paymentLog).Error
}
