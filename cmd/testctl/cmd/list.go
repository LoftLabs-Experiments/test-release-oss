package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List test resources",
	Long:  `List all test resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing test resources:")
		fmt.Println("- test-resource-1")
		fmt.Println("- test-resource-2") 
		fmt.Println("- test-resource-3")
	},
}
