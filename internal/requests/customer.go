package requests

type CreateCustomer struct {
	Code      string  `json:"code" validate:"required,min=4,max=32"`
	FirstName string  `json:"first_name" validate:"required,min=4,max=32"`
	LastName  string  `json:"last_name" validate:"required,min=4,max=32"`
	BirthDate string  `json:"birth_date" validate:"required"`
	Address   Address `json:"address" validate:"required"`
	Phone     string  `json:"phone" validate:"required,min=10,max=10"`
	Email     string  `json:"email" validate:"required,email"`
	Notes     string  `json:"notes" validate:"required"`
	Status    string  `json:"status" validate:"required"`
	Label     string  `json:"label" validate:"required"`
	Priority  int     `json:"priority" validate:"required"`
}
