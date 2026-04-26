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

// getpropCmd represents the getprop command
var getpropCmd = &cobra.Command{
	Use:   "getprop",
	Short: "Get system properties",
	Long:  `Executes "adb shell getprop" and outputs the result as structured JSON.`,
	RunE:  runGetprop,
}

func init() {
	shellCmd.AddCommand(getpropCmd)
}

func runGetprop(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell getprop command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell getprop
	output, err := executor.Execute("shell getprop")
	if err != nil {
		log.Error("Failed to execute adb shell getprop", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell getprop", err)
	}
	log.Debug("ADB shell getprop command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	getpropParser := parser.NewGetpropParser()
	response, err := getpropParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse getprop output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(getpropParser.Name(), err)
	}
	log.Info("Parsed getprop output", map[string]interface{}{"property_count": response.Count})
	
	// Validate result
	if err := getpropParser.Validate(response); err != nil {
		log.Error("Failed to validate getprop output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("getprop", err.Error())
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
	log.Info("Shell getprop command completed successfully", nil)
	
	return nil
}
