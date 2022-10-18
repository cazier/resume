package marshall

import (
	"path/filepath"
	"sort"
	"strings"
	"time"

	shared "github.com/cazier/resume/pkg/shared"

	"github.com/flosch/pongo2/v6"
)

func RenderText(path string, resume Resume) string {
	path = filepath.Join(path, "template.txt.go.j2")

	tpl, err := pongo2.FromFile(path)
	shared.HandleError(err)

	render, err := tpl.Execute(pongo2.Context{
		"data":  resume,
		"funcs": Funcs,
	})
	shared.HandleError(err)

	return render
}

func RenderHtml(path string, resume Resume) string {
	path = filepath.Join(path, "template.html.go.j2")

	tpl, err := pongo2.FromFile(path)
	shared.HandleError(err)

	render, err := tpl.Execute(pongo2.Context{
		"data":  resume,
		"funcs": Funcs,
	})
	shared.HandleError(err)

	return render
}

func Sort(in []Work, reverse bool) []Work {
	parse := func(work Work) time.Time { return time.Time(work.StartDate) }

	out := make([]Work, len(in))
	copy(out, in)

	sort.Slice(out, func(i, j int) bool { return reverse && parse(out[i]).After(parse(out[j])) })

	return out
}

func ZipLanguages(in []Language) [][]string {
	output := make([][]string, 2)

	for index := range output {
		output[index] = make([]string, len(in))
	}

	for _, language := range in {
		output[0] = append(output[0], language.Language)
		output[1] = append(output[1], language.Fluency)
	}

	return output
}

var Funcs map[string]any = map[string]any{
	"join": func(elements []string, with string) string {
		return strings.Join(elements, with)
	},
	"sort": Sort,
}
