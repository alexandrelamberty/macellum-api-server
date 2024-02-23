package requests

type RegisterUser struct {
	FirstName string `json:"first_name" validate:"required,min=4,max=32"`
	LastName  string `json:"last_name" validate:"required,min=4,max=32"`
	Email     string `json:"email" validate:"required,email"`
	// TODO: validate regex
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UpdateUser struct {
	FirstName string `json:"first_name" validate:"required,min=4,max=32"`
	LastName  string `json:"last_name" validate:"required,min=4,max=32"`
}
