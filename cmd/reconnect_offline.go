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

// reconnectOfflineCmd represents the reconnect offline command
var reconnectOfflineCmd = &cobra.Command{
	Use:   "reconnect offline",
	Short: "Reconnect offline device in JSON format",
	Long:  `Executes "adb reconnect offline" and outputs the result as structured JSON.`,
	RunE:  runReconnectOffline,
}

func init() {
	rootCmd.AddCommand(reconnectOfflineCmd)
}

func runReconnectOffline(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting reconnect offline command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb reconnect offline
	output, err := executor.ExecuteWithOutput("reconnect", "offline")
	if err != nil {
		log.Error("Failed to execute adb reconnect offline", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("reconnect offline", err)
	}
	log.Debug("ADB reconnect offline command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	reconnectParser := parser.NewReconnectParser()
	response, err := reconnectParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse reconnect offline output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("reconnect offline", err)
	}
	log.Info("Parsed reconnect offline output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := reconnectParser.Validate(response); err != nil {
		log.Error("Failed to validate reconnect offline output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("reconnect offline", err.Error())
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
	log.Info("Reconnect offline command completed successfully", nil)
	
	return nil
}
