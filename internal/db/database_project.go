package db

import (
	"path/filepath"

	utils "github.com/A-Oez/GoProjectCreator/internal/utils"
)

type DatabaseProject struct {
	ProjectName string
}

func (db *DatabaseProject) CreateDirectories(){
	dbPath := filepath.Join(db.ProjectName, "db")
	subDir := []string{dbPath}
	utils.CreateDir(subDir)	
}