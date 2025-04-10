package todo

import (
	"cli-todo/models"
	"fmt"
)

var tasks []models.Task
var nextID int

func init() {
	loaded, err := LoadTasks()
	if err != nil {
		fmt.Println("Could not load tasks:", err)
	}
	tasks = loaded

	for _, t := range tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
}

func AddTask(title string) {
	task := models.Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}
	nextID++
	tasks = append(tasks, task)
	SaveTasks(tasks)
	fmt.Println("Task added:", task.Title)
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet.")
		return
	}
	for _, t := range tasks {
		status := "❌"
		if t.Completed {
			status = "✅"
		}
		fmt.Printf("[%d] %s %s\n", t.ID, status, t.Title)
	}
}

func CompleteTask(id int) {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			SaveTasks(tasks)
			fmt.Println("Task marked as completed:", t.Title)
			return
		}
	}
	fmt.Println("Task not found.")
}

func DeleteTask(id int) {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			SaveTasks(tasks)
			fmt.Println("Task deleted:", t.Title)
			return
		}
	}
	fmt.Println("Task not found.")
}
