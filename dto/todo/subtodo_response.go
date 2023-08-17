package tododto

type SubTodoResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	File        string `json:"file"`
	TodoID      int    `json:"todo_id"`
}
