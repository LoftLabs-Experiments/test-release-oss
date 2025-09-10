package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "testapp",
	Short: "A dummy test application for workflow testing",
	Long:  `A dummy test application used to test GitHub workflows and CI/CD pipelines.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test application started!")
		logrus.Info("This is a dummy test application")
	},
}

// RunRoot executes the root command.
func RunRoot() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
}
