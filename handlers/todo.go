package handlers

import (
	"net/http"
	"strconv"
	dto "todo-list-thxrhmn/dto/result"
	tododto "todo-list-thxrhmn/dto/todo"
	"todo-list-thxrhmn/models"
	"todo-list-thxrhmn/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerTodo struct {
	TodoRepository repositories.TodoRepository
}

func HandlerTodo(TodoRepository repositories.TodoRepository) *handlerTodo {
	return &handlerTodo{TodoRepository}
}

// CreateTodo godoc
// @Summary Create a todo
// @Description Create a new todo item
// @Tags todo
// @Param title formData string true "Title"
// @Param description formData string true "Description"
// @Param file formData file false "File" format(mime)
// @Accept mpfd
// @Produce json
// @Success 200 {object} dto.SuccessResult
// @Failure 500 {object} dto.ErrorResult
// @Router /todo [post]
func (h *handlerTodo) CreateTodo(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	request := tododto.TodoRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		File:        dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	todo := models.Todo{
		Title:       request.Title,
		Description: request.Description,
		File:        request.File,
	}

	todo, err = h.TodoRepository.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	todo, _ = h.TodoRepository.GetTodo(todo.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: todo})
}

// GetTodo godoc
// @Summary Get a todo by id
// @Description Get a todo by id
// @Tags todo
// @Param id path int true "Id"
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResult
// @Failure 400 {object} dto.ErrorResult
// @Router /todo/{id} [get]
func (h *handlerTodo) GetTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	todo, err := h.TodoRepository.GetTodo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: todo})
}

// GetTodos godoc
// @Summary Get all todo
// @Description Get all todo
// @Tags todo
// @Param page query int false "Search query by pagination"
// @Param per_page query int false "Search query by pagination"
// @Param search_title query string false "Search query by title"
// @Param search_description query string false "Search query by description"
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResult
// @Failure 500 {object} dto.ErrorResult
// @Router /todos [get]
func (h *handlerTodo) GetTodos(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perPage, _ := strconv.Atoi(c.QueryParam("per_page"))
	searchTitle := c.QueryParam("search_title")
	searchDescription := c.QueryParam("search_description")

	var todos []models.Todo
	var err error

	if page != 0 && perPage != 0 {
		todos, err = h.TodoRepository.GetTodosByPagination(page, perPage) 
	} else if searchTitle != "" {
		todos, err = h.TodoRepository.GetTodosByTitle(searchTitle)
	} else if searchDescription != "" {
		todos, err = h.TodoRepository.GetTodosByDescription(searchDescription)
	} else {
		todos, err = h.TodoRepository.GetTodos()
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: todos})
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update a todo
// @Tags todo
// @Param id path int true "Id"
// @Param title formData string true "Title"
// @Param description formData string true "Description"
// @Param file formData file false "File" format(mime)
// @Accept mpfd
// @Produce json
// @Success 200 {object} dto.SuccessResult
// @Failure 500 {object} dto.ErrorResult
// @Router /todo/{id} [patch]
func (h *handlerTodo) UpdateTodo(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	id, _ := strconv.Atoi(c.Param("id"))

	request := tododto.TodoRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		File:        dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	todo, _ := h.TodoRepository.GetTodo(id)

	if request.Title != "" {
		todo.Title = request.Title
	}

	if request.Description != "" {
		todo.Description = request.Description
	}

	if request.File != "" {
		todo.File = request.File
	}

	data, _ := h.TodoRepository.UpdateTodo(todo)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo
// @Tags todo
// @Param id path int true "Id"
// @Accept mpfd
// @Produce json
// @Success 200 {object} dto.SuccessResult
// @Failure 500 {object} dto.ErrorResult
// @Router /todo/{id} [delete]
func (h *handlerTodo) DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	todo, err := h.TodoRepository.GetTodo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TodoRepository.DeleteTodo(todo, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}
