/*
Copyright © 2022 Brendan Cazier
*/
package cmd

import (
	crypt "tool/main/crypt"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a resume JSON file",
	Long: `
Encrypt your JSON file to prevent data scrapers from harvesting
your information about you from a public Github repo.`,
	Run: func(cmd *cobra.Command, args []string) {
		crypt.Conversion(password, input, output, overwrite, crypt.Encrypt)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringVarP(&password, "password", "p", "", "crypto password")
	encryptCmd.Root().MarkFlagRequired("password")
	encryptCmd.Flags().StringVarP(&input, "input", "i", "./resume.json", "path to the json file")
	encryptCmd.Flags().StringVarP(&output, "output", "o", "./resume.json.enc", "path to save file")
	encryptCmd.Flags().BoolVarP(&overwrite, "overwrite", "y", false, "overwrite existing files")
}
