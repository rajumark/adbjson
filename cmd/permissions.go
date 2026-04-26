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

// permissionsCmd represents the permissions command
var permissionsCmd = &cobra.Command{
	Use:   "permissions",
	Short: "List permissions in JSON format",
	Long:  `Executes "adb shell pm list permissions" and outputs the result as structured JSON.`,
	RunE:  runPermissions,
}

func init() {
	rootCmd.AddCommand(permissionsCmd)
}

func runPermissions(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting permissions command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell pm list permissions
	output, err := executor.Execute("shell pm list permissions")
	if err != nil {
		log.Error("Failed to execute adb shell pm list permissions", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell pm list permissions", err)
	}
	log.Debug("ADB shell pm list permissions command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	permissionsParser := parser.NewPermissionsParser()
	response, err := permissionsParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse permissions output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(permissionsParser.Name(), err)
	}
	log.Info("Parsed permissions output", map[string]interface{}{"permission_count": response.Count})
	
	// Validate result
	if err := permissionsParser.Validate(response); err != nil {
		log.Error("Failed to validate permissions output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("permissions", err.Error())
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
	log.Info("Permissions command completed successfully", nil)
	
	return nil
}
