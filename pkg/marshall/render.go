package marshall

import (
	"sort"
	"strings"
	"time"

	shared "github.com/cazier/resume/pkg/shared"
	themes "github.com/cazier/resume/pkg/themes"

	"github.com/flosch/pongo2/v6"
)

func RenderText(theme themes.Theme, resume Resume) string {
	render, err := theme.Txt.Execute(pongo2.Context{
		"data":  resume,
		"funcs": funcs,
	})
	shared.HandleError(err)

	return render
}

func RenderHtml(theme themes.Theme, resume Resume) string {
	render, err := theme.Html.Execute(pongo2.Context{
		"data":  resume,
		"funcs": funcs,
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

func Join(elements []string, with string) string {
	return strings.Join(elements, with)
}

var funcs map[string]any = map[string]any{
	"join": Join,
	"sort": Sort,
}
