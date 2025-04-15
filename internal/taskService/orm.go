package taskService

type Tasks struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
