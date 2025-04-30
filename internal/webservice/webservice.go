package webservice

import (
	"embed"
	"path/filepath"

	"github.com/A-Oez/gpc/pkg/directories"
	"github.com/A-Oez/gpc/pkg/files"
)

//go:embed structure/db/migrations/user_up.sql
//go:embed structure/docker-compose.yml
//go:embed structure/Dockerfile
//go:embed structure/env.txt
//go:embed structure/main.go.txt
//go:embed structure/internal/server/middleware/middleware.go.txt
//go:embed structure/internal/server/router.go.txt
//go:embed structure/internal/server/server.go.txt
//go:embed structure/internal/service/user/entity/entity.go.txt
//go:embed structure/internal/service/user/handler/handler.go.txt
//go:embed structure/internal/service/user/repo/repo.go.txt
//go:embed structure/internal/service/user/errors.go.txt
//go:embed structure/internal/service/user/user.go.txt
var content embed.FS

type Webservice struct{
	ProjectName string
}

func (ws *Webservice) Execute(){
	ws.createDirectories()
	ws.createFiles()
}

func (ws *Webservice) createDirectories() {
	migrationsPath := filepath.Join(ws.ProjectName, "db", "migrations")
	serverPath := filepath.Join(ws.ProjectName, "internal", "server")
	middlewarePath := filepath.Join(ws.ProjectName, "internal", "server", "middleware")
	servicePath := filepath.Join(ws.ProjectName, "internal", "service")
	userPath := filepath.Join(ws.ProjectName, "internal", "service", "user")
	handlerPath := filepath.Join(ws.ProjectName, "internal", "service", "user", "handler")
	entityPath := filepath.Join(ws.ProjectName, "internal", "service", "user", "entity")
	repoPath := filepath.Join(ws.ProjectName, "internal", "service", "user", "repo")
	
	subDir := []string{migrationsPath, serverPath, middlewarePath, servicePath, userPath, handlerPath, entityPath, repoPath}
	directories.CreateDir(subDir)	
}

func (ws *Webservice) createFiles(){
	filesArr := []files.File{
		{
			Path:    []string{ws.ProjectName, "main.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/main.go.txt")),
		},
		{
			Path:    []string{ws.ProjectName, ".env"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/env.txt")),
		},
		{
			Path:    []string{ws.ProjectName, "Dockerfile"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/Dockerfile")),
		},
		{
			Path:    []string{ws.ProjectName, "docker-compose.yml"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/docker-compose.yml")),
		},
		{
			Path:    []string{ws.ProjectName, "db/migrations/user_up.sql"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/db/migrations/user_up.sql")),
		},
		{
			Path:    []string{ws.ProjectName, "internal/server/router.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/internal/server/router.go.txt")),
		},
		{
			Path:    []string{ws.ProjectName, "internal/server/server.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/internal/server/server.go.txt")),
		},
		{
			Path:    []string{ws.ProjectName, "internal/server/middleware/middleware.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/internal/server/middleware/middleware.go.txt")),
		},
		{
			Path:    []string{ws.ProjectName, "internal/service/user/user.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/internal/service/user/user.go.txt")),
		},
		{
			Path:    []string{ws.ProjectName, "internal/service/user/errors.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/internal/service/user/errors.go.txt")),
		},
		{
			Path:    []string{ws.ProjectName, "internal/service/user/entity/entity.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/internal/service/user/entity/entity.go.txt")),
		},
		{
			Path:    []string{ws.ProjectName, "internal/service/user/handler/handler.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/internal/service/user/handler/handler.go.txt")),
		},
		{
			Path:    []string{ws.ProjectName, "internal/service/user/repo/repo.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/internal/service/user/repo/repo.go.txt")),
		},
	}

	for _, fileStr := range filesArr {
		files.CreateFiles(ws.ProjectName, string(fileStr.Content), fileStr.Path)
	}
}