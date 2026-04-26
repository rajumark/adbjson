package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// killserverCmd represents the kill-server command
var killserverCmd = &cobra.Command{
	Use:   "kill-server",
	Short: "Kill ADB server in JSON format",
	Long:  `Executes "adb kill-server" and outputs the result as structured JSON.`,
	RunE:  runKillServer,
}

func init() {
	rootCmd.AddCommand(killserverCmd)
}

func runKillServer(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb kill-server
	output, err := executor.Execute("kill-server")
	if err != nil {
		return fmt.Errorf("failed to execute adb kill-server: %w", err)
	}
	
	// Parse output
	parser := parser.NewKillServerParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse kill-server output: %w", err)
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
