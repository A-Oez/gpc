package utils

import (
	"fmt"
	"os"
)

func CreateDir(paths []string) {

	for _, path := range paths {
		err := os.Mkdir(path, 0755)

		if err != nil {
			fmt.Fprintln(os.Stderr, "use --dir to specify a project name / project already exists")
			os.Exit(1)
		}
	}
}