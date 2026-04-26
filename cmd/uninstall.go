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

var keepData bool

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall <package>",
	Short: "Uninstall package",
	Long:  `Executes "adb uninstall <package>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runUninstall,
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
	
	uninstallCmd.Flags().BoolVarP(&keepData, "keep-data", "k", false, "Uninstall but keep data")
}

func runUninstall(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting uninstall command", map[string]interface{}{"package": args[0]})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Build command with flags
	cmdStr := "uninstall"
	if keepData {
		cmdStr += " -k"
	}
	cmdStr += " " + args[0]
	
	// Run adb uninstall
	output, err := executor.Execute(cmdStr)
	if err != nil {
		log.Error("Failed to execute adb uninstall", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError(cmdStr, err)
	}
	log.Debug("ADB uninstall command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	uninstallParser := parser.NewUninstallParser()
	response, err := uninstallParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse uninstall output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(uninstallParser.Name(), err)
	}
	log.Info("Parsed uninstall output", map[string]interface{}{"success": response.UninstallResult.Success})
	
	// Validate result
	if err := uninstallParser.Validate(response); err != nil {
		log.Error("Failed to validate uninstall output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("uninstall", err.Error())
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
	log.Info("Uninstall command completed successfully", nil)
	
	return nil
}
