package postgres

import (
	"embed"
	"os/exec"
	"path/filepath"

	"github.com/A-Oez/gpc/pkg/commands"
	"github.com/A-Oez/gpc/pkg/directories"
	"github.com/A-Oez/gpc/pkg/files"
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
	directories.CreateDir(subDir)	
}

func (db *DBPostgres) CreateFiles(){
	filesArr := []files.File{
		{
			Path:    []string{db.ProjectName, ".env"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/env.txt")),
		},
		{
			Path:    []string{db.ProjectName, "docker-compose.yml"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/docker-compose.yml")),
		},
		{
			Path:    []string{db.ProjectName, ".gitignore"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/.gitignore.txt")),
		},
		{
			Path:    []string{db.ProjectName, "main.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/main.go.txt")),
		},
		{
			Path:    []string{db.ProjectName, "db", "db_context.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/db/db_context.go.txt")),
		},
		{
			Path:    []string{db.ProjectName, "db", "types", "postgres.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/db/types/postgres.go.txt")),
		},
		{
			Path:    []string{db.ProjectName, "test","connection_test.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/test/connection_test.go.txt")),
		},
	}

	for _, fileStr := range filesArr {
		files.CreateFiles(db.ProjectName, string(fileStr.Content), fileStr.Path)
	}
}

func (db *DBPostgres) UseCommand(){
	cmdEnv := exec.Command("go", "get", "github.com/joho/godotenv")
	cmdEnv.Dir = db.ProjectName
	commands.ExecuteCommand(cmdEnv)

	cmdPQ := exec.Command("go", "get", "github.com/lib/pq")
	cmdPQ.Dir = db.ProjectName
	commands.ExecuteCommand(cmdPQ)
}