# Go To-Do List Application

A simple command-line to-do list application built with Go. Manage your tasks efficiently using the following commands.

## Commands

### 1. Add a New Task

Add a task by using the `add` command followed by the task description.

```bash
go run main.go add "Task Description"
```

**Example:**

```bash
go run main.go add "Read Go book"
```

---

### 2. List All Tasks

List all tasks by using the `lists` command.

```bash
go run main.go lists
```

---

### 3. Mark a Task as Complete

Mark a task as complete by using the `complete` command followed by the task ID.

```bash
go run main.go complete TASK_ID
```

**Example:**

```bash
go run main.go complete 1
```

---

### 4. Delete a Task

Delete a task by using the `delete` command followed by the task ID.

```bash
go run main.go delete TASK_ID
```

**Example:**

```bash
go run main.go delete 1
```

---

## Example Workflow

```bash
go run main.go add "Read Go book"   # Add a new task
go run main.go lists                # List all tasks
go run main.go complete 1           # Mark task 1 as complete
go run main.go delete 1             # Delete task 1
```

---

## License

This project is licensed under the [MIT License](LICENSE).
