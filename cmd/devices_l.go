package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// devicesLCmd represents the devices-l command
var devicesLCmd = &cobra.Command{
	Use:   "devices-l",
	Short: "List connected devices with detailed info in JSON format",
	Long:  `Executes "adb devices -l" and outputs the result as structured JSON.`,
	RunE:  runDevicesL,
}

func init() {
	rootCmd.AddCommand(devicesLCmd)
}

func runDevicesL(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb devices -l
	output, err := executor.Execute("devices", "-l")
	if err != nil {
		return fmt.Errorf("failed to execute adb devices -l: %w", err)
	}
	
	// Parse output
	parser := parser.NewDevicesLParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse devices-l output: %w", err)
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
