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

// dumpsysCmd represents the dumpsys command
var dumpsysCmd = &cobra.Command{
	Use:   "dumpsys",
	Short: "Dump system service information",
	Long:  `Dump system service information in JSON format.`,
}

func init() {
	// Add dumpsys command to shell
	shellCmd.AddCommand(dumpsysCmd)
	
	// Add battery command as subcommand of dumpsys
	batteryCmd := &cobra.Command{
		Use:   "battery",
		Short: "Get battery information in JSON format",
		Long:  `Executes "adb shell dumpsys battery" and outputs the result as structured JSON.`,
		RunE:  runDumpsysBattery,
	}
	dumpsysCmd.AddCommand(batteryCmd)
}

func runDumpsysBattery(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell dumpsys battery command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell dumpsys battery
	output, err := executor.Execute("shell", "dumpsys", "battery")
	if err != nil {
		log.Error("Failed to execute adb shell dumpsys battery", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell dumpsys battery", err)
	}
	log.Debug("ADB shell dumpsys battery command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	batteryParser := parser.NewBatteryParser()
	response, err := batteryParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse dumpsys battery output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("dumpsys battery", err)
	}
	log.Info("Parsed dumpsys battery output", map[string]interface{}{"level": response.Level})
	
	// Format output
	format := formatter.ParseFormat(outputFormat)
	formattedOutput, err := formatter.FormatOutputString(response, format, compactOutput)
	if err != nil {
		log.Error("Failed to format output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewMarshalError(err)
	}
	
	// Print to stdout
	fmt.Println(formattedOutput)
	log.Info("Shell dumpsys battery command completed successfully", nil)
	
	return nil
}
