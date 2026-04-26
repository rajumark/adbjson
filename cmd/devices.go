package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List connected ADB devices in JSON format",
	Long:  `Executes "adb devices" and outputs the result as structured JSON.`,
	RunE:  runDevices,
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}

func runDevices(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb devices
	output, err := executor.Execute("devices")
	if err != nil {
		return fmt.Errorf("failed to execute adb devices: %w", err)
	}
	
	// Parse output
	parser := parser.NewDevicesParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse devices output: %w", err)
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
