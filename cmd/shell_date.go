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

// dateCmd represents the date command
var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Show system date and time in JSON format",
	Long:  `Executes "adb shell date" and outputs the result as structured JSON.`,
	RunE:  runDate,
}

func init() {
	// Add date command to shell
	shellCmd.AddCommand(dateCmd)
}

func runDate(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell date command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell date
	output, err := executor.Execute("shell", "date")
	if err != nil {
		log.Error("Failed to execute adb shell date", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell date", err)
	}
	log.Debug("ADB shell date command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	dateParser := parser.NewDateParser()
	response, err := dateParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse date output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("date", err)
	}
	log.Info("Parsed date output", map[string]interface{}{"datetime": response.DateTime})
	
	// Validate result
	if err := dateParser.Validate(response); err != nil {
		log.Error("Failed to validate date output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("date", err.Error())
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
	log.Info("Shell date command completed successfully", nil)
	
	return nil
}
