package cmd

import (
	"fmt"

	"github.com/LoftLabs-Experiments/test-release-oss/pkg/telemetry"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number of the test application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("testapp version %s\n", telemetry.GetVersion())
	},
}
