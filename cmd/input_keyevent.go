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

// inputKeyeventCmd represents the input keyevent command
var inputKeyeventCmd = &cobra.Command{
	Use:   "keyevent <code>",
	Short: "Send key event in JSON format",
	Long:  `Executes "adb shell input keyevent <code>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runInputKeyevent,
}

func init() {
	// Create a parent input command
	inputCmd := &cobra.Command{
		Use:   "input",
		Short: "Input commands",
		Long:  `Input related commands.`,
	}
	shellCmd.AddCommand(inputCmd)
	inputCmd.AddCommand(inputKeyeventCmd)
}

func runInputKeyevent(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting input keyevent command", nil)

	keycode := args[0]
	log.Info("Sending key event", map[string]interface{}{"keycode": keycode})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell input keyevent <keycode>
	output, err := executor.ExecuteWithOutput("shell", "input", "keyevent", keycode)
	if err != nil {
		log.Error("Failed to execute adb shell input keyevent", map[string]interface{}{"error": err.Error(), "keycode": keycode})
		return apperrors.NewADBExecutionError("shell input keyevent", err)
	}
	log.Debug("ADB shell input keyevent command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	keyeventParser := parser.NewKeyeventParser()
	response, err := keyeventParser.Parse(output, keycode)
	if err != nil {
		log.Error("Failed to parse keyevent output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(keyeventParser.Name(), err)
	}
	log.Info("Parsed keyevent output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := keyeventParser.Validate(response); err != nil {
		log.Error("Failed to validate keyevent output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("keyevent", err.Error())
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
	log.Info("Input keyevent command completed successfully", nil)
	
	return nil
}
