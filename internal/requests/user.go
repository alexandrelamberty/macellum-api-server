package requests

type InviteUser struct {
	FirstName string `json:"first_name" validate:"required,min=4,max=32"`
	LastName  string `json:"last_name" validate:"required,min=4,max=32"`
	Email     string `json:"email" validate:"required,min=4,max=32"`
}

type UpdateUser struct {
	FirstName string `json:"first_name" validate:"required,min=4,max=32"`
	LastName  string `json:"last_name" validate:"required,min=4,max=32"`
	Email     string `json:"email" validate:"required,min=4,max=32"`
}
