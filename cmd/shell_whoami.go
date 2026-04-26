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

// whoamiCmd represents the whoami command
var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Get current user",
	Long:  `Executes "adb shell whoami" and outputs the result as structured JSON.`,
	RunE:  runWhoami,
}

func init() {
	shellCmd.AddCommand(whoamiCmd)
}

func runWhoami(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell whoami command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell whoami
	output, err := executor.Execute("shell whoami")
	if err != nil {
		log.Error("Failed to execute adb shell whoami", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell whoami", err)
	}
	log.Debug("ADB shell whoami command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	whoamiParser := parser.NewWhoamiParser()
	response, err := whoamiParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse whoami output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(whoamiParser.Name(), err)
	}
	log.Info("Parsed whoami output", map[string]interface{}{"username": response.CurrentUser.Username})
	
	// Validate result
	if err := whoamiParser.Validate(response); err != nil {
		log.Error("Failed to validate whoami output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("whoami", err.Error())
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
	log.Info("Shell whoami command completed successfully", nil)
	
	return nil
}
