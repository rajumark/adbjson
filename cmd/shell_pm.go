package cmd

import (
	"github.com/spf13/cobra"
)

// pmCmd represents the pm command under shell
var pmCmd = &cobra.Command{
	Use:   "pm",
	Short: "Package Manager",
	Long:  `Package Manager commands - same as adb shell pm.`,
}

func init() {
	shellCmd.AddCommand(pmCmd)
}
