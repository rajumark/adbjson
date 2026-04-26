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

// reconnectCmd represents the reconnect command
var reconnectCmd = &cobra.Command{
	Use:   "reconnect",
	Short: "Reconnect device in JSON format",
	Long:  `Executes "adb reconnect" and outputs the result as structured JSON.`,
	RunE:  runReconnect,
}

func init() {
	rootCmd.AddCommand(reconnectCmd)
}

func runReconnect(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting reconnect command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb reconnect
	output, err := executor.ExecuteWithOutput("reconnect")
	if err != nil {
		log.Error("Failed to execute adb reconnect", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("reconnect", err)
	}
	log.Debug("ADB reconnect command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	reconnectParser := parser.NewReconnectParser()
	response, err := reconnectParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse reconnect output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("reconnect", err)
	}
	log.Info("Parsed reconnect output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := reconnectParser.Validate(response); err != nil {
		log.Error("Failed to validate reconnect output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("reconnect", err.Error())
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
	log.Info("Reconnect command completed successfully", nil)
	
	return nil
}
