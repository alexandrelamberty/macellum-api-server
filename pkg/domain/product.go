package domain

import "gorm.io/gorm"

type Weight struct {
	Value int
	Unit  string
}

type Stock struct {
	Shelf    string
	Quantity int
	InStore  bool
}

type Store struct {
	Shelf    string
	Quantity int
	InStock  bool
}

type Price struct {
	IncludeVAT float64
	ExcludeVAT float64
	Currency   string
}

type Product struct {
	gorm.Model
	SKU         string `gorm:"unique;not null" json:"sku"`
	Name        string
	Brand       string
	Price       Price      `gorm:"embedded"`
	Weight      Weight     `gorm:"embedded"`
	Stock       Stock      `gorm:"embedded"`
	Store       Store      `gorm:"embedded"`
	Providers   []Provider `gorm:"many2many:product_providers;"`
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
	ProductType string
}

type ProductProvider struct {
	gorm.Model
	ProductID  uint
	ProviderID uint
}
