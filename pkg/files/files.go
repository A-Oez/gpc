package files

import (
	"embed"
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
		log.Fatalf("Failed to create file at %s: %v", path, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("Failed to write content to file at %s: %v", path, err)
	}
}

func GetEmbeddedContent(content embed.FS, fileName string) []byte {
	output, err := content.ReadFile(fileName)

	if err != nil {
		log.Fatalf("Failed to read embedded file %s: %v", fileName, err)
	}

	return output
}

func replaceContent(projectName string, content string) string {
	return strings.ReplaceAll(content, "XPROJECTNAMEX", projectName)
}