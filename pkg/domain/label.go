package domain

type Label struct {
	Entity
	Name    string `gorm:"unique;not null" json:"name"`
	Size uint   `gorm:"not null"`
	Type string `gorm:"not null"` // codebar, label or qrcode
}
