package domain

type Appointment struct {
	Entity
	CustomerID      uint
	Customer        Customer  `gorm:"foreignKey:CustomerID"`
	Date            string    `gorm:"not null"`
}

