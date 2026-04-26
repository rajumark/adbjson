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

var saveToFile string
var allDisplays bool
var displayID string
var pngFormat bool
var hintForSeamless bool

// screencapCmd represents the screencap command
var screencapCmd = &cobra.Command{
	Use:   "screencap [filename]",
	Short: "Capture device screenshot",
	Long:  `Executes "adb shell screencap" and outputs the result as structured JSON. Captures device screenshots with various options.`,
	RunE:  runScreencap,
}

func init() {
	shellCmd.AddCommand(screencapCmd)
	
	screencapCmd.Flags().StringVarP(&saveToFile, "output", "o", "", "Save screenshot to file")
	screencapCmd.Flags().BoolVarP(&allDisplays, "all", "a", false, "Capture all active displays")
	screencapCmd.Flags().StringVarP(&displayID, "display", "d", "", "Specify display ID to capture")
	screencapCmd.Flags().BoolVarP(&pngFormat, "png", "p", false, "Output in PNG format")
	screencapCmd.Flags().BoolVar(&hintForSeamless, "seamless", false, "Use seamless hint path")
}

func runScreencap(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell screencap command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Build arguments for the command
	arguments := []string{"screencap"}
	
	if allDisplays {
		arguments = append(arguments, "-a")
	}
	if displayID != "" {
		arguments = append(arguments, "-d", displayID)
	}
	if pngFormat {
		arguments = append(arguments, "-p")
	}
	if hintForSeamless {
		arguments = append(arguments, "--hint-for-seamless")
	}
	
	// Add filename if provided
	if len(args) > 0 {
		arguments = append(arguments, args[0])
	}
	
	// Run adb shell screencap with arguments
	finalArgs := []string{"shell"}
	finalArgs = append(finalArgs, arguments...)
	output, err := executor.Execute(finalArgs...)
	if err != nil {
		log.Error("Failed to execute adb shell screencap", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell screencap", err)
	}
	log.Debug("ADB shell screencap command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	screencapParser := parser.NewScreencapParser()
	response, err := screencapParser.Parse(output, saveToFile)
	if err != nil {
		log.Error("Failed to parse screencap output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("screencap", err)
	}
	log.Info("Parsed screencap output", map[string]interface{}{"data_size": len(response.Data), "format": response.Format})
	
	// Validate result
	if err := screencapParser.Validate(response); err != nil {
		log.Error("Failed to validate screencap output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("screencap", err.Error())
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
	log.Info("Shell screencap command completed successfully", nil)
	
	return nil
}
