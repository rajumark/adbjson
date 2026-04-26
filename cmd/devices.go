package cmd

import (
	"fmt"
	"adbjson/internal/adb"
	apperrors "adbjson/internal/errors"
	"adbjson/internal/formatter"
	"adbjson/internal/logger"
	"adbjson/internal/model"
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

// -l flag for detailed listing
var devicesLFlag bool

func init() {
	rootCmd.AddCommand(devicesCmd)
	
	// Add -l flag for detailed listing
	devicesCmd.Flags().BoolVarP(&devicesLFlag, "l", "l", false, "List devices in detailed format")
}

func runDevices(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	// Determine command and parser based on -l flag
	var command string
	var parserName string
	
	if devicesLFlag {
		command = "devices -l"
		parserName = "devices -l"
		log.Info("Starting devices command with -l flag", nil)
	} else {
		command = "devices"
		parserName = "devices"
		log.Info("Starting devices command", nil)
	}

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb devices command
	var output string
	var err error
	
	if devicesLFlag {
		output, err = executor.Execute("devices", "-l")
	} else {
		output, err = executor.Execute("devices")
	}
	
	if err != nil {
		log.Error("Failed to execute adb devices", map[string]interface{}{"error": err.Error(), "command": command})
		return apperrors.NewADBExecutionError(command, err)
	}
	log.Debug("ADB devices command executed successfully", map[string]interface{}{"output_length": len(output), "command": command})
	
	// Parse output
	var response interface{}
	var parseErr error
	
	if devicesLFlag {
		devicesLParser := parser.NewDevicesLParser()
		response, parseErr = devicesLParser.Parse(output)
		if parseErr != nil {
			log.Error("Failed to parse devices -l output", map[string]interface{}{"error": parseErr.Error()})
			return apperrors.NewParseError(parserName, parseErr)
		}
		log.Info("Parsed devices -l output", map[string]interface{}{"device_count": len(response.(*model.DevicesResponse).Devices)})
	} else {
		devicesParser := parser.NewDevicesParser()
		response, parseErr = devicesParser.Parse(output)
		if parseErr != nil {
			log.Error("Failed to parse devices output", map[string]interface{}{"error": parseErr.Error()})
			return apperrors.NewParseError(parserName, parseErr)
		}
		log.Info("Parsed devices output", map[string]interface{}{"device_count": len(response.(*model.DevicesResponse).Devices)})
	}
	
	// Validation is handled by the parsers themselves
	
	// Format output
	format := formatter.ParseFormat(outputFormat)
	formattedOutput, err := formatter.FormatOutputString(response, format, compactOutput)
	if err != nil {
		log.Error("Failed to format output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewMarshalError(err)
	}
	
	// Print to stdout
	fmt.Println(formattedOutput)
	log.Info("Devices command completed successfully", nil)
	
	return nil
}
