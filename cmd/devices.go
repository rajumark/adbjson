package cmd

import (
	"encoding/json"
	"fmt"
	"adbjson/internal/adb"
	apperrors "adbjson/internal/errors"
	"adbjson/internal/logger"
	"adbjson/internal/parser"

	"github.com/spf13/cobra"
)

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List connected ADB devices in JSON format",
	Long:  `Executes "adb devices" and outputs the result as structured JSON.`,
	RunE:  runDevices,
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}

func runDevices(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting devices command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb devices
	output, err := executor.Execute("devices")
	if err != nil {
		log.Error("Failed to execute adb devices", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("devices", err)
	}
	log.Debug("ADB devices command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	devicesParser := parser.NewDevicesParser()
	response, err := devicesParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse devices output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(devicesParser.Name(), err)
	}
	log.Info("Parsed devices output", map[string]interface{}{"device_count": len(response.Devices)})
	
	// Validate result
	if err := devicesParser.Validate(response); err != nil {
		log.Error("Failed to validate devices output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("devices", err.Error())
	}
	
	// Determine output format
	var jsonBytes []byte
	if compactOutput {
		jsonBytes, err = json.Marshal(response)
	} else {
		jsonBytes, err = json.MarshalIndent(response, "", "  ")
	}
	
	if err != nil {
		log.Error("Failed to marshal JSON", map[string]interface{}{"error": err.Error()})
		return apperrors.NewMarshalError(err)
	}
	
	// Print to stdout
	fmt.Println(string(jsonBytes))
	log.Info("Devices command completed successfully", nil)
	
	return nil
}
