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

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear <package>",
	Short: "Clear package data",
	Long:  `Executes "adb shell pm clear <package>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runClear,
}

func init() {
	pmCmd.AddCommand(clearCmd)
}

func runClear(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell pm clear command", map[string]interface{}{"package": args[0]})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell pm clear <package>
	cmdStr := "shell pm clear " + args[0]
	output, err := executor.Execute(cmdStr)
	if err != nil {
		log.Error("Failed to execute adb shell pm clear", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError(cmdStr, err)
	}
	log.Debug("ADB shell pm clear command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	pmClearParser := parser.NewPmClearParser()
	response, err := pmClearParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse pm clear output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(pmClearParser.Name(), err)
	}
	log.Info("Parsed pm clear output", map[string]interface{}{"success": response.ClearResult.Success})
	
	// Validate result
	if err := pmClearParser.Validate(response); err != nil {
		log.Error("Failed to validate pm clear output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("pm clear", err.Error())
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
	log.Info("Shell pm clear command completed successfully", nil)
	
	return nil
}
