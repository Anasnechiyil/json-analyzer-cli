package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(countCmd)
}

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count the number of keys in the JSON file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a JSON file.")
			return
		}
		countKeys(args[0])
	},
}

// countKeys counts the total number of keys in a JSON file.
//
// filename is the path to the JSON file to be processed.
// No return values.
func countKeys(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	keyCount := 0
	for decoder.More() {
		var data map[string]interface{}
		if err := decoder.Decode(&data); err != nil {
			fmt.Printf("Error decoding JSON: %v\n", err)
			return
		}
		keyCount += len(data)
	}
	fmt.Printf("Total keys: %d\n", keyCount)
}
