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
	
	// Add connectivity command as subcommand of dumpsys
	connectivityCmd := &cobra.Command{
		Use:   "connectivity",
		Short: "Show connectivity service dump in JSON format",
		Long:  `Executes "adb shell dumpsys connectivity" and outputs the result as structured JSON.`,
		RunE:  runDumpsysConnectivity,
	}
	dumpsysCmd.AddCommand(connectivityCmd)
	
	// Add telephony.registry command as subcommand of dumpsys
	telephonyRegistryCmd := &cobra.Command{
		Use:   "telephony.registry",
		Short: "Show telephony registry dump in JSON format",
		Long:  `Executes "adb shell dumpsys telephony.registry" and outputs the result as structured JSON.`,
		RunE:  runDumpsysTelephonyRegistry,
	}
	dumpsysCmd.AddCommand(telephonyRegistryCmd)
	
	// Add window command as subcommand of dumpsys
	windowCmd := &cobra.Command{
		Use:   "window",
		Short: "Show window manager dump in JSON format",
		Long:  `Executes "adb shell dumpsys window" and outputs the result as structured JSON.`,
		RunE:  runDumpsysWindow,
	}
	dumpsysCmd.AddCommand(windowCmd)
	
	// Add input command as subcommand of dumpsys
	inputCmd := &cobra.Command{
		Use:   "input",
		Short: "Show input system dump in JSON format",
		Long:  `Executes "adb shell dumpsys input" and outputs the result as structured JSON.`,
		RunE:  runDumpsysInput,
	}
	dumpsysCmd.AddCommand(inputCmd)
	
	// Add power command as subcommand of dumpsys
	powerCmd := &cobra.Command{
		Use:   "power",
		Short: "Show power manager dump in JSON format",
		Long:  `Executes "adb shell dumpsys power" and outputs the result as structured JSON.`,
		RunE:  runDumpsysPower,
	}
	dumpsysCmd.AddCommand(powerCmd)
	
	// Add location command as subcommand of dumpsys
	locationCmd := &cobra.Command{
		Use:   "location",
		Short: "Show location service dump in JSON format",
		Long:  `Executes "adb shell dumpsys location" and outputs the result as structured JSON.`,
		RunE:  runDumpsysLocation,
	}
	dumpsysCmd.AddCommand(locationCmd)
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

