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
	Use:   "resume",
	Short: "A resume managing tool",
	Long: `This tool is primarily intended to generate a resume file in various formats
from a provided JSON file. Additionally, there is some support for encoding/decoding
the JSON file (originally intended to prevent having your data scraped from a public
git repository. See the subcommands and their help pages for more details`,
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
