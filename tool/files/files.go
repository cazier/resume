package files

import (
	"os"
	"path/filepath"
	"strings"
)

func Exists(file string) bool {
	resp, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !resp.IsDir()
}

func StripExtensions(fp string) (directory, base string) {
	directory = filepath.Dir(fp)
	file := filepath.Base(fp)

	return directory, strings.Split(file, ".")[0]
}
