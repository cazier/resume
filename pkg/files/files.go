package files

import (
	"os"
)

func Exists(file string) bool {
	resp, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !resp.IsDir()
}
