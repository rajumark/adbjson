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

// idCmd represents the id command
var idCmd = &cobra.Command{
	Use:   "id",
	Short: "Get user ID",
	Long:  `Executes "adb shell id" and outputs the result as structured JSON.`,
	RunE:  runId,
}

func init() {
	shellCmd.AddCommand(idCmd)
}

func runId(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell id command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell id
	output, err := executor.Execute("shell id")
	if err != nil {
		log.Error("Failed to execute adb shell id", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell id", err)
	}
	log.Debug("ADB shell id command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	idParser := parser.NewIdParser()
	response, err := idParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse id output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(idParser.Name(), err)
	}
	log.Info("Parsed id output", map[string]interface{}{"user_id": response.UserInfo.UserID})
	
	// Validate result
	if err := idParser.Validate(response); err != nil {
		log.Error("Failed to validate id output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("id", err.Error())
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
	log.Info("Shell id command completed successfully", nil)
	
	return nil
}
