package tododto

type TodoResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	File        string `json:"file"`
}
