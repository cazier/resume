package themes

import (
	"embed"
	"errors"
	"io"
	iofs "io/fs"
	"path/filepath"
	"strings"

	shared "github.com/cazier/resume/pkg/shared"

	"github.com/flosch/pongo2/v6"
)

type Theme struct {
	Html    *pongo2.Template
	Txt     *pongo2.Template
	Builtin bool
}

const dir string = "templates"

var Builtins []string = []string{"handmade"}

//go:embed templates/*
var fs embed.FS

var templates pongo2.TemplateSet = *pongo2.NewSet("builtins", PongoTemplates{fs})

type PongoTemplates struct {
	vfs embed.FS
}

func (p PongoTemplates) Abs(base string, name string) string {
	var abs string = name

	if base != "" {
		abs = filepath.Join(filepath.Dir(base), name)
	}

	return abs
}

func (p PongoTemplates) Get(path string) (io.Reader, error) {
	return p.vfs.Open(path)
}

func validName(name string) (bool, []string) {
	files, err := fs.ReadDir(filepath.Join(dir, name))

	var types []string = []string{"html", "txt"}

	if errors.Is(err, iofs.ErrNotExist) && strings.Contains(err.Error(), "file does not exist") {
		return false, []string{}
	}

	for _, file := range files {
		for index, value := range types {
			if strings.Contains(file.Name(), value) {
				types = append(types[:index], types[index+1:]...)
			}
		}
	}

	return len(types) == 0, types
}

func load(name string) Theme {
	valid, missing := validName(name)
	if !valid {
		if len(missing) == 0 {
			shared.Exit(1, "The template name (%s) does not exist", name)
		}
		shared.Exit(1, "The template does not have the types: %s", strings.Join(missing, ", "))
	}

	html, err := templates.FromFile(filepath.Join(dir, name, "template.html.j2"))
	shared.HandleError(err)

	txt, err := templates.FromFile(filepath.Join(dir, name, "template.txt.j2"))
	shared.HandleError(err)

	return Theme{html, txt, true}
}

func FindThemeData(name string) Theme {
	for _, b := range Builtins {
		if strings.ToLower(name) == b {
			return load(b)
		}
	}

	html, err := pongo2.FromFile(filepath.Join(name, "template.html.go.j2"))
	shared.HandleError(err)

	txt, err := pongo2.FromFile(filepath.Join(name, "template.txt.go.j2"))
	shared.HandleError(err)

	return Theme{html, txt, false}
}
