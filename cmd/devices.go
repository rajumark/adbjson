package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"adbjson/internal/adb"
	"adbjson/internal/errors"
	"adbjson/internal/ir"
	"adbjson/internal/logger"

	"github.com/spf13/cobra"
)

// devicesCmd represents the devices command using IR architecture
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List connected ADB devices in JSON format",
	Long:  `Executes "adb devices" and outputs the result as structured JSON using the IR engine.`,
	RunE:  runDevices,
}

// -l flag for detailed listing
var devicesLFlag bool
var devicesRawFlag bool
var devicesEnrichFlag bool

func init() {
	rootCmd.AddCommand(devicesCmd)
	
	// Add flags
	devicesCmd.Flags().BoolVarP(&devicesLFlag, "l", "l", false, "List devices in detailed format")
	devicesCmd.Flags().Bool("raw", false, "Show raw ADB output")
	devicesCmd.Flags().Bool("enrich", false, "Add metadata and status envelope")
}

func runDevices(cmd *cobra.Command, args []string) error {
	log := logger.Get()
	
	// Check for --raw flag
	rawMode, _ := cmd.Flags().GetBool("raw")
	if rawMode {
		// Raw mode - execute ADB directly and show output
		executor := adb.NewExecutor()
		
		var output string
		var err error
		
		if devicesLFlag {
			output, err = executor.Execute("devices", "-l")
		} else {
			output, err = executor.Execute("devices")
		}
		
		if err != nil {
			return fmt.Errorf("ADB execution failed: %w", err)
		}
		
		fmt.Print(output)
		return nil
	}
	
	// Determine spec name based on -l flag
	specName := "devices"
	if devicesLFlag {
		specName = "devices-l"
	}

	// Create IR engine
	engine, err := ir.NewEngine()
	if err != nil {
		log.Error("Failed to create IR engine", map[string]interface{}{"error": err.Error()})
		return fmt.Errorf("failed to initialize command engine: %w", err)
	}
	log.Debug("Created IR engine", nil)
	
	// Check for --pretty flag override
	prettyOutput, _ := cmd.Flags().GetBool("pretty")
	if prettyOutput {
		compactOutput = false
	}
	
	// Only log in debug mode
	debugMode, _ := cmd.Flags().GetBool("debug")
	enrichMode, _ := cmd.Flags().GetBool("enrich")
	
	// Execute command using IR approach with format support
	output, err := engine.ExecuteCommandWithFormat(specName, []string{}, outputFormat, compactOutput)
	if err != nil {
		if debugMode {
			log.Error("Command execution failed", map[string]interface{}{"error": err.Error(), "spec": specName})
		}
		
		// Set appropriate exit code based on error type
		if appErr, ok := err.(*errors.AppError); ok {
			switch appErr.GetType() {
			case errors.ADBExecutionError:
				os.Exit(3) // ADB execution error
			case errors.ParseError:
				os.Exit(2) // Parsing error
			case errors.MarshalError:
				os.Exit(2) // JSON marshaling error
			default:
				os.Exit(1) // General error
			}
		}
		return fmt.Errorf("command execution failed: %w", err)
	}
	
	if debugMode {
		log.Debug("Command executed", map[string]interface{}{"spec": specName, "output_length": len(output)})
	}
	
	// Add enrichment if requested
	if enrichMode {
		// Parse the pure output to add metadata
		var pureData interface{}
		if outputFormat == "yaml" {
			// For YAML, we'd need a YAML parser, but for now just wrap
			pureData = map[string]interface{}{"yaml": string(output)}
		} else {
			json.Unmarshal(output, &pureData)
		}
		
		// Calculate count for enrichment
		count := 1
		if arr, ok := pureData.([]interface{}); ok {
			count = len(arr)
		} else if arr, ok := pureData.([]map[string]interface{}); ok {
			count = len(arr)
		}
		
		// Create enriched response
		enriched := map[string]interface{}{
			"status": "success",
			"meta": map[string]interface{}{
				"command": specName,
				"count":   count,
			},
			"data": pureData,
		}
		
		if compactOutput {
			output, _ = json.Marshal(enriched)
		} else {
			output, _ = json.MarshalIndent(enriched, "", "  ")
		}
	}
	
	// Print to stdout
	fmt.Println(string(output))
	
	return nil
}
