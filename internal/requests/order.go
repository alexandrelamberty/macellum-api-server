package requests

type CreateOrder struct {
	Code       string `json:"code" validate:"required,min=4,max=32"`
	ProviderID string `json:"last_name" validate:"required,min=4,max=32"`
}
