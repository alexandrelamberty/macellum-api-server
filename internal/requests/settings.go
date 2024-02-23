package requests

type UpdateSettings struct {
	Name   string `json:"name" validate:"required,min=4,max=32"`
	APIKey string `json:"api_key" validate:"required, md5"`
}
