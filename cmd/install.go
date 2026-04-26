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

var reinstall bool
var protectInstallDir bool
var testOnly bool
var installToSDCard bool
var allowDowngrade bool
var grantAllPermissions bool
var forceABI string

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install <apk>",
	Short: "Install APK",
	Long:  `Executes "adb install <apk>" and outputs the result as structured JSON.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)
	
	installCmd.Flags().BoolVarP(&reinstall, "reinstall", "r", false, "Reinstall package")
	installCmd.Flags().BoolVarP(&protectInstallDir, "protect", "l", false, "Protect installation directory")
	installCmd.Flags().BoolVarP(&testOnly, "test", "t", false, "Install test-only apps")
	installCmd.Flags().BoolVarP(&installToSDCard, "sdcard", "s", false, "Install to sdcard")
	installCmd.Flags().BoolVarP(&allowDowngrade, "downgrade", "d", false, "Allow downgrade")
	installCmd.Flags().BoolVarP(&grantAllPermissions, "grant", "g", false, "Grant all runtime permissions")
	installCmd.Flags().StringVar(&forceABI, "abi", "", "Force specific ABI")
}

func runInstall(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting install command", map[string]interface{}{"apk": args[0]})

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Build command with flags
	cmdStr := "install"
	if reinstall {
		cmdStr += " -r"
	}
	if protectInstallDir {
		cmdStr += " -l"
	}
	if testOnly {
		cmdStr += " -t"
	}
	if installToSDCard {
		cmdStr += " -s"
	}
	if allowDowngrade {
		cmdStr += " -d"
	}
	if grantAllPermissions {
		cmdStr += " -g"
	}
	if forceABI != "" {
		cmdStr += " --abi " + forceABI
	}
	cmdStr += " " + args[0]
	
	// Run adb install
	output, err := executor.Execute(cmdStr)
	if err != nil {
		log.Error("Failed to execute adb install", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError(cmdStr, err)
	}
	log.Debug("ADB install command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	installParser := parser.NewInstallParser()
	response, err := installParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse install output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(installParser.Name(), err)
	}
	log.Info("Parsed install output", map[string]interface{}{"success": response.InstallResult.Success})
	
	// Validate result
	if err := installParser.Validate(response); err != nil {
		log.Error("Failed to validate install output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("install", err.Error())
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
	log.Info("Install command completed successfully", nil)
	
	return nil
}
