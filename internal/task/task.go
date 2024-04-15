package task

import (
	"github.com/google/uuid"
)

// Task represents a task in the task manager.
type Task struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Status bool
}

// NewTask creates a new Task with the given ID and description.
func NewTask(description string) *Task {
	return &Task{
		ID:     uuid.New(),
		Name:   description,
		Status: false,
	}
}

// MarkAsDone marks the task as done.
func (t *Task) MarkAsDone() {
	t.Status = true
}
