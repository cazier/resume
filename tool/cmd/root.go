/*
Copyright © 2022 Brendan Cazier
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var input string
var output string
var password string
var overwrite bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tool",
	Short: "A resume managing tool",
	Long:  `This encompasses encoding and decoding the resume file, as a `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
