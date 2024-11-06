package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var tasks []Task

const filename = "tasks.json"

// Load tasks from file
func loadTasks() error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []Task{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &tasks)
}

// Save tasks to file
func saveTasks() error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Add a new task
func addTask(description string) {
	id := len(tasks) + 1
	task := Task{ID: id, Description: description, Completed: false}
	tasks = append(tasks, task)
	saveTasks()

	fmt.Println("Task added successfully.")
}

// List all tasks
func listTasks() {
	if len(tasks) == 0 {

		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Complete"
		}

		fmt.Printf("%d. %s [%s]\n", task.ID, task.Description, status)
	}
}

// Edit a task description
func editTask(id int, newDescription string) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescription
			saveTasks()

			fmt.Println("Task updated successfully.")
			return
		}
	}
	fmt.Println("Task not found.")
}

// Delete a task by ID
func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()

			fmt.Println("Task deleted successfully.")
			return
		}
	}
	fmt.Println("Task not found.")
}

// Toggle task completion status
func toggleTaskCompletion(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = !task.Completed
			saveTasks()

			fmt.Println("Task completion status updated.")
			return
		}
	}
	fmt.Println("Task not found.")
}

// Main runs a simple to-do list application in a loop, allowing the user to select
// options such as adding a task, listing tasks, editing a task, deleting a task,
// toggling a task's completion status, or exiting the application. The screen is
// cleared before displaying the menu and results.
func main() {
	loadTasks()
	for {

		fmt.Println("Simple To-Do List App")
		fmt.Println("\nTo-Do List:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Edit Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Toggle Task Completion")
		fmt.Println("6. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter task description: ")
			var description string
			fmt.Scan(&description)
			addTask(description)
		case 2:
			listTasks()
		case 3:
			fmt.Print("Enter task ID to edit: ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Enter new description: ")
			var newDescription string
			fmt.Scan(&newDescription)
			editTask(id, newDescription)
		case 4:
			fmt.Print("Enter task ID to delete: ")
			var id int
			fmt.Scan(&id)
			deleteTask(id)
		case 5:
			fmt.Print("Enter task ID to toggle completion: ")
			var id int
			fmt.Scan(&id)
			toggleTaskCompletion(id)
		case 6:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
