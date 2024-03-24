package domain

type Order struct {
	Entity
	Number     string `gorm:"unique" json:"number"`
	ProviderID uint
	Provider   Provider  `gorm:"foreignKey:ProviderID"`
	Products   []Product `gorm:"many2many:order_products;"`
	Total      uint
}

type OrderProduct struct {
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  uint
	Total     uint
}
