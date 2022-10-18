package shared

import (
	"fmt"
)

const major int = 1
const minor int = 1
const patch int = 1

func GetVersion() string {
	return fmt.Sprintf("v%d.%d.%d\n", major, minor, patch)
}
