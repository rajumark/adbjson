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

// librariesCmd represents the libraries command
var librariesCmd = &cobra.Command{
	Use:   "libraries",
	Short: "List device libraries",
	Long:  `Executes "adb shell pm list libraries" and outputs the result as structured JSON.`,
	RunE:  runLibraries,
}

func init() {
	listCmd.AddCommand(librariesCmd)
}

func runLibraries(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell pm list libraries command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell pm list libraries
	output, err := executor.Execute("shell pm list libraries")
	if err != nil {
		log.Error("Failed to execute adb shell pm list libraries", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell pm list libraries", err)
	}
	log.Debug("ADB shell pm list libraries command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	librariesParser := parser.NewLibrariesParser()
	response, err := librariesParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse libraries output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(librariesParser.Name(), err)
	}
	log.Info("Parsed libraries output", map[string]interface{}{"library_count": response.Count})
	
	// Validate result
	if err := librariesParser.Validate(response); err != nil {
		log.Error("Failed to validate libraries output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("libraries", err.Error())
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
	log.Info("Shell pm list libraries command completed successfully", nil)
	
	return nil
}
