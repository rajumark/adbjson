package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	prettyOutput   bool
	compactOutput  bool
	version        = "1.0.0"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adbjson",
	Short: "A CLI tool that wraps ADB commands and outputs structured JSON",
	Long: `adbjson is a cross-platform CLI tool that wraps Android Debug Bridge (ADB) commands
and outputs structured JSON for easy parsing and integration with other tools.`,
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&prettyOutput, "pretty", true, "Pretty print JSON output (default: true)")
	rootCmd.PersistentFlags().BoolVar(&compactOutput, "compact", false, "Compact JSON output (overrides --pretty)")
}
