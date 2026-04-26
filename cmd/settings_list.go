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

// settingsListCmd represents the settings list command
var settingsListCmd = &cobra.Command{
	Use:   "list <namespace>",
	Short: "List settings in JSON format",
	Long:  `Executes "adb shell settings list <namespace>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runSettingsList,
}

func init() {
	// Create a parent settings command
	settingsCmd := &cobra.Command{
		Use:   "settings",
		Short: "Settings commands",
		Long:  `Settings related commands.`,
	}
	shellCmd.AddCommand(settingsCmd)
	settingsCmd.AddCommand(settingsListCmd)
}

func runSettingsList(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting settings list command", nil)

	namespace := args[0]
	
	// Validate namespace
	validNamespaces := []string{"system", "secure", "global"}
	isValid := false
	for _, ns := range validNamespaces {
		if ns == namespace {
			isValid = true
			break
		}
	}
	if !isValid {
		log.Error("Invalid namespace", map[string]interface{}{"namespace": namespace})
		return fmt.Errorf("namespace must be one of: system, secure, global")
	}
	
	log.Info("Listing settings", map[string]interface{}{"namespace": namespace})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell settings list <namespace>
	output, err := executor.ExecuteWithOutput("shell", "settings", "list", namespace)
	if err != nil {
		log.Error("Failed to execute adb shell settings list", map[string]interface{}{"error": err.Error(), "namespace": namespace})
		return apperrors.NewADBExecutionError("shell settings list", err)
	}
	log.Debug("ADB shell settings list command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	settingsListParser := parser.NewSettingsListParser()
	response, err := settingsListParser.Parse(output, namespace)
	if err != nil {
		log.Error("Failed to parse settings list output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(settingsListParser.Name(), err)
	}
	log.Info("Parsed settings list output", map[string]interface{}{"setting_count": response.Count})
	
	// Validate result
	if err := settingsListParser.Validate(response); err != nil {
		log.Error("Failed to validate settings list output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("settings list", err.Error())
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
	log.Info("Settings list command completed successfully", nil)
	
	return nil
}
