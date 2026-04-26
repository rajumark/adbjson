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

// packagesCmd represents the packages command
var packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "List all packages",
	Long:  `Executes "adb shell pm list packages" and outputs the result as structured JSON.`,
	RunE:  runPackages,
}

func init() {
	listCmd.AddCommand(packagesCmd)
}

func runPackages(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell pm list packages command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell pm list packages
	output, err := executor.Execute("shell pm list packages")
	if err != nil {
		log.Error("Failed to execute adb shell pm list packages", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell pm list packages", err)
	}
	log.Debug("ADB shell pm list packages command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	packagesParser := parser.NewPackagesParser()
	response, err := packagesParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse packages output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(packagesParser.Name(), err)
	}
	log.Info("Parsed packages output", map[string]interface{}{"package_count": response.Count})
	
	// Validate result
	if err := packagesParser.Validate(response); err != nil {
		log.Error("Failed to validate packages output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("packages", err.Error())
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
	log.Info("Shell pm list packages command completed successfully", nil)
	
	return nil
}
