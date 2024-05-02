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

func (db *DatabaseProject) CreateFiles(){
	envPath := []string{db.ProjectName,"db",".env"}
	envConfig := []string{"gocreate_config", "db", "postgres", ".env"}

	dockerPath := []string{db.ProjectName,"db","docker-compose.yml"}
	dockerConfig := []string{"gocreate_config", "db", "postgres", "docker-compose.yml"}

	files := []utils.File{
		{
			FilePath: envPath,
			ConfigPath: envConfig,
		},
		{
			FilePath: dockerPath,
			ConfigPath: dockerConfig,
		},
	}

	for _, file := range files{
		utils.CreateFiles(file.FilePath, file.ConfigPath)
	} 
}