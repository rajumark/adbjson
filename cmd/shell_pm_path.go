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

// pathCmd represents the path command
var pathCmd = &cobra.Command{
	Use:   "path <package>",
	Short: "Get package path",
	Long:  `Executes "adb shell pm path <package>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runPath,
}

func init() {
	pmCmd.AddCommand(pathCmd)
}

func runPath(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell pm path command", map[string]interface{}{"package": args[0]})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell pm path <package>
	cmdStr := "shell pm path " + args[0]
	output, err := executor.Execute(cmdStr)
	if err != nil {
		log.Error("Failed to execute adb shell pm path", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError(cmdStr, err)
	}
	log.Debug("ADB shell pm path command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	pmPathParser := parser.NewPmPathParser()
	response, err := pmPathParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse pm path output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(pmPathParser.Name(), err)
	}
	log.Info("Parsed pm path output", map[string]interface{}{"package": response.PackagePath.Package})
	
	// Validate result
	if err := pmPathParser.Validate(response); err != nil {
		log.Error("Failed to validate pm path output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("pm path", err.Error())
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
	log.Info("Shell pm path command completed successfully", nil)
	
	return nil
}
