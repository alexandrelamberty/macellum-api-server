package requests

type Category struct {
	Name    string `json:"page_size" validate:"required,min=4,max=32"`
}
