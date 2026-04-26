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

// disableUserCmd represents the disable-user command
var disableUserCmd = &cobra.Command{
	Use:   "disable-user <package>",
	Short: "Disable package for user",
	Long:  `Executes "adb shell pm disable-user <package>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runDisableUser,
}

func init() {
	pmCmd.AddCommand(disableUserCmd)
}

func runDisableUser(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell pm disable-user command", map[string]interface{}{"package": args[0]})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell pm disable-user <package>
	cmdStr := "shell pm disable-user " + args[0]
	output, err := executor.Execute(cmdStr)
	if err != nil {
		log.Error("Failed to execute adb shell pm disable-user", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError(cmdStr, err)
	}
	log.Debug("ADB shell pm disable-user command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output (reuse disable parser)
	pmDisableParser := parser.NewPmDisableParser()
	response, err := pmDisableParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse pm disable-user output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(pmDisableParser.Name(), err)
	}
	log.Info("Parsed pm disable-user output", map[string]interface{}{"success": response.DisableResult.Success})
	
	// Validate result
	if err := pmDisableParser.Validate(response); err != nil {
		log.Error("Failed to validate pm disable-user output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("pm disable-user", err.Error())
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
	log.Info("Shell pm disable-user command completed successfully", nil)
	
	return nil
}
