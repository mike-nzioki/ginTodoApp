package repositories

import (
	"gorm.io/gorm"
	"todo/models"
)

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	Create(todo models.Todo) (models.Todo, error)
	GetByID(id uint) (models.Todo, error)
	Update(todo models.Todo) (models.Todo, error)
	Delete(id uint) error
}

type todoGormRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoGormRepository{db: db}
}

func (r *todoGormRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *todoGormRepository) Create(todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&todo).Error
	return todo, err
}

func (r *todoGormRepository) GetByID(id uint) (models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	return todo, err
}

func (r *todoGormRepository) Update(todo models.Todo) (models.Todo, error) {
	err := r.db.Save(&todo).Error
	return todo, err
}

func (r *todoGormRepository) Delete(id uint) error {
	err := r.db.Delete(&models.Todo{}, id).Error
	return err
}
