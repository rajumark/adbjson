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

// disconnectCmd represents the disconnect command
var disconnectCmd = &cobra.Command{
	Use:   "disconnect <host:port>",
	Short: "Disconnect from a device via TCP/IP in JSON format",
	Long:  `Executes "adb disconnect <host:port>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runDisconnect,
}

func init() {
	rootCmd.AddCommand(disconnectCmd)
}

func runDisconnect(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting disconnect command", nil)

	target := args[0]
	log.Info("Disconnecting from device", map[string]interface{}{"target": target})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb disconnect
	output, err := executor.ExecuteWithOutput("disconnect", target)
	if err != nil {
		log.Error("Failed to execute adb disconnect", map[string]interface{}{"error": err.Error(), "target": target})
		return apperrors.NewADBExecutionError("disconnect", err)
	}
	log.Debug("ADB disconnect command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	disconnectParser := parser.NewDisconnectParser()
	response, err := disconnectParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse disconnect output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(disconnectParser.Name(), err)
	}
	log.Info("Parsed disconnect output", map[string]interface{}{"disconnected": response.Disconnected})
	
	// Validate result
	if err := disconnectParser.Validate(response); err != nil {
		log.Error("Failed to validate disconnect output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("disconnect", err.Error())
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
	log.Info("Disconnect command completed successfully", nil)
	
	return nil
}
