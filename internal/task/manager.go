package task

import (
	"errors"
)

type Manager struct {
	tasks []*Task
}

// NewManager creates a new TaskManager.
func NewManager() *Manager {
	return &Manager{
		tasks: []*Task{},
	}
}

// MarkTaskAsDone marks a task as done by its ID
func (m *Manager) MarkTaskAsDone(id string) {
	task, err := m.FindTaskByID(id)
	if err != nil {
		panic(err)
	}
	task.MarkAsDone()
}

// Add adds a new task to the task manager.
func (m *Manager) Add(task *Task) {
	m.tasks = append(m.tasks, NewTask(task.Description))
}

// FindAll returns all tasks managed by the task manager.
func (m *Manager) FindAll() []*Task {
	return m.tasks
}

// FindTaskByID returns the task with the given ID.
func (m *Manager) FindTaskByID(id string) (*Task, error) {
	for _, task := range m.tasks {
		if task.ID.String() == id {
			return task, nil
		}
	}
	return nil, errors.New("task not found")
}

// findTaskIndexByID finds the index of a task with the given ID.
func (m *Manager) findTaskIndexByID(id string) (int, error) {
	for idx, task := range m.tasks {
		if task.ID.String() == id {
			return idx, nil
		}
	}
	return -1, errors.New("task not found")
}

// DeleteByID deletes the task with the given ID.
func (m *Manager) DeleteByID(id string) {
	i, err := m.findTaskIndexByID(id)
	if err != nil {
		panic(err)
	}
	m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
}
