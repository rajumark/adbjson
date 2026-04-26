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

var uri string
var where string
var bind string
var selection string

// contentQueryCmd represents the content query command
var contentQueryCmd = &cobra.Command{
	Use:   "query --uri <uri>",
	Short: "Query content provider data",
	Long:  `Executes "adb shell content query" and outputs the result as structured JSON. Queries Android content providers for system and application data.`,
	RunE:  runContentQuery,
}

func init() {
	// Create content command parent
	contentCmd := &cobra.Command{
		Use:   "content",
		Short: "Content provider commands",
		Long:  `Android content provider commands for querying and managing system data.`,
	}
	shellCmd.AddCommand(contentCmd)
	contentCmd.AddCommand(contentQueryCmd)
	
	contentQueryCmd.Flags().StringVar(&uri, "uri", "", "Content provider URI (required)")
	contentQueryCmd.Flags().StringVar(&where, "where", "", "WHERE clause for filtering")
	contentQueryCmd.Flags().StringVar(&bind, "bind", "", "BIND clause for parameter binding")
	contentQueryCmd.Flags().StringVar(&selection, "selection", "", "Selection columns")
}

func runContentQuery(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell content query command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Build arguments for the command
	arguments := []string{"content", "query"}
	
	// Add required URI argument
	if uri == "" {
		return fmt.Errorf("--uri flag is required")
	}
	arguments = append(arguments, "--uri", uri)
	
	// Add optional flags
	if where != "" {
		arguments = append(arguments, "--where", where)
	}
	if bind != "" {
		arguments = append(arguments, "--bind", bind)
	}
	if selection != "" {
		arguments = append(arguments, "--selection", selection)
	}
	
	// Run adb shell content query with arguments
	finalArgs := []string{"shell"}
	finalArgs = append(finalArgs, arguments...)
	output, err := executor.Execute(finalArgs...)
	if err != nil {
		log.Error("Failed to execute adb shell content query", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell content query", err)
	}
	log.Debug("ADB shell content query command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	contentQueryParser := parser.NewContentQueryParser()
	response, err := contentQueryParser.Parse(output, uri)
	if err != nil {
		log.Error("Failed to parse content query output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("content query", err)
	}
	log.Info("Parsed content query output", map[string]interface{}{"row_count": response.Count})
	
	// Validate result
	if err := contentQueryParser.Validate(response); err != nil {
		log.Error("Failed to validate content query output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("content query", err.Error())
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
	log.Info("Shell content query command completed successfully", nil)
	
	return nil
}
