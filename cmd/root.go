package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Global flags
	outputFormat  string
	compactOutput bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adbjson",
	Short: "adbjson - ADB to JSON converter",
	Long: `adbjson is a cross-platform CLI tool that wraps Android Debug Bridge (ADB) commands
and outputs structured JSON for easy parsing and integration with other tools.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVar(&outputFormat, "format", "json", "Output format: json, yaml (default: json)")
	rootCmd.PersistentFlags().BoolVar(&compactOutput, "compact", true, "Compact JSON output (default: true)")
	rootCmd.PersistentFlags().Bool("pretty", false, "Pretty print JSON output")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug logging")
}
