package cmd

import (
	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Execute shell commands",
	Long:  `Executes ADB shell commands with the same structure as original ADB.`,
}

func init() {
	rootCmd.AddCommand(shellCmd)
}
