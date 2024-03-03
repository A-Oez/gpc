package internal_create

import (
	"fmt"
	"os"
	"path/filepath"
)

func allfiles() {
	createFiles("main.go", true)
	createFiles("README.md", true)
}

func createFiles(fileName string, config bool) {
	filePath := filepath.Join(varDirName, fileName)
	file, err := os.Create(filePath)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if config {
		content := getConfigFileContent(fileName)
		writeFileContent(file, content)
	}

	defer file.Close()
}

func writeFileContent(file *os.File, content string) {
	_, err := file.WriteString(content)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getConfigFileContent(fileName string) string {
	content, err := os.ReadFile("config/config_" + fileName + ".txt")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return string(content)
}
