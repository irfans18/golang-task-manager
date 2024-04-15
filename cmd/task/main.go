package main

import (
	"flag"
	"fmt"
	"golang-task-manager/internal/cmd"
	"golang-task-manager/internal/task"
	"os"
)

func main() {
	manager := task.NewManager()
	manager.Add(task.NewTask("Dummy 1"))
	manager.Add(task.NewTask("Dummy 2"))
	manager.Add(task.NewTaskWithID("123", "Dummy 3"))

	// Define command-line flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	//deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

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
		cmd.ListCommand(listCmd, manager)
	case "done":
		cmd.DoneCommand(doneCmd, manager)
		cmd.ListCommand(listCmd, manager)
	//case "delete":
	//	cmd.AddCommand(deleteCmd)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
