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

// uptimeCmd represents the uptime command
var uptimeCmd = &cobra.Command{
	Use:   "uptime",
	Short: "Show system uptime in JSON format",
	Long:  `Executes "adb shell uptime" and outputs the result as structured JSON.`,
	RunE:  runUptime,
}

func init() {
	// Add uptime command to shell
	shellCmd.AddCommand(uptimeCmd)
}

func runUptime(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell uptime command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell uptime
	output, err := executor.Execute("shell", "uptime")
	if err != nil {
		log.Error("Failed to execute adb shell uptime", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell uptime", err)
	}
	log.Debug("ADB shell uptime command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	uptimeParser := parser.NewUptimeParser()
	response, err := uptimeParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse uptime output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("uptime", err)
	}
	log.Info("Parsed uptime output", map[string]interface{}{"uptime": response.Uptime})
	
	// Validate result
	if err := uptimeParser.Validate(response); err != nil {
		log.Error("Failed to validate uptime output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("uptime", err.Error())
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
	log.Info("Shell uptime command completed successfully", nil)
	
	return nil
}
