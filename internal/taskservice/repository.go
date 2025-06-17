// Package taskservice repository.go содержит реализацию репозитория для работы с задачами в базе данных.
//
//nolint:exhaustr
//nolint:exhaustruct
package taskservice

import (
	"fmt"

	"gorm.io/gorm"
)

// TaskRepository provides methods to interact with the database layer.
type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку.
	CreateTask(task Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку.
	GetAllTasks() ([]Task, error)
	// GetTaskByID - находит и возвращает задачу по её ID
	GetTaskByID(id uint) (Task, error)
	// UpdateTaskByID - передаем id и Task, возвращаем обновленный Task
	// и ошибку.
	UpdateTaskByID(id uint, task Task) (Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку.
	DeleteTaskByID(id uint) error
}

// DbTaskRepository реализует репозиторий task, предоставляет методы для взаимодействия с таблицей tasks в БД.
type DbTaskRepository struct {
	db *gorm.DB
}

// NewTaskRepository создает новый экземпляр DbTaskRepository с подключением к БД.
func NewTaskRepository(db *gorm.DB) *DbTaskRepository {
	return &DbTaskRepository{db: db}
}

// CreateTask добавляет новую task в БД и возвращает её с присвоенным ID.
func (r *DbTaskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}

	return task, nil
}

// GetAllTasks находит и возвращает все имеющиеся task из БД.
func (r *DbTaskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task

	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, fmt.Errorf("repository: failed to get all tasks: %w", err)
	}

	return tasks, nil
}

// GetTaskByID находит и возвращает task по ID.
func (r *DbTaskRepository) GetTaskByID(id uint) (Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return Task{}, fmt.Errorf("repository: failed to get task by ID: %w", err)
	}

	return task, nil
}

// UpdateTaskByID ищет task по ID, обновляет её полями из Task и возвращает обновленную task.
func (r *DbTaskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var updateTask Task
	if err := r.db.First(&updateTask, id).Error; err != nil {
		return Task{}, fmt.Errorf("repository: task not found: %w", err)
	}

	// Обновляем только непустые поля
	updates := make(map[string]interface{})
	if task.Task != "" {
		updates["task"] = task.Task
	}

	if task.IsDone != updateTask.IsDone {
		updates["is_done"] = task.IsDone
	}

	if err := r.db.Model(&updateTask).Updates(updates).Error; err != nil {
		return Task{}, fmt.Errorf("repository: failed to update task: %w", err)
	}

	return updateTask, nil
}

// DeleteTaskByID удаляет task с указанным ID.
func (r *DbTaskRepository) DeleteTaskByID(id uint) error {
	return r.db.Delete(&Task{}, id).Error
}
