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

var showAPKPath bool
var thirdPartyOnly bool
var systemOnly bool
var disabledOnly bool
var enabledOnly bool
var showInstaller bool
var includeUninstalled bool
var filter string

// packagesCmd represents the packages command
var packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "List all packages",
	Long:  `Executes "adb shell pm list packages" and outputs the result as structured JSON.`,
	RunE:  runPackages,
}

func init() {
	listCmd.AddCommand(packagesCmd)
	
	packagesCmd.Flags().BoolVarP(&showAPKPath, "show-apk-path", "f", false, "Show APK path")
	packagesCmd.Flags().BoolVarP(&thirdPartyOnly, "third-party", "3", false, "List third-party packages only")
	packagesCmd.Flags().BoolVarP(&systemOnly, "system", "s", false, "List system packages only")
	packagesCmd.Flags().BoolVarP(&disabledOnly, "disabled", "d", false, "List disabled packages only")
	packagesCmd.Flags().BoolVarP(&enabledOnly, "enabled", "e", false, "List enabled packages only")
	packagesCmd.Flags().BoolVarP(&showInstaller, "installer", "i", false, "Show package installer")
	packagesCmd.Flags().BoolVarP(&includeUninstalled, "uninstalled", "u", false, "Include uninstalled packages")
	packagesCmd.Flags().StringVarP(&filter, "filter", "", "", "Filter by package name")
}

func runPackages(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting shell pm list packages command", nil)

	// Create executor
	executor := adb.NewExecutor()
	log.Debug("Created ADB executor", nil)
	
	// Build command with flags
	cmdStr := "shell pm list packages"
	if showAPKPath {
		cmdStr += " -f"
	}
	if thirdPartyOnly {
		cmdStr += " -3"
	}
	if systemOnly {
		cmdStr += " -s"
	}
	if disabledOnly {
		cmdStr += " -d"
	}
	if enabledOnly {
		cmdStr += " -e"
	}
	if showInstaller {
		cmdStr += " -i"
	}
	if includeUninstalled {
		cmdStr += " -u"
	}
	if filter != "" {
		cmdStr += " " + filter
	}
	
	// Run adb shell pm list packages with flags
	output, err := executor.Execute(cmdStr)
	if err != nil {
		log.Error("Failed to execute adb shell pm list packages", map[string]interface{}{"error": err.Error()})
		return apperrors.NewADBExecutionError(cmdStr, err)
	}
	log.Debug("ADB shell pm list packages command executed successfully", map[string]interface{}{"output_length": len(output)})
	
	// Parse output
	packagesParser := parser.NewPackagesParser()
	response, err := packagesParser.Parse(output)
	if err != nil {
		log.Error("Failed to parse packages output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewParseError(packagesParser.Name(), err)
	}
	log.Info("Parsed packages output", map[string]interface{}{"package_count": response.Count})
	
	// Validate result
	if err := packagesParser.Validate(response); err != nil {
		log.Error("Failed to validate packages output", map[string]interface{}{"error": err.Error()})
		return apperrors.NewValidationError("packages", err.Error())
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
	log.Info("Shell pm list packages command completed successfully", nil)
	
	return nil
}
