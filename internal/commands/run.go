package commands

import (
	"fmt"
	"github.com/aaronjoju07/toolr/internal/registry"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

var runCmd = &cobra.Command{
	Use:   "run [name]",
	Short: "Run a registered tool",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		extraArgs := args[1:]

		tool, err := registry.GetTool(name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Combine the registered command with extra CLI arguments
		fullCmd := fmt.Sprintf("cd %s && %s %s", tool.Path, tool.Command, strings.Join(extraArgs, " "))

		execCmd := exec.Command("sh", "-c", fullCmd)
		execCmd.Stdout = cmd.OutOrStdout()
		execCmd.Stderr = cmd.OutOrStderr()
		execCmd.Run()
	},
}
