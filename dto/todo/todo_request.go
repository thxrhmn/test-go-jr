package tododto

type TodoRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	File        string `json:"file" form:"file"`
}
