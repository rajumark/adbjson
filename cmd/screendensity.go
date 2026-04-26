package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// screendensityCmd represents the screendensity command
var screendensityCmd = &cobra.Command{
	Use:   "wm-density",
	Short: "Get screen density in JSON format",
	Long:  `Executes "adb shell wm density" and outputs the result as structured JSON.`,
	RunE:  runScreenDensity,
}

func init() {
	rootCmd.AddCommand(screendensityCmd)
}

func runScreenDensity(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb shell wm density
	output, err := executor.Execute("shell", "wm", "density")
	if err != nil {
		return fmt.Errorf("failed to execute adb shell wm density: %w", err)
	}
	
	// Parse output
	parser := parser.NewScreenDensityParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse screendensity output: %w", err)
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
