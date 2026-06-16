// Package models holds the domain types for TaskFlow.
package models

import "time"

// Priority levels a task can have. Stored as text for readability.
const (
	PriorityLow    = "low"
	PriorityMedium = "medium"
	PriorityHigh   = "high"
)

// IsValidPriority reports whether p is one of the allowed priorities.
func IsValidPriority(p string) bool {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh:
		return true
	}
	return false
}

// Task is the core domain entity.
// JSON tags define the API's wire format (Session 15).
type Task struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	Priority  string    `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateTaskInput is the payload accepted when creating a task.
// Keeping it separate from Task means clients can't set ID or CreatedAt.
type CreateTaskInput struct {
	Title    string `json:"title"`
	Priority string `json:"priority"` // optional; defaults to "medium"
}

// UpdateTaskInput is the payload accepted when updating a task.
type UpdateTaskInput struct {
	Title    string `json:"title"`
	Done     bool   `json:"done"`
	Priority string `json:"priority"` // optional; defaults to "medium"
}

// TaskFilter narrows a task listing. A nil/empty field means "don't filter".
type TaskFilter struct {
	Done     *bool  // ?done=true / ?done=false
	Priority string // ?priority=high  ("" = any)
}
