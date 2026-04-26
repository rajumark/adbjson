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

var tapDisplayID int
var source string

// inputTapCmd represents the input tap command
var inputTapCmd = &cobra.Command{
	Use:   "tap <x> <y>",
	Short: "Simulate touch tap at coordinates",
	Long:  `Executes "adb shell input tap" and outputs the result as structured JSON. Simulates touch events at specified screen coordinates.`,
	Args:  cobra.ExactArgs(2),
	RunE:  runInputTap,
}

func init() {
	inputCmd.AddCommand(inputTapCmd)
	
	inputTapCmd.Flags().IntVarP(&tapDisplayID, "display", "d", 0, "Specify display ID to tap on")
	inputTapCmd.Flags().StringVarP(&source, "source", "s", "touchscreen", "Input source type")
}

func runInputTap(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell input tap command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Parse coordinates from args
	x := args[0]
	y := args[1]
	
	// Build arguments for the command
	arguments := []string{"input"}
	
	if source != "" {
		arguments = append(arguments, source)
	}
	
	if tapDisplayID != 0 {
		arguments = append(arguments, "-d", fmt.Sprintf("%d", tapDisplayID))
	}
	
	arguments = append(arguments, "tap", x, y)
	
	// Run adb shell input tap with arguments
	finalArgs := []string{"shell"}
	finalArgs = append(finalArgs, arguments...)
	output, err := executor.Execute(finalArgs...)
	if err != nil {
		log.Error("Failed to execute adb shell input tap", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell input tap", err)
	}
	log.Debug("ADB shell input tap command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	inputTapParser := parser.NewInputTapParser()
	response, err := inputTapParser.Parse(output, x, y, source, tapDisplayID)
	if err != nil {
		log.Error("Failed to parse input tap output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("input tap", err)
	}
	log.Info("Parsed input tap output", map[string]interface{}{"x": response.X, "y": response.Y, "source": response.Source})
	
	// Validate result
	if err := inputTapParser.Validate(response); err != nil {
		log.Error("Failed to validate input tap output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("input tap", err.Error())
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
	log.Info("Shell input tap command completed successfully", nil)
	
	return nil
}
