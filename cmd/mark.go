package cmd

import (
    "fmt"
    "task-tracker/task"

    "github.com/spf13/cobra"
)

var markCmd = &cobra.Command{
    Use:   "mark [status] [id]",
    Short: "Mark a task as in-progress or done",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        status := args[0]  // Status ("in-progress" or "done")
        id := args[1]      // ID as a string
        if err := task.MarkTask(id, status); err != nil {
            fmt.Println("Error marking task:", err)
            return
        }
        fmt.Println("Task marked successfully.")
    },
}

func init() {
    rootCmd.AddCommand(markCmd)
}
