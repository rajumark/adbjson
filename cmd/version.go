package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "adb-version",
	Short: "Get ADB version information in JSON format",
	Long:  `Executes "adb version" and outputs the result as structured JSON.`,
	RunE:  runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb version
	output, err := executor.Execute("version")
	if err != nil {
		return fmt.Errorf("failed to execute adb version: %w", err)
	}
	
	// Parse output
	parser := parser.NewVersionParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse version output: %w", err)
	}
	
	// Determine output format
	var jsonBytes []byte
	if compactOutput {
		jsonBytes, err = json.Marshal(response)
	} else {
		jsonBytes, err = json.MarshalIndent(response, "", "  ")
	}
	
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	
	// Print to stdout
	fmt.Println(string(jsonBytes))
	
	return nil
}
