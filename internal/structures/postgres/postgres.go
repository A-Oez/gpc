package postgres

import (
	"embed"
	"os/exec"
	"path/filepath"

	utils "github.com/A-Oez/GoProjectCreator/pkg/utils"
)

//go:embed config/main.go.txt
//go:embed config/docker-compose.yml
//go:embed config/.gitignore.txt
//go:embed config/env.txt
//go:embed config/test/connection_test.go.txt
//go:embed config/db/db_context.go.txt
//go:embed config/db/types/postgres.go.txt
var content embed.FS


type DBPostgres struct {
	ProjectName string
}

func (db *DBPostgres) CreateDirectories(){
	dbPath := filepath.Join(db.ProjectName, "db")
	typesPath := filepath.Join(db.ProjectName, "db", "types")
	testPath := filepath.Join(db.ProjectName, "test")
	subDir := []string{dbPath, typesPath, testPath}
	utils.CreateDir(subDir)	
}

func (db *DBPostgres) CreateFiles(){
	files := []utils.File{
		{
			Path:    []string{db.ProjectName, ".env"},
			Content: []byte(utils.GetEmbeddedContent(content, "config/env.txt")),
		},
		{
			Path:    []string{db.ProjectName, "docker-compose.yml"},
			Content: []byte(utils.GetEmbeddedContent(content, "config/docker-compose.yml")),
		},
		{
			Path:    []string{db.ProjectName, ".gitignore"},
			Content: []byte(utils.GetEmbeddedContent(content, "config/.gitignore.txt")),
		},
		{
			Path:    []string{db.ProjectName, "main.go"},
			Content: []byte(utils.GetEmbeddedContent(content, "config/main.go.txt")),
		},
		{
			Path:    []string{db.ProjectName, "db", "db_context.go"},
			Content: []byte(utils.GetEmbeddedContent(content, "config/db/db_context.go.txt")),
		},
		{
			Path:    []string{db.ProjectName, "db", "types", "postgres.go"},
			Content: []byte(utils.GetEmbeddedContent(content, "config/db/types/postgres.go.txt")),
		},
		{
			Path:    []string{db.ProjectName, "test","connection_test.go"},
			Content: []byte(utils.GetEmbeddedContent(content, "config/test/connection_test.go.txt")),
		},
	}

	for _, fileStr := range files {
		utils.CreateFiles(db.ProjectName, string(fileStr.Content), fileStr.Path)
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