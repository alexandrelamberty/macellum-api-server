package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Number     string `gorm:"unique" json:"number"`
	ProviderID uint
	Provider   Provider  `gorm:"foreignKey:ProviderID;"`
	Products   []Product `gorm:"many2many:order_products;"`
}

type OrderProduct struct {
	gorm.Model
	OrderID    uint
	ProductID  uint
	Quantity   uint
	TotalPrice uint
}
