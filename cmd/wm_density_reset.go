package cmd

import (
	"fmt"
	"adbjson/internal/adb"
	apperrors "adbjson/internal/errors"
	"adbjson/internal/formatter"
	"adbjson/internal/logger"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// wmDensityResetCmd represents the wm density reset command
var wmDensityResetCmd = &cobra.Command{
	Use:   "density reset",
	Short: "Reset screen density in JSON format",
	Long:  `Executes "adb shell wm density reset" and outputs the result as structured JSON.`,
	RunE:  runWmDensityReset,
}

func init() {
	// Find existing wm command or create new one
	for _, cmd := range shellCmd.Commands() {
		if cmd.Name() == "wm" {
			cmd.AddCommand(wmDensityResetCmd)
			return
		}
	}
	
	// Create a parent wm command if it doesn't exist
	wmCmd := &cobra.Command{
		Use:   "wm",
		Short: "Window manager commands",
		Long:  `Window manager related commands.`,
	}
	shellCmd.AddCommand(wmCmd)
	wmCmd.AddCommand(wmDensityResetCmd)
}

func runWmDensityReset(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting wm density reset command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell wm density reset
	output, err := executor.ExecuteWithOutput("shell", "wm", "density", "reset")
	if err != nil {
		log.Error("Failed to execute adb shell wm density reset", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell wm density reset", err)
	}
	log.Debug("ADB shell wm density reset command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	wmDensityResetParser := parser.NewWmDensityResetParser()
	response, err := wmDensityResetParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse wm density reset output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(wmDensityResetParser.Name(), err)
	}
	log.Info("Parsed wm density reset output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := wmDensityResetParser.Validate(response); err != nil {
		log.Error("Failed to validate wm density reset output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("wm density reset", err.Error())
	}
	
	// Format output
	format := formatter.ParseFormat(outputFormat)
	formattedOutput, err := formatter.FormatOutputString(response, format, compactOutput)
	if err != nil {
		log.Error("Failed to format output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewMarshalError(err)
	}
	
	// Print to stdout
	fmt.Println(formattedOutput)
	log.Info("WM density reset command completed successfully", nil)
	
	return nil
}
