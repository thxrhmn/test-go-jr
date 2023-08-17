package repositories

import (
	"todo-list-thxrhmn/models"

	"gorm.io/gorm"
)

type SubTodoRepository interface {
	CreateSubTodo(subtodo models.SubTodo) (models.SubTodo, error)
	GetSubTodo(ID int) (models.SubTodo, error)
	GetSubTodos() ([]models.SubTodo, error)
	GetSubTodosByTitle(title string) ([]models.SubTodo, error)
	GetSubTodosByDescription(description string) ([]models.SubTodo, error)
	GetSubTodosByTitleAndDescription(title, description string) ([]models.SubTodo, error)
	UpdateSubTodo(subtodo models.SubTodo) (models.SubTodo, error)
	DeleteSubTodo(subtodo models.SubTodo, ID int) (models.SubTodo, error)
}

func RepositorySubTodo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateSubTodo(subtodo models.SubTodo) (models.SubTodo, error) {
	err := r.db.Create(&subtodo).Error

	return subtodo, err
}

func (r *repository) GetSubTodo(ID int) (models.SubTodo, error) {
	var subtodo models.SubTodo
	err := r.db.First(&subtodo, ID).Error

	return subtodo, err
}

func (r *repository) GetSubTodos() ([]models.SubTodo, error) {
	var subtodos []models.SubTodo
	err := r.db.Find(&subtodos).Error

	return subtodos, err
}

func (r *repository) GetSubTodosByTitle(title string) ([]models.SubTodo, error) {
	var subtodos []models.SubTodo
	err := r.db.Where("title ILIKE ?", "%"+title+"%").Find(&subtodos).Error

	return subtodos, err
}

func (r *repository) GetSubTodosByDescription(description string) ([]models.SubTodo, error) {
	var subtodos []models.SubTodo
	err := r.db.Where("description ILIKE ?", "%"+description+"%").Find(&subtodos).Error

	return subtodos, err
}

func (r *repository) GetSubTodosByTitleAndDescription(title, description string) ([]models.SubTodo, error) {
	var subtodos []models.SubTodo
	err := r.db.Where("title ILIKE ? OR description ILIKE ?", "%"+title+"%", "%"+description+"%").Find(&subtodos).Error

	return subtodos, err
}

func (r *repository) UpdateSubTodo(subtodo models.SubTodo) (models.SubTodo, error) {
	err := r.db.Save(&subtodo).Error

	return subtodo, err
}

func (r *repository) DeleteSubTodo(subtodo models.SubTodo, ID int) (models.SubTodo, error) {
	err := r.db.Delete(&subtodo).Error

	return subtodo, err
}