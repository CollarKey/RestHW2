package userservice

import (
	"fmt"
	"log"
)

// UserService реализует бизнес-логику работы с user через репозиторий.
type UserService struct {
	repo UserRepository
}

// NewUserService создает новый экземпляр UserService для работы с пользователями.
// Принимает: интерфейс UserRepository в качестве зависимости для доступа к данным.
// Возвращает: экземпляр UserService, готовый к работе.
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser реализует бизнес-логику на создание нового User, используя методы репозитория.
func (s *UserService) CreateUser(user User) (User, error) {
	users, err := s.repo.CreateUser(user)
	if err != nil{
		return User{}, fmt.Errorf("service: failed to create user: %w", err)
	}

	return users, nil
}

// GetAllUsers реализует бизнес-логику на получение всех User, используя методы репозитория.
func (s *UserService) GetAllUsers() ([]User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("service: failed to get all users: %w", err)
	}

	return users, nil
}

// GetUsersByID реализует бизнес-логику на получение User по его ID, используя методы репозитория.
func (s *UserService) GetUsersByID(id uint) (User, error) {
	userByID, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, fmt.Errorf("service: failed to get the user by ID: %w", err)
	}

	return userByID, nil
}

// UpdateUserByID реализует бизнес-логику на обновление полей User по его ID, используя методы репозитория
func (s *UserService) UpdateUserByID(id uint, user User) (User, error) {
	userToUpdate, err := s.GetUsersByID(id)
	if err != nil {
		return User{}, fmt.Errorf("service: failed to get the user to update: %w", err)
	}

	if user.Email != "" {
		userToUpdate.Email = user.Email
	}

	if user.Password != "" {
		userToUpdate.Password = user.Password
	}

	updatedUser, err := s.repo.UpdateUser(id, userToUpdate)
	if err != nil {
		return User{}, fmt.Errorf("service: failed to update user: %w", err)
	}

	return updatedUser, nil
}

// DeleteUser реализует бизнес-логику на удаление User по его ID, используя методы репозитория.
func (s *UserService) DeleteUser(id uint) error {
	log.Printf("Удаление User с ID: %d", id)

	err := s.repo.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("service: failed to delete User %w", err)
	}

	return nil
}