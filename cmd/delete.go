package cmd

import (
    "fmt"
    "task-tracker/task"

    "github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
    Use:   "delete [id]",
    Short: "Delete a task by ID",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        id := args[0] // id should be passed as a string
        if err := task.DeleteTask(id); err != nil {
            fmt.Println("Error deleting task:", err)
            return
        }
        fmt.Println("Task deleted successfully.")
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}
