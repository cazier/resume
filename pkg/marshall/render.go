package marshall

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	shared "github.com/cazier/resume/pkg/shared"
	themes "github.com/cazier/resume/pkg/themes"
	"github.com/playwright-community/playwright-go"

	"github.com/flosch/pongo2/v6"
)

func RenderText(theme themes.Theme, resume Resume) []byte {
	render, err := theme.Txt.Execute(pongo2.Context{
		"data":  resume,
		"funcs": funcs,
	})
	shared.HandleError(err)

	return []byte(render)
}

func RenderHtml(theme themes.Theme, resume Resume) []byte {
	render, err := theme.Html.Execute(pongo2.Context{
		"data":  resume,
		"funcs": funcs,
	})
	shared.HandleError(err)

	return []byte(render)
}

func RenderPdf(theme themes.Theme, resume Resume) []byte {
	dir, err := os.MkdirTemp("", "resume")
	shared.HandleError(err)

	err = os.WriteFile(filepath.Join(dir, "resume.html"), []byte(RenderHtml(theme, resume)), 0644)
	shared.HandleError(err)

	pw, err := playwright.Run(&playwright.RunOptions{Browsers: []string{"chromium"}})
	shared.HandleError(err)

	browser, err := pw.Chromium.Launch()
	shared.HandleError(err)

	page, err := browser.NewPage()
	shared.HandleError(err)

	_, err = page.Goto("file://"+filepath.Join(dir, "resume.html"), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	})
	shared.HandleError(err)

	_, err = page.PDF(playwright.PagePdfOptions{
		Path: playwright.String(filepath.Join(dir, "resume.pdf")),
	})
	shared.HandleError(err)

	pdf, err := os.ReadFile(filepath.Join(dir, "resume.pdf"))
	shared.HandleError(err)

	return pdf
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
