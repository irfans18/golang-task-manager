package task

import (
	"github.com/google/uuid"
)

// Task represents a task in the task manager.
type Task struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	status      bool
}

// NewTask creates a new Task with the given ID and description.
func NewTask(description string) *Task {
	return &Task{
		ID:          uuid.New(),
		Description: description,
		status:      false,
	}
}

// MarkAsDone marks the task as done.
func (t *Task) MarkAsDone() {
	t.status = true
}
