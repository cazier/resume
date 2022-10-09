package shared

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func Exit(code int, message string, args ...any) {
	fmt.Printf(message+"\n", args...)
	os.Exit(code)
}
