package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

// rootCmd is the root command of your CLI
var rootCmd = &cobra.Command{
    Use:   "task-cli",
    Short: "A task management CLI",
    Long:  "A simple CLI for managing tasks and keeping track of to-dos.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to the Task Tracker CLI! Use 'task-cli add' to add a task.")
    },
}

// Execute is the entry point for the CLI
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println("Error executing command:", err)
        os.Exit(1)
    }
}
