package cmd

import (
	"fmt"
	"adbjson/internal/adb"
	apperrors "adbjson/internal/errors"
	"adbjson/internal/formatter"
	"adbjson/internal/logger"
	"adbjson/internal/parser"
	"strings"

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
	
	// Add activity command as subcommand of dumpsys
	activityCmd := &cobra.Command{
		Use:   "activity",
		Short: "Show activity manager dump in JSON format",
		Long:  `Executes "adb shell dumpsys activity" and outputs the result as structured JSON.`,
	}
	dumpsysCmd.AddCommand(activityCmd)
	
	// Add activities command as subcommand of activity
	activitiesCmd := &cobra.Command{
		Use:   "activities",
		Short: "Show activity manager activities dump in JSON format",
		Long:  `Executes "adb shell dumpsys activity activities" and outputs the result as structured JSON.`,
		RunE:  runDumpsysActivity,
	}
	activityCmd.AddCommand(activitiesCmd)
	
	// Add wifi command as subcommand of dumpsys
	wifiCmd := &cobra.Command{
		Use:   "wifi",
		Short: "Show WiFi service dump in JSON format",
		Long:  `Executes "adb shell dumpsys wifi" and outputs the result as structured JSON.`,
		RunE:  runDumpsysWifi,
	}
	dumpsysCmd.AddCommand(wifiCmd)
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

func runDumpsysActivity(cmd *cobra.Command, cmdArgs []string) error {
	log := logger.Get()
	
	// Determine if this is the activities subcommand
	commandPath := cmd.CommandPath()
	var dumpsysCommand string
	var logMessage string
	
	if strings.HasSuffix(commandPath, "dumpsys activity activities") {
		dumpsysCommand = "dumpsys activity activities"
		logMessage = "Starting shell dumpsys activity activities command"
	} else {
		dumpsysCommand = "dumpsys activity"
		logMessage = "Starting shell dumpsys activity command"
	}
	
	log.Info(logMessage, nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell dumpsys command
	parts := strings.Split(dumpsysCommand, " ")
	var args []string
	args = append(args, "shell")
	args = append(args, parts...)
	output, err := executor.Execute(args...)
	if err != nil {
		log.Error("Failed to execute adb shell "+dumpsysCommand, map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell "+dumpsysCommand, err)
	}
	log.Debug("ADB shell "+dumpsysCommand+" command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	dumpsysActivityParser := parser.NewDumpsysActivityParser()
	response, err := dumpsysActivityParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse dumpsys activity output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("dumpsys activity", err)
	}
	log.Info("Parsed dumpsys activity output", map[string]interface{}{"section_count": len(response.Sections)})
	
	// Validate result
	if err := dumpsysActivityParser.Validate(response); err != nil {
		log.Error("Failed to validate dumpsys activity output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("dumpsys activity", err.Error())
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
	log.Info("Shell "+dumpsysCommand+" command completed successfully", nil)
	
	return nil
}

func runDumpsysWifi(cmd *cobra.Command, cmdArgs []string) error {
	log := logger.Get()
	log.Info("Starting shell dumpsys wifi command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell dumpsys wifi
	output, err := executor.Execute("shell", "dumpsys", "wifi")
	if err != nil {
		log.Error("Failed to execute adb shell dumpsys wifi", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell dumpsys wifi", err)
	}
	log.Debug("ADB shell dumpsys wifi command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	dumpsysWifiParser := parser.NewDumpsysWifiParser()
	response, err := dumpsysWifiParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse dumpsys wifi output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("dumpsys wifi", err)
	}
	log.Info("Parsed dumpsys wifi output", map[string]interface{}{"section_count": len(response.Sections)})
	
	// Validate result
	if err := dumpsysWifiParser.Validate(response); err != nil {
		log.Error("Failed to validate dumpsys wifi output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("dumpsys wifi", err.Error())
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
	log.Info("Shell dumpsys wifi command completed successfully", nil)
	
	return nil
}
