package internal_create

import (
	"fmt"
	"os"
	"path/filepath"
)

func mainDir() {
	createDir([]string{varDirName})
}

func subDir() {

	pkgPath := filepath.Join(varDirName, "pkg")
	cmdPath := filepath.Join(varDirName, "cmd")
	internalPath := filepath.Join(varDirName, "internal")
	configPath := filepath.Join(varDirName, "config")

	paths := []string{pkgPath, cmdPath, internalPath, configPath}

	createDir(paths)
}

func createDir(paths []string) {

	for _, path := range paths {
		err := os.Mkdir(path, 0755)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
