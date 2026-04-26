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

var showBatteries bool
var showAdapters bool
var showCooling bool
var showTemperatures bool
var showAll bool

// acpiCmd represents the acpi command
var acpiCmd = &cobra.Command{
	Use:   "acpi",
	Short: "Show ACPI power and thermal information",
	Long:  `Executes "adb shell acpi" and outputs the result as structured JSON.`,
	RunE:  runAcpi,
}

func init() {
	shellCmd.AddCommand(acpiCmd)
	
	acpiCmd.Flags().BoolVarP(&showBatteries, "batteries", "b", false, "Show batteries")
	acpiCmd.Flags().BoolVarP(&showAdapters, "adapters", "a", false, "Show power adapters")
	acpiCmd.Flags().BoolVarP(&showCooling, "cooling", "c", false, "Show cooling device state")
	acpiCmd.Flags().BoolVarP(&showTemperatures, "temperatures", "t", false, "Show temperatures")
	acpiCmd.Flags().BoolVarP(&showAll, "all", "V", false, "Show everything")
}

func runAcpi(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell acpi command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Build arguments for the command
	acpiArgs := []string{"acpi"}
	if showBatteries {
		acpiArgs = append(acpiArgs, "-b")
	}
	if showAdapters {
		acpiArgs = append(acpiArgs, "-a")
	}
	if showCooling {
		acpiArgs = append(acpiArgs, "-c")
	}
	if showTemperatures {
		acpiArgs = append(acpiArgs, "-t")
	}
	if showAll {
		acpiArgs = append(acpiArgs, "-V")
	}
	
	// If no flags specified, default to showing everything
	if len(acpiArgs) == 1 {
		acpiArgs = append(acpiArgs, "-V")
	}
	
	// Run adb shell acpi with flags
	var output string
	var err error
	
	// Build the command arguments
	cmdArgs := acpiArgs
	
	// Execute with output to handle non-zero exit codes
	if len(cmdArgs) == 2 {
		output, err = executor.ExecuteWithOutput("shell", cmdArgs[0], cmdArgs[1])
	} else if len(cmdArgs) == 3 {
		output, err = executor.ExecuteWithOutput("shell", cmdArgs[0], cmdArgs[1], cmdArgs[2])
	} else {
		// For simplicity, use the regular Execute method for now
		finalArgs := []string{"shell"}
		finalArgs = append(finalArgs, cmdArgs...)
		output, err = executor.Execute(finalArgs...)
	}
	
	if err != nil {
		// Check if we got any output despite the error
		if output == "" {
			log.Error("Failed to execute adb shell acpi", map[string]interface{}{"error": err.Error()})
			return apperrors.NewADBExecutionError("shell acpi", err)
		}
		// Log warning but continue with the output we got
		log.Warn("ADB shell acpi command returned non-zero exit code but produced output", map[string]interface{}{"error": err.Error()})
	}
	log.Debug("ADB shell acpi command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	acpiParser := parser.NewAcpiParser()
	response, err := acpiParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse acpi output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("acpi", err)
	}
	log.Info("Parsed acpi output", map[string]interface{}{"section_count": len(response.Sections)})
	
	// Validate result
	if err := acpiParser.Validate(response); err != nil {
		log.Error("Failed to validate acpi output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("acpi", err.Error())
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
	log.Info("Shell acpi command completed successfully", nil)
	
	return nil
}
