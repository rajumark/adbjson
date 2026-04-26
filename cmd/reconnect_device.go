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

// reconnectDeviceCmd represents the reconnect device command
var reconnectDeviceCmd = &cobra.Command{
	Use:   "reconnect device",
	Short: "Reconnect device in JSON format",
	Long:  `Executes "adb reconnect device" and outputs the result as structured JSON.`,
	RunE:  runReconnectDevice,
}

func init() {
	rootCmd.AddCommand(reconnectDeviceCmd)
}

func runReconnectDevice(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting reconnect device command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb reconnect device
	output, err := executor.ExecuteWithOutput("reconnect", "device")
	if err != nil {
		log.Error("Failed to execute adb reconnect device", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("reconnect device", err)
	}
	log.Debug("ADB reconnect device command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	reconnectParser := parser.NewReconnectParser()
	response, err := reconnectParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse reconnect device output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("reconnect device", err)
	}
	log.Info("Parsed reconnect device output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := reconnectParser.Validate(response); err != nil {
		log.Error("Failed to validate reconnect device output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("reconnect device", err.Error())
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
	log.Info("Reconnect device command completed successfully", nil)
	
	return nil
}
