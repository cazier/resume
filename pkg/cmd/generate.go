/*
Copyright © 2022 Brendan Cazier
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	files "github.com/cazier/resume/pkg/files"
	marshall "github.com/cazier/resume/pkg/marshall"
	shared "github.com/cazier/resume/pkg/shared"
	themes "github.com/cazier/resume/pkg/themes"

	"github.com/spf13/cobra"
)

var format []string
var template string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "create the output resume files",
	Long: `
Generate the generated resume files by filling a passed in template with the supplied resume data.
The output can contain plain text (.txt) files or websites (.html), as well as a binary PDF document
(PENDING).

If an output FILE is provided, the tool will output that exact filename, regardless of the formats 
generated. If an output DIRECTORY is provided, the tool will output the file(s) in that directory
(creating it, if needed), with the filename "resume.EXT" (where "EXT" is the proper extension).
If an output DIRECTORY is provided with a BUILTIN template, a subdirectory with the template name
will be used.

The built-in themes include: ` + strings.Join(themes.Builtins, ", "),
	Run: func(cmd *cobra.Command, args []string) {
		resume := marshall.LoadJsonFile(input)
		theme := themes.FindThemeData(template)

		for _, f := range format {
			var fn func(themes.Theme, marshall.Resume) []byte
			var fout string

			switch strings.ToLower(f) {
			case "html":
				fn = marshall.RenderHtml
			case "txt":
				fn = marshall.RenderText
			case "pdf":
				fn = marshall.RenderPdf
			default:
				shared.Exit(1, "The only supported formats are `html`, `txt`, and `pdf`.")
			}

			files.MakeDirectories(output, true)

			info, err := os.Stat(output)
			if !os.IsNotExist(err) {
				shared.HandleError(err)
			}

			if info.IsDir() {
				if theme.Builtin {
					fout = filepath.Join(output, template, fmt.Sprintf("resume.%s", f))
				} else {
					fout = filepath.Join(output, fmt.Sprintf("resume.%s", f))
				}
			} else {
				fout = output
			}

			files.MakeDirectories(fout, false)

			gen := fn(theme, resume)

			if files.Exists(fout, true) && !overwrite {
				shared.Exit(
					1,
					"The output (%s) exists, and the overwrite flag is not provided.",
					fout,
				)
			}

			err = os.WriteFile(fout, gen, 0644)
			shared.HandleError(err)

			fmt.Printf("Saved output to: %s\n", fout)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&input, "input", "i", "./resume.json", "path to the json file")
	generateCmd.Flags().StringVarP(&output, "output", "o", ".", "directory to save the documents")
	generateCmd.Flags().StringVarP(&template, "template", "t", "", "theme name or path to the template(s)")
	rootCmd.MarkFlagRequired("template")
	generateCmd.Flags().StringSliceVarP(&format, "format", "f", []string{"html"}, "file format to generate")
	generateCmd.Flags().BoolVarP(&overwrite, "overwrite", "y", false, "overwrite existing files")
}
