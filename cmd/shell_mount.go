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

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Show mount points in JSON format",
	Long:  `Executes "adb shell mount" and outputs the result as structured JSON.`,
	RunE:  runMount,
}

func init() {
	// Add mount command to shell
	shellCmd.AddCommand(mountCmd)
}

func runMount(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell mount command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell mount
	output, err := executor.Execute("shell", "mount")
	if err != nil {
		log.Error("Failed to execute adb shell mount", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell mount", err)
	}
	log.Debug("ADB shell mount command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	mountParser := parser.NewMountParser()
	response, err := mountParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse mount output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("mount", err)
	}
	log.Info("Parsed mount output", map[string]interface{}{"mount_count": len(response.MountPoints)})
	
	// Validate result
	if err := mountParser.Validate(response); err != nil {
		log.Error("Failed to validate mount output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("mount", err.Error())
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
	log.Info("Shell mount command completed successfully", nil)
	
	return nil
}
