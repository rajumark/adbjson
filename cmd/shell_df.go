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

// dfCmd represents the df command
var dfCmd = &cobra.Command{
	Use:   "df",
	Short: "Show disk usage in JSON format",
	Long:  `Executes "adb shell df" and outputs the result as structured JSON.`,
	RunE:  runDf,
}

// --human-readable flag for human readable format
var dfHumanReadableFlag bool

func init() {
	// Add df command to shell
	shellCmd.AddCommand(dfCmd)
	
	// Add --human-readable flag for human readable format
	dfCmd.Flags().BoolVar(&dfHumanReadableFlag, "human-readable", false, "Human readable format")
}

func runDf(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	
	// Determine command and parser based on --human-readable flag
	var command string
	var parserName string
	
	if dfHumanReadableFlag {
		command = "shell df -h"
		parserName = "df -h"
		log.Info("Starting shell df -h command", nil)
	} else {
		command = "shell df"
		parserName = "df"
		log.Info("Starting shell df command", nil)
	}

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell df command
	var output string
	var err error
	
	if dfHumanReadableFlag {
		output, err = executor.Execute("shell", "df", "-h")
	} else {
		output, err = executor.Execute("shell", "df")
	}
	
	if err != nil {
		log.Error("Failed to execute adb shell df", map[string]interface{}{"error": err.Error(), "command": command})
		return apperrors.NewADBExecutionError(command, err)
	}
	log.Debug("ADB shell df command executed successfully", map[string]interface{}{"output_length": len(output), "command": command})
	
	// Parse output
	dfParser := parser.NewDfParser()
	response, err := dfParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse df output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(parserName, err)
	}
	log.Info("Parsed df output", map[string]interface{}{"filesystem_count": len(response.Filesystems)})
	
	// Validate result
	if err := dfParser.Validate(response); err != nil {
		log.Error("Failed to validate df output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError(parserName, err.Error())
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
	log.Info("Shell df command completed successfully", nil)
	
	return nil
}
