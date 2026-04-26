package cmd

import (
	"adbjson/internal/formatter"
	"adbjson/internal/version"

	"github.com/spf13/cobra"
)

// cliVersionCmd represents the CLI version command
var cliVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show adbjson CLI version information",
	Long:  `Displays the current version of adbjson CLI with build information.`,
	RunE:  runCLIVersion,
}

func init() {
	rootCmd.AddCommand(cliVersionCmd)
}

func runCLIVersion(cmd *cobra.Command, args []string) error {
	buildInfo := version.GetBuildInfo()
	
	// Format output
	format := formatter.ParseFormat(outputFormat)
	output, err := formatter.FormatOutputString(buildInfo, format, compactOutput)
	if err != nil {
		return err
	}
	
	// Print to stdout
	cmd.Println(output)
	
	return nil
}
