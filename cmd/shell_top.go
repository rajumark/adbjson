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

// topCmd represents the top command
var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Show running processes with resource usage in JSON format",
	Long:  `Executes "adb shell top" and outputs the result as structured JSON.`,
	RunE:  runTop,
}

// -n flag for number of iterations
var topNFlag string

func init() {
	// Add top command to shell
	shellCmd.AddCommand(topCmd)
	
	// Add -n flag for number of iterations
	topCmd.Flags().StringVarP(&topNFlag, "n", "n", "1", "Number of iterations")
}

func runTop(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell top command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell top -n <iterations>
	output, err := executor.Execute("shell", "top", "-n", topNFlag)
	if err != nil {
		log.Error("Failed to execute adb shell top", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell top", err)
	}
	log.Debug("ADB shell top command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	topParser := parser.NewTopParser()
	response, err := topParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse top output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("top", err)
	}
	log.Info("Parsed top output", map[string]interface{}{"process_count": len(response.Processes)})
	
	// Validate result
	if err := topParser.Validate(response); err != nil {
		log.Error("Failed to validate top output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("top", err.Error())
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
	log.Info("Shell top command completed successfully", nil)
	
	return nil
}
