package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// devpathCmd represents the devpath command
var devpathCmd = &cobra.Command{
	Use:   "get-devpath",
	Short: "Get device path in JSON format",
	Long:  `Executes "adb get-devpath" and outputs the result as structured JSON.`,
	RunE:  runDevPath,
}

func init() {
	rootCmd.AddCommand(devpathCmd)
}

func runDevPath(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb get-devpath
	output, err := executor.Execute("get-devpath")
	if err != nil {
		return fmt.Errorf("failed to execute adb get-devpath: %w", err)
	}
	
	// Parse output
	parser := parser.NewDevPathParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse devpath output: %w", err)
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
