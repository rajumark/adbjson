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

// lsRootCmd represents the ls / command
var lsRootCmd = &cobra.Command{
	Use:   "ls-root",
	Short: "List root directory contents in JSON format",
	Long:  `Executes "adb shell ls /" and outputs the result as structured JSON.`,
	RunE:  runLsRoot,
}

func init() {
	// Add ls / command to shell
	shellCmd.AddCommand(lsRootCmd)
}

func runLsRoot(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell ls / command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell ls /
	output, err := executor.Execute("shell", "ls", "/")
	if err != nil {
		log.Error("Failed to execute adb shell ls /", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell ls /", err)
	}
	log.Debug("ADB shell ls / command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	lsRootParser := parser.NewLsRootParser()
	response, err := lsRootParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse ls / output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("ls /", err)
	}
	log.Info("Parsed ls / output", map[string]interface{}{"item_count": len(response.Items)})
	
	// Validate result
	if err := lsRootParser.Validate(response); err != nil {
		log.Error("Failed to validate ls / output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("ls /", err.Error())
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
	log.Info("Shell ls / command completed successfully", nil)
	
	return nil
}
