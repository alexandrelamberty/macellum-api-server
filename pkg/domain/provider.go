package domain

import "gorm.io/gorm"

type Provider struct {
	gorm.Model
	Name    string  `gorm:"unique;not null" json:"name"`
	Address Address `gorm:"embedded;embeddedPrefix:address_"`
	Website string
	Phone   string
	Orders  []Order ` gorm:"foreignKey:ProviderID"`
	// Products []Product ` gorm:"foreignKey:ProviderID"`
}
