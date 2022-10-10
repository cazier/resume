package marshall

import (
	"path/filepath"
	"sort"
	"strings"
	"time"

	shared "tool/main/shared"

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

func parseDate(date string) time.Time {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		t, err = time.Parse("2006-01", date)
		shared.HandleError(err)
	}
	return t
}

func Sort(in []Work, reverse bool) []Work {
	parse := func(work Work) time.Time { return parseDate(work.StartDate) }

	out := make([]Work, len(in))
	copy(out, in)

	sort.Slice(out, func(i, j int) bool { return reverse && parse(out[i]).After(parse(out[j])) })

	return out
}

func DateFormat(in string) string {
	return parseDate(in).Format("Jan 2006")
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
	"date": DateFormat,
}
