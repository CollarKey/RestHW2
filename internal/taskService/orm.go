package taskService

import "gorm.io/gorm"

type Tasks struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
