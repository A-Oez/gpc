package internal_create

//TODO: methode schreiben, die inhalte in files übergben können -> json configs in ressource hinterlegen

import (
	"fmt"
	"os"
	"path/filepath"
)

func allfiles() {
	mainGo()
	readme()
}

func mainGo() {
	createFiles("main.go")
}

func readme() {
	if varReadme {
		createFiles("README.md")
	}
}

func createFiles(fileName string) {
	filePath := filepath.Join(varDirName, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()
}
