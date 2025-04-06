package taskService

import "gorm.io/gorm"

type RequestBody struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
