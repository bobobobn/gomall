package model

import (
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint32 `json:"user_id" gorm:"type:int(11);not null;index:idx_user_id"`
	ProductID uint32 `json:"product_id" gorm:"type:int(11);not null;"`
	Quantity  uint32 `json:"quantity" gorm:"type:int(11);not null;default:1"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddCart(db *gorm.DB, item *Cart) error {
	return db.Create(item).Error
}

func GetCartByUserID(db *gorm.DB, userID uint32) ([]Cart, error) {
	var carts []Cart
	err := db.Where("user_id =?", userID).Find(&carts).Error
	return carts, err
}

func EmptyCart(db *gorm.DB, userID uint32) error {
	if userID == 0 {
		return errors.New("user_id cannot be 0")
	}
	return db.Where("user_id =?", userID).Delete(&Cart{}).Error
}
