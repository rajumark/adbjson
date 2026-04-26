package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// serialnoCmd represents the serialno command
var serialnoCmd = &cobra.Command{
	Use:   "get-serialno",
	Short: "Get device serial number in JSON format",
	Long:  `Executes "adb get-serialno" and outputs the result as structured JSON.`,
	RunE:  runSerialNo,
}

func init() {
	rootCmd.AddCommand(serialnoCmd)
}

func runSerialNo(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb get-serialno
	output, err := executor.Execute("get-serialno")
	if err != nil {
		return fmt.Errorf("failed to execute adb get-serialno: %w", err)
	}
	
	// Parse output
	parser := parser.NewSerialNoParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse serialno output: %w", err)
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
