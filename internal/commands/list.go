package commands

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/aaronjoju07/toolr/internal/registry"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all registered tools",
    Run: func(cmd *cobra.Command, args []string) {
        tools := registry.ListTools()
        for _, tool := range tools {
            fmt.Printf("- %s (%s): %s\n", tool.Name, tool.Path, tool.Command)
        }
    },
}
