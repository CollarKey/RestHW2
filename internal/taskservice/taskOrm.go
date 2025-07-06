// Package taskservice taskOrm.go содержит модели таблиц базы данных для task.
package taskservice

import "gorm.io/gorm"

// Task модель таблицы task в базе данных.
type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}
