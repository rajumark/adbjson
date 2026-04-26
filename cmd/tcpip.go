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

// tcpipCmd represents the tcpip command
var tcpipCmd = &cobra.Command{
	Use:   "tcpip <port>",
	Short: "Allow device to listen on TCP/IP port in JSON format",
	Long:  `Executes "adb tcpip <port>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTcpip,
}

func init() {
	rootCmd.AddCommand(tcpipCmd)
}

func runTcpip(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting tcpip command", nil)

	port := args[0]
	log.Info("Setting TCP/IP port", map[string]interface{}{"port": port})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb tcpip <port>
	output, err := executor.ExecuteWithOutput("tcpip", port)
	if err != nil {
		log.Error("Failed to execute adb tcpip", map[string]interface{}{"error": err.Error(), "port": port})
		return apperrors.NewADBExecutionError("tcpip", err)
	}
	log.Debug("ADB tcpip command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	tcpipParser := parser.NewTcpipParser()
	response, err := tcpipParser.Parse(output, port)
	if err != nil {
		log.Error("Failed to parse tcpip output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(tcpipParser.Name(), err)
	}
	log.Info("Parsed tcpip output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := tcpipParser.Validate(response); err != nil {
		log.Error("Failed to validate tcpip output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("tcpip", err.Error())
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
	log.Info("Tcpip command completed successfully", nil)
	
	return nil
}
