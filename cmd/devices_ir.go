package cmd

import (
	"fmt"
	"adbjson/internal/ir"
	"adbjson/internal/logger"

	"github.com/spf13/cobra"
)

// devicesIRCmd represents the new IR-based devices command
var devicesIRCmd = &cobra.Command{
	Use:   "devices-ir",
	Short: "List connected ADB devices using new IR-based architecture",
	Long:  `Executes "adb devices" using the new IR-based architecture for better scalability and maintainability.`,
	RunE:  runDevicesIR,
}

func init() {
	rootCmd.AddCommand(devicesIRCmd)
}

func runDevicesIR(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	log.Info("Starting IR-based devices command", nil)

	// Create IR engine
	engine, err := ir.NewEngine()
	if err != nil {
		log.Error("Failed to create IR engine", map[string]interface{}{"error": err.Error()})
		return err
	}
	log.Debug("Created IR engine", nil)

	// Execute command using IR approach
	jsonOutput, err := engine.ExecuteCommand("devices", args)
	if err != nil {
		log.Error("Failed to execute IR-based devices command", map[string]interface{}{"error": err.Error()})
		return err
	}
	log.Debug("IR-based devices command executed successfully", map[string]interface{}{"output_length": len(jsonOutput)})

	// Print to stdout
	fmt.Println(string(jsonOutput))
	log.Info("IR-based devices command completed successfully", nil)

	return nil
}
