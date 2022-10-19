package files

import (
	"errors"
	"os"
	"path/filepath"

	shared "github.com/cazier/resume/pkg/shared"
)

func Exists(path string, is_file bool) bool {
	resp, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else if errors.Is(err, os.ErrPermission) {
		shared.Exit(1, "The destination path (%s) has bad permissions", path)
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
