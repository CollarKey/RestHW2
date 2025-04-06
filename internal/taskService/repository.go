package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	//CreateTask - Передлаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task RequestBody) (RequestBody, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]RequestBody, error)
	// UpdateTaskByID - передаем id и Task, возврщаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task RequestBody) (RequestBody, error)
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
func (r *taskRepository) CreateTask(task RequestBody) (RequestBody, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return RequestBody{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]RequestBody, error) {
	var tasks []RequestBody
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task RequestBody) (RequestBody, error) {
	var updateTask RequestBody
	if err := r.db.First(&updateTask, id).Error; err != nil {
		return RequestBody{}, err
	}

	err := r.db.Model(&updateTask).Updates(task)
	if err.Error != nil {
		return RequestBody{}, err.Error
	}

	return updateTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	return r.db.Delete(&RequestBody{}, id).Error
	// написал ранее result := r.db.Delete(&requestBody{}, id)
	// if result.Error != nil {
	//	return result.Error
	//}
	//return nil
}
