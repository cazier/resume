/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	crypt "tool/main/crypt"

	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt an encrypted JSON file back to plaintext",
	Long: `
Encrypt your JSON file to prevent data scrapers from harvesting
your information about you from a public Github repo.`,
	Run: func(cmd *cobra.Command, args []string) {
		crypt.Conversion(password, input, output, overwrite, crypt.Encrypt)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringVarP(&password, "password", "p", "", "crypto password")
	decryptCmd.Root().MarkFlagRequired("password")
	decryptCmd.Flags().StringVarP(&input, "input", "i", "./resume.json.enc", "path to the encrypted file")
	decryptCmd.Flags().StringVarP(&output, "output", "o", "./resume.json", "path to save file")
	decryptCmd.Flags().BoolVarP(&overwrite, "overwrite", "y", false, "overwrite existing files")
}
