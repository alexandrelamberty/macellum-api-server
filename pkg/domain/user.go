package domain

type User struct {
	Entity
	FirstName string `gorm:"required" json:"first_name"`
	LastName  string `gorm:"required" json:"last_name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `gorm:"required" json:"-"`
}
