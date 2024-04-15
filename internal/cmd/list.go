package cmd

import (
	"flag"
	"fmt"
	"golang-task-manager/internal/task"
)

// ListCommand represents the "list" command.
func ListCommand(fs *flag.FlagSet, manager *task.Manager) {
	tasks := manager.FindAll()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	// Print column headers
	fmt.Printf("%-5s %-15s %-10s\n", "ID", "Name", "Status")
	for _, task := range tasks {
		status := "Not Done"
		if task.Status {
			status = "Done"
		}
		fmt.Printf("%-5d %-15s %s\n", task.ID, task.Name, status)
	}
}
