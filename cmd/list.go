package cmd

import (
    "fmt"
    "task-tracker/task"

    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list [status]",
    Short: "List tasks by status (all, todo, in-progress, done)",
    Args:  cobra.MaximumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        status := ""
        if len(args) == 1 {
            status = args[0]
            if status != "all" && status != "todo" && status != "in-progress" && status != "done" {
                fmt.Println("Invalid status. Use 'all', 'todo', 'in-progress', or 'done'.")
                return
            }
        }
        tasks, err := task.ListTasks(status)
        if err != nil {
            fmt.Println("Error listing tasks:", err)
            return
        }
        if len(tasks) == 0 {
            fmt.Println("No tasks found.")
            return
        }
        for _, t := range tasks {
            fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n",
                t.ID, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04:05"), t.UpdatedAt.Format("2006-01-02 15:04:05"))
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
