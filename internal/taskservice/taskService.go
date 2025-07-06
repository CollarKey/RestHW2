// Package taskservice файла taskService.go содержит реализацию бизнес-логики для работы с tasks.
package taskservice

import (
	"fmt"
	"log"
)

// TaskService реализует бизнес-логику для task, используя TaskRepository для доступа к БД.
type TaskService struct {
	repo TaskRepository
}

// NewService создает экземпляр TaskService с внедренным репозиторием.
func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask реализует бизнес-логику для создания task, используя методы репозитория.
func (s *TaskService) CreateTask(task Task) (Task, error) {
	tasks, err := s.repo.CreateTask(task)
	if err != nil {
		return Task{}, fmt.Errorf("service: failed to create the task: %w", err)
	}

	return tasks, nil
}

// GetAllTasks реализует бизнес-логику для получения всех task, используя методы репозитория.
func (s *TaskService) GetAllTasks() ([]Task, error) {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("service: failed to get all tasks: %w", err)
	}

	return tasks, nil
}

// GetTaskByID реализует бизнес-логику для получения task по ID, используя методы репозитория.
func (s *TaskService) GetTaskByID(id uint) (Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, fmt.Errorf("service: failed to find the task by ID: %w", err)
	}

	return task, nil
}

// UpdateTaskByID реализует бизнес-логику для обновления task по ID, используя методы репозитория.
func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	currentTask, err := s.GetTaskByID(id)
	if err != nil {
		return Task{}, fmt.Errorf("service: task not found: %w", err)
	}

	// Обновляем только те поля, которые были переданы
	if task.Task != "" {
		currentTask.Task = task.Task
	}

	if task.IsDone != currentTask.IsDone {
		currentTask.IsDone = task.IsDone
	}

	updatedTask, err := s.repo.UpdateTaskByID(id, currentTask)
	if err != nil {
		return Task{}, fmt.Errorf("service: failed to update the task: %w", err)
	}

	return updatedTask, nil
}

// DeleteTask реализует бизнес-логику для удаления task по ID, используя методы репозитория.
func (s *TaskService) DeleteTask(id uint) error {
	log.Printf("Удаление задачи с ID %d", id)

	err := s.repo.DeleteTaskByID(id)
	if err != nil {
		return fmt.Errorf("service: failed to delete the task: %w", err)
	}

	return nil
}
