package postgres

import (
	"embed"
	"os/exec"
	"path/filepath"

	utils "github.com/A-Oez/gpc/pkg/utils"
)

//go:embed structure/main.go.txt
//go:embed structure/docker-compose.yml
//go:embed structure/.gitignore.txt
//go:embed structure/env.txt
//go:embed structure/test/connection_test.go.txt
//go:embed structure/db/db_context.go.txt
//go:embed structure/db/types/postgres.go.txt
var content embed.FS


type DBPostgres struct {
	ProjectName string
	OpenEditor bool
}

func (db *DBPostgres) CreateDirectories(){
	dbPath := filepath.Join(db.ProjectName, "db")
	typesPath := filepath.Join(db.ProjectName, "db", "types")
	subDir := []string{dbPath, typesPath}
	utils.CreateDir(subDir)	
}

func (db *DBPostgres) CreateFiles(){
	files := []utils.File{
		{
			Path:    []string{db.ProjectName, ".env"},
			Content: []byte(utils.GetEmbeddedContent(content, "structure/env.txt")),
		},
		{
			Path:    []string{db.ProjectName, "docker-compose.yml"},
			Content: []byte(utils.GetEmbeddedContent(content, "structure/docker-compose.yml")),
		},
		{
			Path:    []string{db.ProjectName, ".gitignore"},
			Content: []byte(utils.GetEmbeddedContent(content, "structure/.gitignore.txt")),
		},
		{
			Path:    []string{db.ProjectName, "main.go"},
			Content: []byte(utils.GetEmbeddedContent(content, "structure/main.go.txt")),
		},
		{
			Path:    []string{db.ProjectName, "db", "db_context.go"},
			Content: []byte(utils.GetEmbeddedContent(content, "structure/db/db_context.go.txt")),
		},
		{
			Path:    []string{db.ProjectName, "db", "types", "postgres.go"},
			Content: []byte(utils.GetEmbeddedContent(content, "structure/db/types/postgres.go.txt")),
		},
		{
			Path:    []string{db.ProjectName, "test","connection_test.go"},
			Content: []byte(utils.GetEmbeddedContent(content, "structure/test/connection_test.go.txt")),
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