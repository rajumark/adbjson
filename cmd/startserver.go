package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// startserverCmd represents the start-server command
var startserverCmd = &cobra.Command{
	Use:   "start-server",
	Short: "Start ADB server in JSON format",
	Long:  `Executes "adb start-server" and outputs the result as structured JSON.`,
	RunE:  runStartServer,
}

func init() {
	rootCmd.AddCommand(startserverCmd)
}

func runStartServer(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb start-server
	output, err := executor.Execute("start-server")
	if err != nil {
		return fmt.Errorf("failed to execute adb start-server: %w", err)
	}
	
	// Parse output
	parser := parser.NewStartServerParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse start-server output: %w", err)
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
