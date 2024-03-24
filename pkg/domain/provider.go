package domain

type Provider struct {
	Entity
	Name    string  `gorm:"unique;not null" json:"name"`
	Address Address `gorm:"embedded;embeddedPrefix:address_"`
	Website string
	Email   string
	Phone   string
	// Orders   []Order
	Products []Product
}
