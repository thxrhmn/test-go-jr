package routes

import (
	"todo-list-thxrhmn/handlers"
	"todo-list-thxrhmn/pkg/postgresql"
	"todo-list-thxrhmn/repositories"

	"github.com/labstack/echo/v4"
)

func SubTodoRoutes(e *echo.Group) {
	subTodoRepository := repositories.RepositorySubTodo(postgresql.DB)
	h := handlers.HandlerSubTodo(subTodoRepository)

	e.POST("/subtodo", h.CreateSubTodo)
	e.GET("/subtodo/:id", h.GetSubTodo)
	e.GET("/subtodos", h.GetSubTodos)
	e.DELETE("/subtodo/:id", h.DeleteSubTodo)
	e.PATCH("/subtodo/:id", h.UpdateSubTodo)
}
