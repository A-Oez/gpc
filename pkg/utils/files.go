package utils

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type File struct { 	
	Path    []string
	Content []byte
}

func CreateFiles(projectName string, content string, filePath []string) {
	content = replaceContent(projectName, content)
	path := filepath.Join(filePath...)

	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}


func GetEmbeddedContent(content embed.FS, fileName string) []byte {
	output, err := content.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return output
}

func replaceContent(projectName string, content string) string {
	return strings.ReplaceAll(content, "XPROJECTNAMEX", projectName)
}