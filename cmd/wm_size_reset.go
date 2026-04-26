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

// wmSizeResetCmd represents the wm size reset command
var wmSizeResetCmd = &cobra.Command{
	Use:   "size reset",
	Short: "Reset screen resolution in JSON format",
	Long:  `Executes "adb shell wm size reset" and outputs the result as structured JSON.`,
	RunE:  runWmSizeReset,
}

func init() {
	// Create a parent wm command
	wmCmd := &cobra.Command{
		Use:   "wm",
		Short: "Window manager commands",
		Long:  `Window manager related commands.`,
	}
	shellCmd.AddCommand(wmCmd)
	wmCmd.AddCommand(wmSizeResetCmd)
}

func runWmSizeReset(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting wm size reset command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell wm size reset
	output, err := executor.ExecuteWithOutput("shell", "wm", "size", "reset")
	if err != nil {
		log.Error("Failed to execute adb shell wm size reset", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell wm size reset", err)
	}
	log.Debug("ADB shell wm size reset command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	wmSizeResetParser := parser.NewWmSizeResetParser()
	response, err := wmSizeResetParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse wm size reset output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(wmSizeResetParser.Name(), err)
	}
	log.Info("Parsed wm size reset output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := wmSizeResetParser.Validate(response); err != nil {
		log.Error("Failed to validate wm size reset output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("wm size reset", err.Error())
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
	log.Info("WM size reset command completed successfully", nil)
	
	return nil
}
