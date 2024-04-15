package task

import (
	"flag"
	"fmt"
	"golang-task-manager/internal/cmd"
	"golang-task-manager/internal/task"
	"os"
)

func main() {
	manager := task.NewManager()

	// Define command-line flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// Parse the command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: task <command> [<args>]")
		os.Exit(1)
	}

	// Switch between commands
	switch os.Args[1] {
	case "add":
		cmd.AddCommand(addCmd, manager)
	case "list":
		cmd.AddCommand(listCmd)
	case "done":
		cmd.AddCommand(doneCmd)
	case "delete":
		cmd.AddCommand(deleteCmd)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
