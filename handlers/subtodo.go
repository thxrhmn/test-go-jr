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

type handlerSubTodo struct {
	SubTodoRepository repositories.SubTodoRepository
}

func HandlerSubTodo(SubTodoRepository repositories.SubTodoRepository) *handlerSubTodo {
	return &handlerSubTodo{SubTodoRepository}
}

// CreateSubTodo godoc
// @Summary Create a subtodo
// @Description Create a new subtodo item
// @Tags subtodo
// @Param todo_id formData int true "TodoID"
// @Param title formData string true "Title"
// @Param description formData string true "Description"
// @Param file formData string false "File"
// @Accept mpfd
// @Produce json
// @Success 200 {object} dto.SuccessResult
// @Failure 500 {object} dto.ErrorResult
// @Router /subtodo [post]
func (h *handlerSubTodo) CreateSubTodo(c echo.Context) error {
	todoId, _ := strconv.Atoi(c.FormValue("todo_id"))

	request := tododto.SubTodoRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		File:        c.FormValue("file"),
		TodoID:      todoId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	subtodo := models.SubTodo{
		Title:       request.Title,
		Description: request.Description,
		File:        request.File,
		TodoID:      request.TodoID,
	}

	subtodo, err = h.SubTodoRepository.CreateSubTodo(subtodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	subtodo, _ = h.SubTodoRepository.GetSubTodo(subtodo.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: subtodo})
}

// GetSubTodo godoc
// @Summary Get a subtodo by id
// @Description Get a subtodo by id
// @Tags subtodo
// @Param id path int true "Id"
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResult
// @Failure 400 {object} dto.ErrorResult
// @Router /subtodo/{id} [get]
func (h *handlerSubTodo) GetSubTodo(c echo.Context) error {
	todoId, _ := strconv.Atoi(c.Param("id"))

	todo, err := h.SubTodoRepository.GetSubTodo(todoId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: todo})
}

// GetSubTodos godoc
// @Summary Get all subtodo
// @Description Get all subtodo
// @Tags subtodo
// @Param page query int false "Search query by pagination"
// @Param per_page query int false "Search query by pagination"
// @Param search_title query string false "Search query by title"
// @Param search_description query string false "Search query by description"
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResult
// @Failure 500 {object} dto.ErrorResult
// @Router /subtodos [get]
func (h *handlerSubTodo) GetSubTodos(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perPage, _ := strconv.Atoi(c.QueryParam("per_page"))
	searchTitle := c.QueryParam("search_title")
	searchDescription := c.QueryParam("search_description")

	var todos []models.SubTodo
	var err error

	if page != 0 && perPage != 0 {
		todos, err = h.SubTodoRepository.GetSubTodosByPagination(page, perPage) 
	} else if searchTitle != "" {
		todos, err = h.SubTodoRepository.GetSubTodosByTitle(searchTitle)
	} else if searchDescription != "" {
		todos, err = h.SubTodoRepository.GetSubTodosByDescription(searchDescription)
	} else {
		todos, err = h.SubTodoRepository.GetSubTodos()
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: todos})
}

// UpdateSubTodo godoc
// @Summary Update a subtodo
// @Description Update a subtodo
// @Tags subtodo
// @Param id path int true "Id"
// @Param title formData string true "Title"
// @Param description formData string true "Description"
// @Param file formData string false "File"
// @Accept mpfd
// @Produce json
// @Success 200 {object} dto.SuccessResult
// @Failure 500 {object} dto.ErrorResult
// @Router /subtodo/{id} [patch]
func (h *handlerSubTodo) UpdateSubTodo(c echo.Context) error {
	todoId, _ := strconv.Atoi(c.Param("id"))

	request := tododto.SubTodoRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		File:        c.FormValue("file"),
		TodoID:      todoId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	subTodo, _ := h.SubTodoRepository.GetSubTodo(todoId)

	if request.Title != "" {
		subTodo.Title = request.Title
	}

	if request.Description != "" {
		subTodo.Description = request.Description
	}

	if request.File != "" {
		subTodo.File = request.File
	}

	if request.TodoID != 0 {
		subTodo.TodoID = request.TodoID
	}

	data, _ := h.SubTodoRepository.UpdateSubTodo(subTodo)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

// DeleteSubTodo godoc
// @Summary Delete a subtodo
// @Description Delete a subtodo
// @Tags subtodo
// @Param id path int true "Id"
// @Accept mpfd
// @Produce json
// @Success 200 {object} dto.SuccessResult
// @Failure 500 {object} dto.ErrorResult
// @Router /subtodo/{id} [delete]
func (h *handlerSubTodo) DeleteSubTodo(c echo.Context) error {
	subTodoId, _ := strconv.Atoi(c.Param("id"))
	// subtodoId, _ := strconv.Atoi(c.Param("subtodo_id"))

	subTodo, err := h.SubTodoRepository.GetSubTodo(subTodoId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.SubTodoRepository.DeleteSubTodo(subTodo, subTodoId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}
