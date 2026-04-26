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

// vmstatCmd represents the vmstat command
var vmstatCmd = &cobra.Command{
	Use:   "vmstat",
	Short: "Show virtual memory statistics in JSON format",
	Long:  `Executes "adb shell vmstat" and outputs the result as structured JSON.`,
	RunE:  runVmstat,
}

func init() {
	// Add vmstat command to shell
	shellCmd.AddCommand(vmstatCmd)
}

func runVmstat(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell vmstat command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell vmstat
	output, err := executor.Execute("shell", "vmstat")
	if err != nil {
		log.Error("Failed to execute adb shell vmstat", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell vmstat", err)
	}
	log.Debug("ADB shell vmstat command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	vmstatParser := parser.NewVmstatParser()
	response, err := vmstatParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse vmstat output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("vmstat", err)
	}
	log.Info("Parsed vmstat output", map[string]interface{}{"free_memory": response.Memory.Free})
	
	// Validate result
	if err := vmstatParser.Validate(response); err != nil {
		log.Error("Failed to validate vmstat output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("vmstat", err.Error())
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
	log.Info("Shell vmstat command completed successfully", nil)
	
	return nil
}
