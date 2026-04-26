package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// batteryCmd represents the battery command
var batteryCmd = &cobra.Command{
	Use:   "battery",
	Short: "Get battery information in JSON format",
	Long:  `Executes "adb shell dumpsys battery" and outputs the result as structured JSON.`,
	RunE:  runBattery,
}

func init() {
	rootCmd.AddCommand(batteryCmd)
}

func runBattery(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb shell dumpsys battery
	output, err := executor.Execute("shell", "dumpsys", "battery")
	if err != nil {
		return fmt.Errorf("failed to execute adb shell dumpsys battery: %w", err)
	}
	
	// Parse output
	parser := parser.NewBatteryParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse battery output: %w", err)
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
