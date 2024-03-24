package domain

type Group struct {
	Entity
	Name string `gorm:"unique" json:"name"`
}
