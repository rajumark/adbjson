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

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable <package>",
	Short: "Disable package",
	Long:  `Executes "adb shell pm disable <package>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runDisable,
}

func init() {
	pmCmd.AddCommand(disableCmd)
}

func runDisable(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell pm disable command", map[string]interface{}{"package": args[0]})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell pm disable <package>
	cmdStr := "shell pm disable " + args[0]
	output, err := executor.Execute(cmdStr)
	if err != nil {
		log.Error("Failed to execute adb shell pm disable", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError(cmdStr, err)
	}
	log.Debug("ADB shell pm disable command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	pmDisableParser := parser.NewPmDisableParser()
	response, err := pmDisableParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse pm disable output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(pmDisableParser.Name(), err)
	}
	log.Info("Parsed pm disable output", map[string]interface{}{"success": response.DisableResult.Success})
	
	// Validate result
	if err := pmDisableParser.Validate(response); err != nil {
		log.Error("Failed to validate pm disable output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("pm disable", err.Error())
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
	log.Info("Shell pm disable command completed successfully", nil)
	
	return nil
}