func runDumpsysConnectivity(cmd *cobra.Command, cmdArgs []string) error {
	log := logger.Get()
	log.Info("Starting shell dumpsys connectivity command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell dumpsys connectivity
	output, err := executor.Execute("shell", "dumpsys", "connectivity")
	if err != nil {
		log.Error("Failed to execute adb shell dumpsys connectivity", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell dumpsys connectivity", err)
	}
	log.Debug("ADB shell dumpsys connectivity command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	dumpsysConnectivityParser := parser.NewDumpsysConnectivityParser()
	response, err := dumpsysConnectivityParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse dumpsys connectivity output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("dumpsys connectivity", err)
	}
	log.Info("Parsed dumpsys connectivity output", map[string]interface{}{"section_count": len(response.Sections)})
	
	// Validate result
	if err := dumpsysConnectivityParser.Validate(response); err != nil {
		log.Error("Failed to validate dumpsys connectivity output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("dumpsys connectivity", err.Error())
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
	log.Info("Shell dumpsys connectivity command completed successfully", nil)
	
	return nil
}

func runDumpsysTelephonyRegistry(cmd *cobra.Command, cmdArgs []string) error {
	log := logger.Get()
	log.Info("Starting shell dumpsys telephony.registry command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell dumpsys telephony.registry
	output, err := executor.Execute("shell", "dumpsys", "telephony.registry")
	if err != nil {
		log.Error("Failed to execute adb shell dumpsys telephony.registry", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell dumpsys telephony.registry", err)
	}
	log.Debug("ADB shell dumpsys telephony.registry command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	dumpsysTelephonyRegistryParser := parser.NewDumpsysTelephonyRegistryParser()
	response, err := dumpsysTelephonyRegistryParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse dumpsys telephony.registry output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("dumpsys telephony.registry", err)
	}
	log.Info("Parsed dumpsys telephony.registry output", map[string]interface{}{"section_count": len(response.Sections)})
	
	// Validate result
	if err := dumpsysTelephonyRegistryParser.Validate(response); err != nil {
		log.Error("Failed to validate dumpsys telephony.registry output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("dumpsys telephony.registry", err.Error())
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
	log.Info("Shell dumpsys telephony.registry command completed successfully", nil)
	
	return nil
}

func runDumpsysWindow(cmd *cobra.Command, cmdArgs []string) error {
	log := logger.Get()
	log.Info("Starting shell dumpsys window command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell dumpsys window
	output, err := executor.Execute("shell", "dumpsys", "window")
	if err != nil {
		log.Error("Failed to execute adb shell dumpsys window", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell dumpsys window", err)
	}
	log.Debug("ADB shell dumpsys window command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	dumpsysWindowParser := parser.NewDumpsysWindowParser()
	response, err := dumpsysWindowParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse dumpsys window output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("dumpsys window", err)
	}
	log.Info("Parsed dumpsys window output", map[string]interface{}{"section_count": len(response.Sections)})
	
	// Validate result
	if err := dumpsysWindowParser.Validate(response); err != nil {
		log.Error("Failed to validate dumpsys window output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("dumpsys window", err.Error())
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
	log.Info("Shell dumpsys window command completed successfully", nil)
	
	return nil
}

func runDumpsysInput(cmd *cobra.Command, cmdArgs []string) error {
	log := logger.Get()
	log.Info("Starting shell dumpsys input command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell dumpsys input
	output, err := executor.Execute("shell", "dumpsys", "input")
	if err != nil {
		log.Error("Failed to execute adb shell dumpsys input", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell dumpsys input", err)
	}
	log.Debug("ADB shell dumpsys input command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	dumpsysInputParser := parser.NewDumpsysInputParser()
	response, err := dumpsysInputParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse dumpsys input output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("dumpsys input", err)
	}
	log.Info("Parsed dumpsys input output", map[string]interface{}{"section_count": len(response.Sections)})
	
	// Validate result
	if err := dumpsysInputParser.Validate(response); err != nil {
		log.Error("Failed to validate dumpsys input output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("dumpsys input", err.Error())
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
	log.Info("Shell dumpsys input command completed successfully", nil)
	
	return nil
}

func runDumpsysPower(cmd *cobra.Command, cmdArgs []string) error {
	log := logger.Get()
	log.Info("Starting shell dumpsys power command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell dumpsys power
	output, err := executor.Execute("shell", "dumpsys", "power")
	if err != nil {
		log.Error("Failed to execute adb shell dumpsys power", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell dumpsys power", err)
	}
	log.Debug("ADB shell dumpsys power command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	dumpsysPowerParser := parser.NewDumpsysPowerParser()
	response, err := dumpsysPowerParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse dumpsys power output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("dumpsys power", err)
	}
	log.Info("Parsed dumpsys power output", map[string]interface{}{"section_count": len(response.Sections)})
	
	// Validate result
	if err := dumpsysPowerParser.Validate(response); err != nil {
		log.Error("Failed to validate dumpsys power output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("dumpsys power", err.Error())
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
	log.Info("Shell dumpsys power command completed successfully", nil)
	
	return nil
}

func runDumpsysLocation(cmd *cobra.Command, cmdArgs []string) error {
	log := logger.Get()
	log.Info("Starting shell dumpsys location command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell dumpsys location
	output, err := executor.Execute("shell", "dumpsys", "location")
	if err != nil {
		log.Error("Failed to execute adb shell dumpsys location", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell dumpsys location", err)
	}
	log.Debug("ADB shell dumpsys location command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	dumpsysLocationParser := parser.NewDumpsysLocationParser()
	response, err := dumpsysLocationParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse dumpsys location output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("dumpsys location", err)
	}
	log.Info("Parsed dumpsys location output", map[string]interface{}{"section_count": len(response.Sections)})
	
	// Validate result
	if err := dumpsysLocationParser.Validate(response); err != nil {
		log.Error("Failed to validate dumpsys location output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("dumpsys location", err.Error())
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
	log.Info("Shell dumpsys location command completed successfully", nil)
	
	return nil
}
