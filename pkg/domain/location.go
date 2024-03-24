package domain

type Location struct {
	Entity
	Name    string  `gorm:"unique;not null" json:"name"`
	Address Address `gorm:"embedded;embeddedPrefix:address_"`
	Phone   string  `json:"phone"`
	Email   string  `json:"email"`
}
