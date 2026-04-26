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

var textDisplayID int
var textSource string

// inputTextCmd represents the input text command
var inputTextCmd = &cobra.Command{
	Use:   "text <string>",
	Short: "Input text string",
	Long:  `Executes "adb shell input text" and outputs the result as structured JSON. Simulates text input with support for multiple input sources and displays.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runInputText,
}

func init() {
	inputCmd.AddCommand(inputTextCmd)
	
	inputTextCmd.Flags().IntVarP(&textDisplayID, "display", "d", 0, "Specify display ID to input text on")
	inputTextCmd.Flags().StringVarP(&textSource, "source", "s", "keyboard", "Input source type")
}

func runInputText(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell input text command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Parse text from args
	text := args[0]
	
	// Build arguments for the command
	arguments := []string{"input"}
	
	if textSource != "" {
		arguments = append(arguments, textSource)
	}
	
	if textDisplayID != 0 {
		arguments = append(arguments, "-d", fmt.Sprintf("%d", textDisplayID))
	}
	
	arguments = append(arguments, "text", text)
	
	// Run adb shell input text with arguments
	finalArgs := []string{"shell"}
	finalArgs = append(finalArgs, arguments...)
	output, err := executor.Execute(finalArgs...)
	if err != nil {
		log.Error("Failed to execute adb shell input text", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell input text", err)
	}
	log.Debug("ADB shell input text command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	inputTextParser := parser.NewInputTextParser()
	response, err := inputTextParser.Parse(output, text, textSource, textDisplayID)
	if err != nil {
		log.Error("Failed to parse input text output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("input text", err)
	}
	log.Info("Parsed input text output", map[string]interface{}{"text": response.Text, "source": response.Source})
	
	// Validate result
	if err := inputTextParser.Validate(response); err != nil {
		log.Error("Failed to validate input text output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("input text", err.Error())
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
	log.Info("Shell input text command completed successfully", nil)
	
	return nil
}
