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

// setenforceCmd represents the setenforce command
var setenforceCmd = &cobra.Command{
	Use:   "setenforce <0|1>",
	Short: "Set SELinux enforcing mode in JSON format",
	Long:  `Executes "adb shell setenforce <0|1>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runSetenforce,
}

func init() {
	shellCmd.AddCommand(setenforceCmd)
}

func runSetenforce(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell setenforce command", nil)

	mode := args[0]
	
	// Validate mode argument
	if mode != "0" && mode != "1" {
		log.Error("Invalid mode argument", map[string]interface{}{"mode": mode})
		return fmt.Errorf("mode must be 0 (permissive) or 1 (enforcing)")
	}
	
	log.Info("Setting SELinux enforcing mode", map[string]interface{}{"mode": mode})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell setenforce <mode>
	output, err := executor.ExecuteWithOutput("shell", "setenforce", mode)
	if err != nil {
		log.Error("Failed to execute adb shell setenforce", map[string]interface{}{"error": err.Error(), "mode": mode})
		return apperrors.NewADBExecutionError("shell setenforce", err)
	}
	log.Debug("ADB shell setenforce command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	setenforceParser := parser.NewSetenforceParser()
	response, err := setenforceParser.Parse(output, mode)
	if err != nil {
		log.Error("Failed to parse setenforce output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(setenforceParser.Name(), err)
	}
	log.Info("Parsed setenforce output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := setenforceParser.Validate(response); err != nil {
		log.Error("Failed to validate setenforce output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("setenforce", err.Error())
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
	log.Info("Shell setenforce command completed successfully", nil)
	
	return nil
}
