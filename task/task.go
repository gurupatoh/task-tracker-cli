package task

import "time"

// Task represents a task with its properties
type Task struct {
    ID          int       `json:"id"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}
