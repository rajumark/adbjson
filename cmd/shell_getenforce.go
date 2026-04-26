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

// getenforceCmd represents the getenforce command
var getenforceCmd = &cobra.Command{
	Use:   "getenforce",
	Short: "Get SELinux status",
	Long:  `Executes "adb shell getenforce" and outputs the result as structured JSON.`,
	RunE:  runGetenforce,
}

func init() {
	shellCmd.AddCommand(getenforceCmd)
}

func runGetenforce(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell getenforce command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell getenforce
	output, err := executor.Execute("shell getenforce")
	if err != nil {
		log.Error("Failed to execute adb shell getenforce", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell getenforce", err)
	}
	log.Debug("ADB shell getenforce command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	getenforceParser := parser.NewGetenforceParser()
	response, err := getenforceParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse getenforce output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(getenforceParser.Name(), err)
	}
	log.Info("Parsed getenforce output", map[string]interface{}{"selinux_status": response.SELinuxStatus.Status})
	
	// Validate result
	if err := getenforceParser.Validate(response); err != nil {
		log.Error("Failed to validate getenforce output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("getenforce", err.Error())
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
	log.Info("Shell getenforce command completed successfully", nil)
	
	return nil
}
