package main

import "gorm.io/gorm"

type requestBody struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
