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

// rootCmd represents the root command
var rootCmdADB = &cobra.Command{
	Use:   "root",
	Short: "Restart adbd as root in JSON format",
	Long:  `Executes "adb root" and outputs the result as structured JSON.`,
	RunE:  runRoot,
}

func init() {
	rootCmd.AddCommand(rootCmdADB)
}

func runRoot(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting root command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb root
	output, err := executor.ExecuteWithOutput("root")
	if err != nil {
		log.Error("Failed to execute adb root", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("root", err)
	}
	log.Debug("ADB root command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	rootParser := parser.NewRootParser()
	response, err := rootParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse root output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(rootParser.Name(), err)
	}
	log.Info("Parsed root output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := rootParser.Validate(response); err != nil {
		log.Error("Failed to validate root output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("root", err.Error())
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
	log.Info("Root command completed successfully", nil)
	
	return nil
}
