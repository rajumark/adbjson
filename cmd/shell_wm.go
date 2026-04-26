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

// wmCmd represents the wm command
var wmCmd = &cobra.Command{
	Use:   "wm",
	Short: "Window manager commands",
	Long:  `Window manager related commands.`,
}

func init() {
	// Add wm command to shell
	shellCmd.AddCommand(wmCmd)
	
	// Add density command
	wmDensityCmd := &cobra.Command{
		Use:   "density",
		Short: "Get screen density in JSON format",
		Long:  `Executes "adb shell wm density" and outputs the result as structured JSON.`,
		RunE:  runWmDensity,
	}
	wmCmd.AddCommand(wmDensityCmd)
	
	// Add density reset command as subcommand of density
	wmDensityResetCmd := &cobra.Command{
		Use:   "reset",
		Short: "Reset screen density in JSON format",
		Long:  `Executes "adb shell wm density reset" and outputs the result as structured JSON.`,
		RunE:  runWmDensityReset,
	}
	wmDensityCmd.AddCommand(wmDensityResetCmd)
	
	// Add size command
	wmSizeCmd := &cobra.Command{
		Use:   "size",
		Short: "Get screen size in JSON format",
		Long:  `Executes "adb shell wm size" and outputs the result as structured JSON.`,
		RunE:  runWmSize,
	}
	wmCmd.AddCommand(wmSizeCmd)
	
	// Add size reset command as subcommand of size
	wmSizeResetCmd := &cobra.Command{
		Use:   "reset",
		Short: "Reset screen resolution in JSON format",
		Long:  `Executes "adb shell wm size reset" and outputs the result as structured JSON.`,
		RunE:  runWmSizeReset,
	}
	wmSizeCmd.AddCommand(wmSizeResetCmd)
}

func runWmDensity(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell wm density command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell wm density
	output, err := executor.Execute("shell", "wm", "density")
	if err != nil {
		log.Error("Failed to execute adb shell wm density", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell wm density", err)
	}
	log.Debug("ADB shell wm density command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	densityParser := parser.NewScreenDensityParser()
	response, err := densityParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse wm density output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("wm density", err)
	}
	log.Info("Parsed wm density output", map[string]interface{}{"density": response.PhysicalDensity})
	
	// Format output
	format := formatter.ParseFormat(outputFormat)
	formattedOutput, err := formatter.FormatOutputString(response, format, compactOutput)
	if err != nil {
		log.Error("Failed to format output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewMarshalError(err)
	}
	
	// Print to stdout
	fmt.Println(formattedOutput)
	log.Info("Shell wm density command completed successfully", nil)
	
	return nil
}

func runWmSize(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell wm size command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell wm size
	output, err := executor.Execute("shell", "wm", "size")
	if err != nil {
		log.Error("Failed to execute adb shell wm size", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell wm size", err)
	}
	log.Debug("ADB shell wm size command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	sizeParser := parser.NewScreenSizeParser()
	response, err := sizeParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse wm size output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError("wm size", err)
	}
	log.Info("Parsed wm size output", map[string]interface{}{"size": response.PhysicalSize})
	
	// Format output
	format := formatter.ParseFormat(outputFormat)
	formattedOutput, err := formatter.FormatOutputString(response, format, compactOutput)
	if err != nil {
		log.Error("Failed to format output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewMarshalError(err)
	}
	
	// Print to stdout
	fmt.Println(formattedOutput)
	log.Info("Shell wm size command completed successfully", nil)
	
	return nil
}

func runWmDensityReset(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell wm density reset command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell wm density reset
	output, err := executor.ExecuteWithOutput("shell", "wm", "density", "reset")
	if err != nil {
		log.Error("Failed to execute adb shell wm density reset", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell wm density reset", err)
	}
	log.Debug("ADB shell wm density reset command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	wmDensityResetParser := parser.NewWmDensityResetParser()
	response, err := wmDensityResetParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse wm density reset output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(wmDensityResetParser.Name(), err)
	}
	log.Info("Parsed wm density reset output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := wmDensityResetParser.Validate(response); err != nil {
		log.Error("Failed to validate wm density reset output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("wm density reset", err.Error())
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
	log.Info("Shell wm density reset command completed successfully", nil)
	
	return nil
}

func runWmSizeReset(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell wm size reset command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Run adb shell wm size reset
	output, err := executor.ExecuteWithOutput("shell", "wm", "size", "reset")
	if err != nil {
		log.Error("Failed to execute adb shell wm size reset", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError("shell wm size reset", err)
	}
	log.Debug("ADB shell wm size reset command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	wmSizeResetParser := parser.NewWmSizeResetParser()
	response, err := wmSizeResetParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse wm size reset output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(wmSizeResetParser.Name(), err)
	}
	log.Info("Parsed wm size reset output", map[string]interface{}{"success": response.Success})
	
	// Validate result
	if err := wmSizeResetParser.Validate(response); err != nil {
		log.Error("Failed to validate wm size reset output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("wm size reset", err.Error())
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
	log.Info("Shell wm size reset command completed successfully", nil)
	
	return nil
}
