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

// lsProcCmd represents the ls /proc command
var lsProcCmd = &cobra.Command{
	Use:   "ls /proc",
	Short: "List /proc directory contents in JSON format",
	Long:  `Executes "adb shell ls /proc" and outputs the result as structured JSON.`,
	RunE:  runLsProc,
}

func init() {
	// Add ls /proc command to shell
	shellCmd.AddCommand(lsProcCmd)
}

func runLsProc(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell ls /proc command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell ls /proc
	output, err := executor.Execute("shell", "ls", "/proc")
	if err != nil {
		log.Error("Failed to execute adb shell ls /proc", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell ls /proc", err)
	}
	log.Debug("ADB shell ls /proc command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	lsProcParser := parser.NewLsProcParser()
	response, err := lsProcParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse ls /proc output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("ls /proc", err)
	}
	log.Info("Parsed ls /proc output", map[string]interface{}{"item_count": len(response.Items)})
	
	// Validate result
	if err := lsProcParser.Validate(response); err != nil {
		log.Error("Failed to validate ls /proc output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("ls /proc", err.Error())
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
	log.Info("Shell ls /proc command completed successfully", nil)
	
	return nil
}
