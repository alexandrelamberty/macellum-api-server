package domain

type Cart struct {
	Entity
	CustomerID      uint
	Customer        Customer  `gorm:"foreignKey:CustomerID"`
	Products        []Product `gorm:"many2many:cart_products;"`
	Total           uint
	PaymentMethodID uint
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
}

type CartProduct struct {
	CartID    uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  uint
	Total     uint
}
