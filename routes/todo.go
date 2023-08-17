package routes

import (
	"todo-list-thxrhmn/handlers"
	"todo-list-thxrhmn/pkg/middleware"
	"todo-list-thxrhmn/pkg/postgresql"
	"todo-list-thxrhmn/repositories"

	"github.com/labstack/echo/v4"
)

func TodoRoutes(e *echo.Group) {
	todoRepository := repositories.RepositoryTodo(postgresql.DB)
	h := handlers.HandlerTodo(todoRepository)

	e.POST("/todo", middleware.UploadFile(h.CreateTodo))
	e.GET("/todo/:id", h.GetTodo)
	e.GET("/todos", h.GetTodos)
	e.DELETE("/todo/:id", h.DeleteTodo)
	e.PATCH("/todo/:id", middleware.UploadFile(h.UpdateTodo))
}
