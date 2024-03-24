package domain

type Price struct {
	Entity
	ProductID uint    `gorm:"not null" json:"-"`
	Price     float64 `gorm:"not null" json:"price"`
	Date      string  `gorm:"not null" json:"date"`
}

type Weight struct {
	Value int    `gorm:"not null" json:"value"`
	Unit  string `gorm:"not null" json:"unit"`
}

type Stock struct {
	Shelf    string `gorm:"not null" json:"shelf"`
	Quantity int    `gorm:"not null" json:"quantity"`
	Minimum  int    `gorm:"not null" json:"minimum"`
	Maximum  int    `gorm:"not null" json:"maximum"`
	InStock  bool   `gorm:"not null" json:"in_stock"`
}

type Store struct {
	Shelf    string `gorm:"not null" json:"shelf"`
	Quantity int    `gorm:"not null" json:"quantity"`
	Minimum  int    `gorm:"not null" json:"minimum"`
	Maximum  int    `gorm:"not null" json:"maximum"`
	InStore  bool   `gorm:"not null" json:"in_store"`
}

type Product struct {
	Entity
	Code        string  `gorm:"unique;not null" json:"code"`
	Name        string  `gorm:"not null" json:"name"`
	Brand       string  `gorm:"not null" json:"brand"`
	Prices      []Price `json:"prices"`
	Weight      Weight  `gorm:"embedded;embeddedPrefix:weight_" json:"weight"`
	Stock       Stock   `gorm:"embedded;embeddedPrefix:stock_" json:"stock"`
	Store       Store   `gorm:"embedded;embeddedPrefix:store_" json:"store"`
	ProviderID  uint    `gorm:"not null" json:"-"`
	Provider    Provider
	CategoryID  uint `gorm:"not null" json:"-"`
	Category    Category
	GroupID     uint `gorm:"not null" json:"-"`
	Group       Group
	ProductType string `gorm:"not null" json:"product_type"`
}
