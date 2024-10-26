package base

import (
	"embed"
	"os/exec"

	"path/filepath"

	utils "github.com/A-Oez/GoProjectCreator/pkg/utils"
)

//go:embed config/main.go.txt
//go:embed config/README.md.txt
var content embed.FS

type BaseProject struct{
	ProjectName string
	OpenEditor bool
}

func (bp *BaseProject) CreateMainDirectory(){
	utils.CreateDir([]string{bp.ProjectName})
}

func (bp *BaseProject) CreateDirectories(){
	cmdPath := filepath.Join(bp.ProjectName, "cmd")
	internalPath := filepath.Join(bp.ProjectName, "internal")
	pkgPath := filepath.Join(bp.ProjectName, "pkg")
	testPath := filepath.Join(bp.ProjectName, "test")

	subDir := []string{cmdPath, internalPath, pkgPath, testPath}
	utils.CreateDir(subDir)	
}

func (bp *BaseProject) UseCommands() {
	cmdGoMod := exec.Command("go", "mod", "init", bp.ProjectName)
	cmdGoMod.Dir = bp.ProjectName
	utils.ExecuteCommand(cmdGoMod)

	if bp.OpenEditor {
		cmdCode := exec.Command("code", ".")
		cmdCode.Dir = bp.ProjectName
		utils.ExecuteCommand(cmdCode)
	}
}

func (bp *BaseProject) CreateFiles(){
	files := []utils.File{
		{
			Path:    []string{bp.ProjectName, "main.go"},
			Content: []byte(utils.GetEmbeddedContent(content, "config/main.go.txt")),
		},
		{
			Path:    []string{bp.ProjectName, "README.md"},
			Content: []byte(utils.GetEmbeddedContent(content, "config/README.md.txt")),
		},
	}

	for _, fileStr := range files {
		utils.CreateFiles(bp.ProjectName, string(fileStr.Content), fileStr.Path)
	}
}