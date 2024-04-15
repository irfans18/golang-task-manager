package cmd

import (
	"flag"
	"fmt"
	"golang-task-manager/internal/task"
	"os"
)

// AddCommand represents the "add" command.
func AddCommand(fs *flag.FlagSet, manager *task.Manager) {
	// Parse command-line flags
	fs.Parse(os.Args[2:])

	// Get the task name from command-line arguments
	args := fs.Args()
	if len(args) < 1 {
		fmt.Println("Usage: task add <task name>")
		return
	}

	payload := task.NewTask(args[0])

	// Add the task to the task manager
	err := manager.Add(payload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Task added successfully:", payload.Name)
}
