package errors

import (
	"fmt"
	"os"
)

func ErrAndExit(err error) {
	fmt.Printf("%s", err)
	os.Exit(1)
}
