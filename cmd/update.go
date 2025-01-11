package cmd

import (
    "fmt"
    "task-tracker/task"

    "github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
    Use:   "update [id] [description]",
    Short: "Update the description of a task",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        id := args[0]           // ID as a string
        description := args[1]  // New description
        if err := task.UpdateTask(id, description); err != nil {
            fmt.Println("Error updating task:", err)
            return
        }
        fmt.Println("Task updated successfully.")
    },
}

func init() {
    rootCmd.AddCommand(updateCmd)
}
