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

// featuresCmd represents the features command
var featuresCmd = &cobra.Command{
	Use:   "features",
	Short: "List device features",
	Long:  `Executes "adb shell pm list features" and outputs the result as structured JSON.`,
	RunE:  runFeatures,
}

func init() {
	listCmd.AddCommand(featuresCmd)
}

func runFeatures(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell pm list features command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell pm list features
	output, err := executor.Execute("shell pm list features")
	if err != nil {
		log.Error("Failed to execute adb shell pm list features", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell pm list features", err)
	}
	log.Debug("ADB shell pm list features command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	featuresParser := parser.NewFeaturesParser()
	response, err := featuresParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse features output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(featuresParser.Name(), err)
	}
	log.Info("Parsed features output", map[string]interface{}{"feature_count": response.Count})
	
	// Validate result
	if err := featuresParser.Validate(response); err != nil {
		log.Error("Failed to validate features output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("features", err.Error())
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
	log.Info("Shell pm list features command completed successfully", nil)
	
	return nil
}
