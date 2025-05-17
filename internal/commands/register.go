package commands

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/aaronjoju07/toolr/internal/registry"
)

var registerCmd = &cobra.Command{
    Use:   "register [name] [path] [command]",
    Short: "Register a tool",
    Args:  cobra.ExactArgs(3),
    Run: func(cmd *cobra.Command, args []string) {
        name, path, command := args[0], args[1], args[2]
        err := registry.RegisterTool(name, path, command)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Println("Tool registered successfully.")
        }
    },
}
