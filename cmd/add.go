package cmd

import (
    "fmt"
    "task-tracker/task"  // Correct the import to task, as per the directory structure

    "github.com/spf13/cobra"
)

// Define the add command
var addCmd = &cobra.Command{
    Use:   "add [description]",
    Short: "Add a new task",
    Args:  cobra.ExactArgs(1), // Ensure that exactly one argument is passed (the task description)
    Run: func(cmd *cobra.Command, args []string) {
        // The task description is passed as the first argument
        description := args[0]

        // Call the AddTask function from task to add the task
        if err := task.AddTask(description); err != nil {
            fmt.Println("Error adding task:", err) // Handle errors
            return
        }
        fmt.Println("Task added successfully.") // If the task is added successfully
    },
}

// init function registers the addCmd to the root command
func init() {
    rootCmd.AddCommand(addCmd) // Add the 'add' command to the root command
}
