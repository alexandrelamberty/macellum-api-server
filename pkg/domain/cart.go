package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CustomerID uint
	Customer   Customer  `gorm:"foreignKey:CustomerID"`
	Products   []Product `gorm:"many2many:cart_products;"`
	Total      uint
}

type CartProduct struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Quantity  uint
	Total     uint
}
