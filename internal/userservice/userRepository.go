// Package userservice userRepository.go содержит реализацию репозитория для работы с users в базе данных.
//nolint:exhaustruct
package userservice

import (
	"fmt"

	"gorm.io/gorm"
)

// UserRepository определяет методы взаимодействия с БД.
type UserRepository interface {
	CreateUser(usr User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (User, error)
	UpdateUser(id uint, user User) (User, error)
	DeleteUser(id uint) error
}

// DBUserRepository реализует репозиторий user, предоставляет методы для взаимодействия с БД.
type DBUserRepository struct {
	db *gorm.DB
}

// NewUserRepository конструкт создания экземпляра репозитория для работы с User.
func NewUserRepository(db *gorm.DB) *DBUserRepository {
	return &DBUserRepository{db:db}
}

// CreateUser добавляет нового user в БД, возвращает его с присвоенным ID.
func (r *DBUserRepository) CreateUser(user User) (User, error)  {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

// GetAllUsers возвращает слайс всех User из БД.
func (r *DBUserRepository) GetAllUsers() ([]User, error) {
	var Users []User

	err := r.db.Find(&Users).Error
	if err != nil {
		return nil, fmt.Errorf("repository: failed to get users %w", err)
	}

	return Users, nil
}

// GetUserByID возвращает User по его ID.
func (r *DBUserRepository) GetUserByID(id uint) (User, error) {
	var user User

	err := r.db.First(&user, id).Error
	if err != nil{
		return User{}, fmt.Errorf("repository: failed to find user by id: %w", err)
	}

	return user, nil
}

// UpdateUser получает user при помощи метода GetUserByID, обновляет поля User, возвращает обновленного user.
func (r *DBUserRepository) UpdateUser(id uint, user User) (User, error) {
	userToUpdate, err := r.GetUserByID(id)
	if err != nil {
		return User{}, err
	}

	updates := make(map[string]interface{})
	if user.Email != "" {
		updates["Email"] = user.Email
	}

	if user.Password != "" {
		updates["Password"] = user.Password
	}

	if err := r.db.Model(&userToUpdate).Updates(updates).Error; err != nil {
		return User{}, fmt.Errorf("repository: failed to update user: %w", err)
	}

	if err := r.db.First(&userToUpdate, id).Error; err != nil {
		return User{}, fmt.Errorf("repository: failed to reload user after update: %w", err)
	}

	return userToUpdate, nil
}

// DeleteUser удаляет User из БД.
func (r *DBUserRepository) DeleteUser(id uint) error {
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return fmt.Errorf("repository: failed to delete user by id: %w", err)
	}

	return nil
}