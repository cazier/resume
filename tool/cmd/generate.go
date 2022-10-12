/*
Copyright © 2022 Brendan Cazier
*/
package cmd

import (
	"os"

	files "tool/main/files"
	marshall "tool/main/marshall"
	shared "tool/main/shared"

	"github.com/spf13/cobra"
)

var format []string
var template string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "create the output resume files",
	Long: `
Generate the generated resume files by filling a passed in template with the
supplied resume data. The output can contain plain text (.txt) files or websites
(.html), as well as a binary PDF document.`,
	Run: func(cmd *cobra.Command, args []string) {
		resume := marshall.LoadJsonFile(input)

		for _, fmt := range format {
			var fn func(string, marshall.Resume) string

			if fmt == "html" {
				fn = marshall.RenderHtml
			} else if fmt == "txt" {
				fn = marshall.RenderText
			} else {
				shared.Exit(1, "The only supported formats are `html` and `txt`.")
			}

			output = files.WithExtension(output, fmt)
			out := fn(template, resume)

			if files.Exists(output) && !overwrite {
				shared.Exit(1, "The output file exists, and the overwrite flag is not provided.")
			}

			os.WriteFile(output, []byte(out), 0644)
			// shared.Exit(0, "Saving the file to: %s", output)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&input, "input", "i", "./resume.json", "path to the json file")
	generateCmd.Flags().StringVarP(&output, "output", "o", "resume.html", "path to save the documents")
	generateCmd.Flags().StringVarP(&template, "template", "t", "", "path to the template(s)")
	rootCmd.MarkFlagRequired("template")
	generateCmd.Flags().StringSliceVarP(&format, "format", "f", []string{"html"}, "theme for output")
	generateCmd.Flags().BoolVarP(&overwrite, "overwrite", "y", false, "overwrite existing files")
}
