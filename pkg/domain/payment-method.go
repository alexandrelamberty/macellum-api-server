package domain

type PaymentMethod struct {
	Entity
	Name string `gorm:"unique;not null" json:"name"`
}
