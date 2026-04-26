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

// unameCmd represents the uname command
var unameCmd = &cobra.Command{
	Use:   "uname",
	Short: "Get system information",
	Long:  `Executes "adb shell uname -a" and outputs the result as structured JSON.`,
	RunE:  runUname,
}

func init() {
	shellCmd.AddCommand(unameCmd)
}

func runUname(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell uname command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell uname -a
	output, err := executor.Execute("shell uname -a")
	if err != nil {
		log.Error("Failed to execute adb shell uname -a", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell uname -a", err)
	}
	log.Debug("ADB shell uname -a command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	unameParser := parser.NewUnameParser()
	response, err := unameParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse uname output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(unameParser.Name(), err)
	}
	log.Info("Parsed uname output", map[string]interface{}{"kernel_name": response.SystemInfo.KernelName})
	
	// Validate result
	if err := unameParser.Validate(response); err != nil {
		log.Error("Failed to validate uname output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("uname", err.Error())
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
	log.Info("Shell uname command completed successfully", nil)
	
	return nil
}
