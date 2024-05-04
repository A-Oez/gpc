package db

import (
	"os/exec"
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

	dbContextPath := []string{db.ProjectName, "db","db_context.go"}
	dbContextConfig := []string{"gocreate_config", "postgres", "db", "db_context.go.txt"}
	
	dbTypePath := []string{db.ProjectName, "db", "database_types", "postgres.go"}
	dbTypeConfig := []string{"gocreate_config", "postgres", "db", "database_types", "postgres.go.txt"}

	userEntityPath := []string{db.ProjectName, "db", "entity", "user.go"}
	userEntityConfig := []string{"gocreate_config", "postgres", "db", "entity", "user.go.txt"}

	repositoryPath := []string{db.ProjectName, "db", "repository", "repository.go"}
	repositoryConfig := []string{"gocreate_config", "postgres", "db", "repository", "repository.go.txt"}
	
	userRepositoryFactoryPath := []string{db.ProjectName, "db", "repository", "user", "user_repository_factory.go"}
	userRepositoryFactoryConfig := []string{"gocreate_config", "postgres", "db", "repository", "user", "user_repository_factory.go.txt"}
	
	userRepositoryPath := []string{db.ProjectName, "db", "repository", "user", "postgres", "user_repository.go"}
	userRepositoryConfig := []string{"gocreate_config", "postgres", "db", "repository", "user", "postgres", "user_repository.go.txt"}

	userServicePath := []string{db.ProjectName, "service", "user_service.go"}
	userServiceConfig := []string{"gocreate_config", "postgres", "service", "user_service.go.txt"}

	files := []utils.File{
		{
			FilePath: envPath,
			ConfigPath: envConfig,
		},
		{
			FilePath: dockerPath,
			ConfigPath: dockerConfig,
		},
		{
			FilePath: dbContextPath,
			ConfigPath: dbContextConfig,
		},
		{
			FilePath: dbTypePath,
			ConfigPath: dbTypeConfig,
		},
		{
			FilePath: userEntityPath,
			ConfigPath: userEntityConfig,
		},
		{
			FilePath: repositoryPath,
			ConfigPath: repositoryConfig,
		},
		{
			FilePath: userRepositoryFactoryPath,
			ConfigPath: userRepositoryFactoryConfig,
		},
		{
			FilePath: userRepositoryPath,
			ConfigPath: userRepositoryConfig,
		},
		{
			FilePath: userServicePath,
			ConfigPath: userServiceConfig,
		},
	}

	for _, file := range files{
		utils.CreateFiles(file.FilePath, file.ConfigPath)
	} 
}

func (db *DBPostgres) UseCommand(){
	cmdEnv := exec.Command("go", "get", "github.com/joho/godotenv")
	cmdEnv.Dir = db.ProjectName
	utils.ExecuteCommand(cmdEnv)

	cmdPQ := exec.Command("go", "get", "github.com/lib/pq")
	cmdPQ.Dir = db.ProjectName
	utils.ExecuteCommand(cmdPQ)
}