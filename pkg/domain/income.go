package domain

type Income struct {
	Entity
	Name string `gorm:"unique;not null" json:"name"`
}
