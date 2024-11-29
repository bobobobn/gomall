package model

import (
	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       string
}

type Order struct {
	gorm.Model
	OrderId    string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId     uint32      `gorm:"type:int(11)"`
	Consignee  Consignee   `gorm:"embedded"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
