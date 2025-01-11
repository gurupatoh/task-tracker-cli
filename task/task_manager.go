package task

import (
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "time"
)

const taskFile = "./data/tasks.json"

// readTasks reads and parses the tasks file into a slice of Task objects
func readTasks() ([]Task, error) {
    file, err := os.Open(taskFile)
    if err != nil {
        if os.IsNotExist(err) {
            return []Task{}, nil // Return an empty slice if the file does not exist
        }
        return nil, fmt.Errorf("error opening tasks file: %v", err)
    }
    defer file.Close()

    var tasks []Task
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&tasks)
    if err != nil {
        return nil, fmt.Errorf("error decoding tasks: %v", err)
    }
    return tasks, nil
}

// writeTasks writes a slice of Task objects back to the tasks file
func writeTasks(tasks []Task) error {
    file, err := os.Create(taskFile)
    if err != nil {
        return fmt.Errorf("error creating tasks file: %v", err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ") // Pretty print the JSON
    err = encoder.Encode(tasks)
    if err != nil {
        return fmt.Errorf("error encoding tasks to file: %v", err)
    }
    return nil
}

// AddTask adds a new task to the tasks file
func AddTask(description string) error {
    tasks, err := readTasks()
    if err != nil {
        return err
    }

    newID := 1
    if len(tasks) > 0 {
        newID = tasks[len(tasks)-1].ID + 1 // Assign a new ID
    }

    task := Task{
        ID:          newID,
        Description: description,
        Status:      "todo",
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    tasks = append(tasks, task)
    return writeTasks(tasks)
}

// UpdateTask updates the description of an existing task
func UpdateTask(id string, description string) error {
    tasks, err := readTasks()
    if err != nil {
        return err
    }

    taskID, err := strconv.Atoi(id)
    if err != nil {
        return fmt.Errorf("invalid task ID")
    }

    for i, t := range tasks {
        if t.ID == taskID {
            tasks[i].Description = description
            tasks[i].UpdatedAt = time.Now()
            return writeTasks(tasks)
        }
    }
    return fmt.Errorf("task with ID %d not found", taskID)
}

// DeleteTask deletes a task from the tasks list
func DeleteTask(id string) error {
    tasks, err := readTasks()
    if err != nil {
        return err
    }

    taskID, err := strconv.Atoi(id)
    if err != nil {
        return fmt.Errorf("invalid task ID")
    }

    for i, t := range tasks {
        if t.ID == taskID {
            tasks = append(tasks[:i], tasks[i+1:]...) // Remove task from slice
            return writeTasks(tasks)
        }
    }
    return fmt.Errorf("task with ID %d not found", taskID)
}

// MarkTask marks a task as either in-progress or done
func MarkTask(id string, status string) error {
    tasks, err := readTasks()
    if err != nil {
        return err
    }

    taskID, err := strconv.Atoi(id)
    if err != nil {
        return fmt.Errorf("invalid task ID")
    }

    for i, t := range tasks {
        if t.ID == taskID {
            if status != "in-progress" && status != "done" {
                return fmt.Errorf("invalid status. Valid values are 'in-progress' or 'done'")
            }
            tasks[i].Status = status
            tasks[i].UpdatedAt = time.Now()
            return writeTasks(tasks)
        }
    }
    return fmt.Errorf("task with ID %d not found", taskID)
}

// ListTasks lists tasks based on their status (all, todo, in-progress, done)
func ListTasks(status string) ([]Task, error) {
    tasks, err := readTasks()
    if err != nil {
        return nil, err
    }

    var filteredTasks []Task
    for _, t := range tasks {
        if status == "all" || t.Status == status {
            filteredTasks = append(filteredTasks, t)
        }
    }
    return filteredTasks, nil
}
