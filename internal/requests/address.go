package requests

type Address struct {
	Street  string `json:"street" validate:"required,min=4,max=32"`
	Number  string `json:"number" validate:"required,min=1,max=3"`
	ZipCode string `json:"zip_code" validate:"required,min=4,max=4"`
	City    string `json:"city" validate:"required,min=4,max=32"`
	State   string `json:"state" validate:"required,min=4,max=32"`
	Country string `json:"country" validate:"required,min=4,max=32"`
}
