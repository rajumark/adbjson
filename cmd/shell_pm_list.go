package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command under pm
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items",
	Long:  `List packages, features, libraries, instrumentation, permissions.`,
}

func init() {
	pmCmd.AddCommand(listCmd)
}
