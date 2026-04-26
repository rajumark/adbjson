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

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List processes in JSON format",
	Long:  `Executes "adb shell ps" and outputs the result as structured JSON.`,
	RunE:  runPs,
}

// -A flag for showing all processes
var psAFlag bool

func init() {
	// Add ps command to shell
	shellCmd.AddCommand(psCmd)
	
	// Add -A flag for showing all processes
	psCmd.Flags().BoolVarP(&psAFlag, "A", "A", false, "Show all processes")
}

func runPs(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	
	// Determine command and parser based on -A flag
	var command string
	var parserName string
	
	if psAFlag {
		command = "shell ps -A"
		parserName = "ps -A"
		log.Info("Starting shell ps -A command", nil)
	} else {
		command = "shell ps"
		parserName = "ps"
		log.Info("Starting shell ps command", nil)
	}

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell ps command
	var output string
	var err error
	
	if psAFlag {
		output, err = executor.Execute("shell", "ps", "-A")
	} else {
		output, err = executor.Execute("shell", "ps")
	}
	
	if err != nil {
		log.Error("Failed to execute adb shell ps", map[string]interface{}{"error": err.Error(), "command": command})
		return apperrors.NewADBExecutionError(command, err)
	}
	log.Debug("ADB shell ps command executed successfully", map[string]interface{}{"output_length": len(output), "command": command})
	
	// Parse output
	psParser := parser.NewPsParser()
	response, err := psParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse ps output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(parserName, err)
	}
	log.Info("Parsed ps output", map[string]interface{}{"process_count": len(response.Processes)})
	
	// Validate result
	if err := psParser.Validate(response); err != nil {
		log.Error("Failed to validate ps output", map[string]interface{}{"error": err.Error()})
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
	log.Info("Shell ps command completed successfully", nil)
	
	return nil
}
