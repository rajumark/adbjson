package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// stateCmd represents the state command
var stateCmd = &cobra.Command{
	Use:   "get-state",
	Short: "Get device state in JSON format",
	Long:  `Executes "adb get-state" and outputs the result as structured JSON.`,
	RunE:  runState,
}

func init() {
	rootCmd.AddCommand(stateCmd)
}

func runState(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb get-state
	output, err := executor.Execute("get-state")
	if err != nil {
		return fmt.Errorf("failed to execute adb get-state: %w", err)
	}
	
	// Parse output
	parser := parser.NewStateParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse state output: %w", err)
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
