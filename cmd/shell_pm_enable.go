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

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable <package>",
	Short: "Enable package",
	Long:  `Executes "adb shell pm enable <package>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runEnable,
}

func init() {
	pmCmd.AddCommand(enableCmd)
}

func runEnable(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell pm enable command", map[string]interface{}{"package": args[0]})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell pm enable <package>
	cmdStr := "shell pm enable " + args[0]
	output, err := executor.Execute(cmdStr)
	if err != nil {
		log.Error("Failed to execute adb shell pm enable", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError(cmdStr, err)
	}
	log.Debug("ADB shell pm enable command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	pmEnableParser := parser.NewPmEnableParser()
	response, err := pmEnableParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse pm enable output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(pmEnableParser.Name(), err)
	}
	log.Info("Parsed pm enable output", map[string]interface{}{"success": response.EnableResult.Success})
	
	// Validate result
	if err := pmEnableParser.Validate(response); err != nil {
		log.Error("Failed to validate pm enable output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("pm enable", err.Error())
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
	log.Info("Shell pm enable command completed successfully", nil)
	
	return nil
}
