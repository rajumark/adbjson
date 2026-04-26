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

// instrumentationCmd represents the instrumentation command
var instrumentationCmd = &cobra.Command{
	Use:   "instrumentation",
	Short: "List instrumentation in JSON format",
	Long:  `Executes "adb shell pm list instrumentation" and outputs the result as structured JSON.`,
	RunE:  runInstrumentation,
}

func init() {
	rootCmd.AddCommand(instrumentationCmd)
}

func runInstrumentation(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting instrumentation command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell pm list instrumentation
	output, err := executor.Execute("shell pm list instrumentation")
	if err != nil {
		log.Error("Failed to execute adb shell pm list instrumentation", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell pm list instrumentation", err)
	}
	log.Debug("ADB shell pm list instrumentation command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	instrumentationParser := parser.NewInstrumentationParser()
	response, err := instrumentationParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse instrumentation output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(instrumentationParser.Name(), err)
	}
	log.Info("Parsed instrumentation output", map[string]interface{}{"instrumentation_count": response.Count})
	
	// Validate result
	if err := instrumentationParser.Validate(response); err != nil {
		log.Error("Failed to validate instrumentation output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("instrumentation", err.Error())
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
	log.Info("Instrumentation command completed successfully", nil)
	
	return nil
}
