package db

import (
	"path/filepath"

	utils "github.com/A-Oez/GoProjectCreator/internal/utils"
)

type DBPostgres struct {
	ProjectName string
}

func (db *DBPostgres) CreateDirectories(){
	dbPath := filepath.Join(db.ProjectName, "db")
	service := filepath.Join(db.ProjectName, "service")
	database_types := filepath.Join(db.ProjectName, "db", "database_types")
	entity := filepath.Join(db.ProjectName, "db", "entity")
	repository := filepath.Join(db.ProjectName, "db", "repository")
	user := filepath.Join(db.ProjectName, "db", "repository", "user")
	postgres := filepath.Join(db.ProjectName, "db", "repository", "user", "postgres")
	subDir := []string{dbPath, service, database_types, entity, repository, user, postgres}
	utils.CreateDir(subDir)	
}

func (db *DBPostgres) CreateFiles(){
	envPath := []string{db.ProjectName, ".env"}
	envConfig := []string{"gocreate_config", "postgres", ".env"}

	dockerPath := []string{db.ProjectName, "docker-compose.yml"}
	dockerConfig := []string{"gocreate_config", "postgres", "docker-compose.yml"}

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

func (db *DBPostgres) UseCommand(){
	//go get github.com/joho/godotenv
	//go get github.com/lib/pq
}