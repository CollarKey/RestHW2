package taskService

import "log"

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Tasks) (Tasks, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Tasks, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, task Tasks) (Tasks, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTask(id uint) error {
	log.Printf("Удаление задачи с ID %d", id)
	return s.repo.DeleteTaskByID(id)
}
