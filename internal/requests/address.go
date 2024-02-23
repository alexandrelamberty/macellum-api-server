package requests

type Address struct {
	Street  string `json:"code" validate:"required,min=4,max=32"`
	Number  string `json:"first_name" validate:"required,min=1,max=3"`
	ZipCode string `json:"last_name" validate:"required,min=4,max=4"`
	City    string `json:"last_name" validate:"required,min=4,max=32"`
	State   string `json:"last_name" validate:"required,min=4,max=32"`
	Country string `json:"last_name" validate:"required,min=4,max=32"`
}
