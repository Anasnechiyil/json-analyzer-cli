package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "json-analyzer-cli",
	Short: "CLI tool to analyze large JSON files",
	Long:  `A tool to filter, query, and analyze large JSON files efficiently.`,
}

// Execute executes the root command and handles any errors that occur.
//
// No parameters.
// No return values.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
