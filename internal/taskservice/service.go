package taskservice

import (
	"fmt"
	"log"
)

// TaskService реализует бизнес-логику для task, используя TaskRepository для доступа к БД
type TaskService struct {
	repo TaskRepository
}

// NewService создает экземпляр TaskService с внедренным репозиторием.
func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask реализует бизнес-логику для создания task, используя методы репозитория
func (s *TaskService) CreateTask(task Task) (Task, error) {
	tasks, err := s.repo.CreateTask(task)
	if err != nil {
		return Task{}, fmt.Errorf("service: failed to create the task: %w", err)
	}

	return tasks, nil
}

// GetAllTasks реализует бизнес-логику для получения всех task, используя методы репозитория
func (s *TaskService) GetAllTasks() ([]Task, error) {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("service: failed to get all tasks: %w", err)
	}

	return tasks, nil
}

// UpdateTaskByID реализует бизнес-логику для обновления task по ID, используя методы репозитория
func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	tasks, err := s.repo.UpdateTaskByID(id, task)
	if err != nil {
		return Task{}, fmt.Errorf("service: failed to updated the task: %w", err)
	}

	return tasks, nil
}

// DeleteTask реализует бизнес-логику для удаления task по ID, используя методы репозитория
func (s *TaskService) DeleteTask(id uint) error {
	log.Printf("Удаление задачи с ID %d", id)

	err := s.repo.DeleteTaskByID(id)
	if err != nil {
		return fmt.Errorf("service: failed to delete the task: %w", err)
	}

	return nil
}