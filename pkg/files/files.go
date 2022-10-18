package files

import (
	"os"
	"path/filepath"

	shared "github.com/cazier/resume/pkg/shared"
)

func Exists(path string, is_file bool) bool {
	resp, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return is_file != resp.IsDir()
}

func MakeDirectories(path string, include bool) {
	var fpath string

	if include && filepath.Ext(path) == "" {
		fpath = path
	} else {
		fpath = filepath.Dir(path)
	}

	err := os.MkdirAll(fpath, 0775)
	shared.HandleError(err)
}
