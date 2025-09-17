// Package userservice userOrm.go содержит модели таблиц базы данных для User.
package userservice

import (
	"CheckingErrorsHW2/internal/taskservice"

	"gorm.io/gorm"
)

// User модель таблицы user в базе данных.
type User struct {
	gorm.Model
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskservice.Task `json:"tasks,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdatre:CASCADE,OnDelete:CASCADE;"`
}
