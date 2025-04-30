package directories

import (
	"log"
	"os"
)

func CreateDir(paths []string) {
	for _, path := range paths {
		if err := os.MkdirAll(path, 0755); err != nil {
			log.Fatalf("Failed to create directory '%s': %v", path, err)
		}
	}
}