package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// reconnectCmd represents the reconnect command
var reconnectCmd = &cobra.Command{
	Use:   "reconnect",
	Short: "Reconnect device in JSON format",
	Long:  `Executes "adb reconnect" and outputs the result as structured JSON.`,
	RunE:  runReconnect,
}

func init() {
	rootCmd.AddCommand(reconnectCmd)
}

func runReconnect(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb reconnect
	output, err := executor.Execute("reconnect")
	if err != nil {
		return fmt.Errorf("failed to execute adb reconnect: %w", err)
	}
	
	// Parse output
	parser := parser.NewReconnectParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse reconnect output: %w", err)
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
