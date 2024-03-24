package domain

type Limit struct {
	Entity
	Name    string `gorm:"unique;not null" json:"name"`
	Peoples uint   `gorm:"not null"`
	Days    uint   `gorm:"not null"`
	Amount  uint   `gorm:"not null"`
}
