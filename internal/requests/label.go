package requests

type CreateLabel struct {
	PageSize    string `json:"page_size" validate:"required,min=4,max=32"`
	Orientation string `json:"orientation" validate:"required,min=1,max=1"`
	LabelFormat string `json:"label_format" validate:"required,min=4,max=32"`
}
