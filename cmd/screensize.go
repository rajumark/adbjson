package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// screensizeCmd represents the screensize command
var screensizeCmd = &cobra.Command{
	Use:   "wm-size",
	Short: "Get screen size in JSON format",
	Long:  `Executes "adb shell wm size" and outputs the result as structured JSON.`,
	RunE:  runScreenSize,
}

func init() {
	rootCmd.AddCommand(screensizeCmd)
}

func runScreenSize(cmd *cobra.Command, args []string) error {
	// Create executor
	executor := adb.NewExecutor()
	
	// Run adb shell wm size
	output, err := executor.Execute("shell", "wm", "size")
	if err != nil {
		return fmt.Errorf("failed to execute adb shell wm size: %w", err)
	}
	
	// Parse output
	parser := parser.NewScreenSizeParser()
	response, err := parser.Parse(output)
	if err != nil {
		return fmt.Errorf("failed to parse screensize output: %w", err)
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
