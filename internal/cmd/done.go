package cmd

import (
	"flag"
	"fmt"
	"golang-task-manager/internal/task"
	"os"
)

// DoneCommand represents the "done" command.
func DoneCommand(fs *flag.FlagSet, manager *task.Manager) {
	// Parse command-line flags
	fs.Parse(os.Args[2:])

	// Get the task name from command-line arguments
	args := fs.Args()
	if len(args) < 1 {
		fmt.Println("Usage: task add <task name>")
		return
	}

	id := args[0]
	task, err := manager.FindTaskByID(id)
	if err != nil {
	}
	task.MarkAsDone()
}
