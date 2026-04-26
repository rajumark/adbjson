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

// freeCmd represents the free command
var freeCmd = &cobra.Command{
	Use:   "free",
	Short: "Show memory usage in JSON format",
	Long:  `Executes "adb shell free" and outputs the result as structured JSON.`,
	RunE:  runFree,
}

func init() {
	// Add free command to shell
	shellCmd.AddCommand(freeCmd)
}

func runFree(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell free command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell free
	output, err := executor.Execute("shell", "free")
	if err != nil {
		log.Error("Failed to execute adb shell free", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell free", err)
	}
	log.Debug("ADB shell free command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	freeParser := parser.NewFreeParser()
	response, err := freeParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse free output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("free", err)
	}
	log.Info("Parsed free output", map[string]interface{}{"memory_total": response.Memory.Total})
	
	// Validate result
	if err := freeParser.Validate(response); err != nil {
		log.Error("Failed to validate free output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("free", err.Error())
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
	log.Info("Shell free command completed successfully", nil)
	
	return nil
}
