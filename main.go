package main

import (
	"fmt"
	"cli-todo/todo"
	"os"
	"strconv"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  cli-todo add <task>       	Add a new task")
	fmt.Println("  cli-todo lists            	List all tasks")
	fmt.Println("  cli-todo complete <id>   	Mark a task as completed")
	fmt.Println("  cli-todo delete <id>     	Delete a task")
}

func main() {
	args := os.Args

	// $ go run main.go add "Read Go"
	// ↓		  0			1		2
	// args = ["main.go", "add", "Read Go"]
	// ↓
	// args[1] = "add"
	// ↓
	// Call todo.AddTask("Read Go")


	if len(args) < 2 {
		printUsage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) < 3 {
			fmt.Println("Please provide a task title.")
			return
		}
		todo.AddTask(args[2])
	case "lists":
		todo.ListTasks()
	case "complete":
		if len(args) < 3 {
			fmt.Println("Please provide a task ID.")
			return
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}
		todo.CompleteTask(id)
	case "delete":
		if len(args) < 3 {
			fmt.Println("Please provide a task ID.")
			return
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}
		todo.DeleteTask(id)
	default:
		fmt.Println("Unknown command:", args[1])
		printUsage()

	}

}