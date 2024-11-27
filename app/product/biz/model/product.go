package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string     `json:"name" gorm:"type:varchar(128)"`
	Description string     `json:"description" gorm:"type:varchar(256)"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p *Product) TableName() string {
	return "products"
}

func GetProductById(db *gorm.DB, id int64) (p Product, err error) {
	err = db.Model(&Product{}).Where("id =?", id).First(&p).Error
	return
}

func SearchProduct(db *gorm.DB, q string) (products []*Product, err error) {
	err = db.Model(&Product{}).Where("name LIKE? or description LIKE?",
		"%"+q+"%", "%"+q+"%",
	).Find(&products).Error
	return
}
