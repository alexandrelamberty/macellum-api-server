package domain

type Category struct {
	Entity
	Name string `gorm:"unique" json:"name"`
}
