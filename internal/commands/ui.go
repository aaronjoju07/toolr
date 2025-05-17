package commands

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/spf13/cobra"
    "github.com/aaronjoju07/toolr/internal/ui"
)

var uiCmd = &cobra.Command{
    Use:   "ui",
    Short: "Launch interactive TUI",
    Run: func(cmd *cobra.Command, args []string) {
        p := tea.NewProgram(ui.InitialModel())
        if err := p.Start(); err != nil {
            fmt.Println("Error running TUI:", err)
        }
    },
}
