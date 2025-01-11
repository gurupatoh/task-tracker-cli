
# Task Tracker CLI

**Task Tracker CLI** is a command-line application built in Go for managing tasks efficiently. It allows you to add, update, delete, and list tasks, as well as mark tasks as in progress or completed. Tasks are stored persistently in a JSON file, ensuring that your task list is available even after restarting the application.

---

## Features

- **Add Tasks**: Add new tasks to your to-do list.
- **Update Tasks**: Update the description of existing tasks.
- **Delete Tasks**: Remove tasks from the list.
- **List Tasks**: View all tasks or filter tasks by status (e.g., `todo`, `in-progress`, `done`).
- **Mark Tasks**: Mark tasks as `in-progress` or `done`.
- **Persistent Storage**: Tasks are stored in a `tasks.json` file, ensuring data persistence.

---

## Getting Started

### Prerequisites

- **Go (>=1.16)**: Ensure you have Go installed. Download it [here](https://golang.org/dl/).
- A terminal or command prompt.

---

### Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/task-tracker-cli.git
   cd task-tracker-cli
   ```

2. Build the application:
   ```bash
   go build -o task-tracker
   ```

3. Run the application:
   ```bash
   ./task-tracker
   ```

---

## Usage

Run the CLI tool by invoking the built binary followed by the desired command.

### Commands

#### Add a Task
```bash
./task-tracker add "Buy groceries"
```
Output: 
```
Task added successfully.
```

#### Update a Task
```bash
./task-tracker update <task_id> "Buy groceries and cook dinner"
```
Output: 
```
Task updated successfully.
```

#### Delete a Task
```bash
./task-tracker delete <task_id>
```
Output:
```
Task deleted successfully.
```

#### Mark a Task
```bash
# Mark a task as in-progress
./task-tracker mark in-progress <task_id>

# Mark a task as done
./task-tracker mark done <task_id>
```
Output:
```
Task marked as in-progress/done successfully.
```

#### List Tasks
```bash
# List all tasks
./task-tracker list

# List tasks by status
./task-tracker list todo
./task-tracker list in-progress
./task-tracker list done
```

---

## File Structure

```plaintext
task-tracker/
├── cmd/                # Cobra CLI commands
│   ├── add.go          # Add task functionality
│   ├── delete.go       # Delete task functionality
│   ├── list.go         # List tasks functionality
│   ├── mark.go         # Mark tasks functionality
│   ├── update.go       # Update task functionality
│   └── root.go         # Root command configuration
├── data/               # Data storage folder
│   └── tasks.json      # JSON file to persist tasks
├── task/               # Task management logic
│   ├── task.go         # Task struct and helper functions
│   ├── task_manager.go # Core task operations (CRUD)
├── main.go             # Main application entry point
├── go.mod              # Go module file
└── go.sum              # Dependencies
```

---

## Task Properties

Each task includes the following properties:

| Property   | Description                                 |
|------------|---------------------------------------------|
| `id`       | Unique identifier for the task             |
| `description` | Short description of the task          |
| `status`   | Status of the task (`todo`, `in-progress`, `done`) |
| `createdAt`| Timestamp when the task was created        |
| `updatedAt`| Timestamp when the task was last updated   |

---

## Error Handling

Errors are gracefully handled to ensure the application doesn't crash unexpectedly. Common errors (e.g., invalid input, missing files) provide descriptive messages to the user.

---

## Future Enhancements

- Add support for due dates and reminders.
- Implement sorting tasks by creation date or priority.
- Allow users to prioritize tasks.
- Introduce categories or labels for tasks.

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for improvements or bug fixes.

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your message here"
   ```
4. Push to your branch:
   ```bash
   git push origin feature/your-feature-name
   ```
5. Open a pull request.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## Acknowledgments

- Built with Go and the Cobra library for CLI functionality.
- Inspired by the simplicity of command-line productivity tools.

---

Feel free to reach out with any questions or feedback! 🚀