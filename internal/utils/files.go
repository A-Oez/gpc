package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteFileContent(file *os.File, content string) {
	_, err := file.WriteString(content)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func GetConfigFileContent(filePath ...string) (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	executableDir := filepath.Dir(executablePath)

	paths := append([]string{executableDir}, filePath...)
	configPath := filepath.Join(paths...)

	content, err := os.ReadFile(configPath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
