package repositories

import (
	"todo-list-thxrhmn/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(todo models.Todo) (models.Todo, error)
	GetTodo(ID int) (models.Todo, error)
	GetTodos() ([]models.Todo, error)
	GetTodosByTitle(title string) ([]models.Todo, error)
	GetTodosByDescription(description string) ([]models.Todo, error)
	GetTodosByTitleAndDescription(title, description string) ([]models.Todo, error)
	UpdateTodo(todo models.Todo) (models.Todo, error)
	DeleteTodo(todo models.Todo, ID int) (models.Todo, error)
}

func RepositoryTodo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&todo).Error

	return todo, err
}

func (r *repository) GetTodo(ID int) (models.Todo, error) {
	var todo models.Todo
	err := r.db.Preload("SubTodos").First(&todo, ID).Error

	return todo, err
}

func (r *repository) GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Preload("SubTodos").Find(&todos).Error

	return todos, err
}

func (r *repository) GetTodosByTitle(title string) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Preload("SubTodos").Where("title ILIKE ?", "%"+title+"%").Find(&todos).Error

	return todos, err
}

func (r *repository) GetTodosByDescription(description string) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Preload("SubTodos").Where("description ILIKE ?", "%"+description+"%").Find(&todos).Error

	return todos, err
}

func (r *repository) GetTodosByTitleAndDescription(title, description string) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Preload("SubTodos").Where("title ILIKE ? OR description ILIKE ?", "%"+title+"%", "%"+description+"%").Find(&todos).Error

	return todos, err
}

func (r *repository) UpdateTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Save(&todo).Error

	return todo, err
}

func (r *repository) DeleteTodo(todo models.Todo, ID int) (models.Todo, error) {
	err := r.db.Delete(&todo).Error

	return todo, err
}