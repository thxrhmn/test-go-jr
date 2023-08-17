package tododto

type SubTodoRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	File        string `json:"file" form:"file"`
	TodoID      int    `json:"todo_id" form:"todo_id" validate:"required"`
}
