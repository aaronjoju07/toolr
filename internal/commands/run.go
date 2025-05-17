package commands

import (
    "fmt"
    "os/exec"
    "github.com/spf13/cobra"
    "github.com/aaronjoju07/toolr/internal/registry"
)

var runCmd = &cobra.Command{
    Use:   "run [name]",
    Short: "Run a registered tool",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        tool, err := registry.GetTool(args[0])
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        execCmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s && %s", tool.Path, tool.Command))
        execCmd.Stdout = cmd.OutOrStdout()
        execCmd.Stderr = cmd.OutOrStderr()
        execCmd.Run()
    },
}
