package domain

type Customer struct {
	Entity
	Code      string `gorm:"unique"`
	FirstName string
	LastName  string
	BirthDate string
	Address   Address `gorm:"embedded;embeddedPrefix:address_"`
	Phone     string
	Email     string
	Notes     string
	Status    string
	Label     string
	Priority  int
	Carts     []Cart  `gorm:"foreignKey:CustomerID"`
	Events    []Event `gorm:"foreignKey:CustomerID"`
}
