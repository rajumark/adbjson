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

// unrootCmd represents the unroot command
var unrootCmd = &cobra.Command{
	Use:   "unroot",
	Short: "Restore adbd non-root privileges in JSON format",
	Long:  `Executes "adb unroot" and outputs the result as structured JSON.`,
	RunE:  runUnroot,
}

func init() {
	rootCmd.AddCommand(unrootCmd)
}

func runUnroot(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting unroot command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb unroot
	output, err := executor.ExecuteWithOutput("unroot")
	if err != nil {
		log.Error("Failed to execute adb unroot", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("unroot", err)
	}
	log.Debug("ADB unroot command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	unrootParser := parser.NewUnrootParser()
	response, err := unrootParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse unroot output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(unrootParser.Name(), err)
	}
	log.Info("Parsed unroot output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := unrootParser.Validate(response); err != nil {
		log.Error("Failed to validate unroot output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("unroot", err.Error())
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
	log.Info("Unroot command completed successfully", nil)
	
	return nil
}
