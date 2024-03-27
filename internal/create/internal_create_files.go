package internal_create

import (
	"fmt"
	"os"
	"path/filepath"
)

func allfiles() {
	createFiles("main.go", true)

	if varReadme {
		createFiles("README.md", true)
	}
}

func createFiles(fileName string, config bool) {
	filePath := filepath.Join(varDirName, fileName)
	file, err := os.Create(filePath)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if config {
		content, err := getConfigFileContent(fileName)

		if err == nil {
			writeFileContent(file, content)
		}

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

func getConfigFileContent(fileName string) (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	executableDir := filepath.Dir(executablePath)

	configPath := filepath.Join(executableDir, "gocreate_config", "config_"+fileName+".txt")

	content, err := os.ReadFile(configPath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
