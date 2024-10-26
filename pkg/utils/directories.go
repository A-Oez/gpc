package utils

import (
	"fmt"
	"os"
)

func CreateDir(paths []string) {

	for _, path := range paths {
		err := os.Mkdir(path, 0755)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}