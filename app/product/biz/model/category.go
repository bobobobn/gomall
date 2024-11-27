package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"product" gorm:"many2many:product_category"`
}

func (c *Category) TableName() string {
	return "categories"
}

func GetProductsByCategoryName(db *gorm.DB, name string) (c *Category, err error) {
	err = db.Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&c).Error
	return
}
