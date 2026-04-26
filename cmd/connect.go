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

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect <host:port>",
	Short: "Connect to a device via TCP/IP in JSON format",
	Long:  `Executes "adb connect <host:port>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runConnect,
}

func init() {
	rootCmd.AddCommand(connectCmd)
}

func runConnect(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting connect command", nil)

	target := args[0]
	log.Info("Connecting to device", map[string]interface{}{"target": target})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb connect
	output, err := executor.Execute("connect", target)
	if err != nil {
		log.Error("Failed to execute adb connect", map[string]interface{}{"error": err.Error(), "target": target})
		return apperrors.NewADBExecutionError("connect", err)
	}
	log.Debug("ADB connect command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	connectParser := parser.NewConnectParser()
	response, err := connectParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse connect output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(connectParser.Name(), err)
	}
	log.Info("Parsed connect output", map[string]interface{}{"connected": response.Connected})
	
	// Validate result
	if err := connectParser.Validate(response); err != nil {
		log.Error("Failed to validate connect output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("connect", err.Error())
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
	log.Info("Connect command completed successfully", nil)
	
	return nil
}
