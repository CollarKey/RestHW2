package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	//CreateTask - Передлаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Tasks) (Tasks, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Tasks, error)
	// UpdateTaskByID - передаем id и Task, возврщаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task Tasks) (Tasks, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию
func (r *taskRepository) CreateTask(task Tasks) (Tasks, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Tasks{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Tasks, error) {
	var tasks []Tasks
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Tasks) (Tasks, error) {
	var updateTask Tasks
	if err := r.db.First(&updateTask, id).Error; err != nil {
		return Tasks{}, err
	}

	err := r.db.Model(&updateTask).Updates(task)
	if err.Error != nil {
		return Tasks{}, err.Error
	}

	return updateTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	return r.db.Delete(&Tasks{}, id).Error
	// написал ранее result := r.db.Delete(&requestBody{}, id)
	// if result.Error != nil {
	//	return result.Error
	//}
	//return nil
}
