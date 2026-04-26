package cmd

import (
	"github.com/spf13/cobra"
)

// inputCmd represents the input command
var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Input related commands",
	Long:  `Input related commands.`,
}

func init() {
	shellCmd.AddCommand(inputCmd)
}
