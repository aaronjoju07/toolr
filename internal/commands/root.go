package commands

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "toolr",
    Short: "Toolr is a helper launcher CLI for dev tools",
    Long:  `Toolr lets you register and run helper CLI tools from one place.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.AddCommand(registerCmd)
    rootCmd.AddCommand(runCmd)
    rootCmd.AddCommand(listCmd)
    rootCmd.AddCommand(uiCmd)
}
