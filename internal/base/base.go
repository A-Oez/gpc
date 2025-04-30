package base

import (
	"embed"
	"os/exec"

	"github.com/A-Oez/gpc/pkg/commands"
	"github.com/A-Oez/gpc/pkg/directories"
	"github.com/A-Oez/gpc/pkg/files"
)

//go:embed structure/main.go.txt
//go:embed structure/README.md.txt
var content embed.FS

type BaseProject struct{
	ProjectName string
	OpenEditor bool
}

func (bp *BaseProject) Execute(){
	bp.createMainDirectory()
	bp.createFiles()
	bp.useCommands()
}

func (bp *BaseProject) createMainDirectory(){
	directories.CreateDir([]string{bp.ProjectName})
}

func (bp *BaseProject) useCommands() {
	cmdGoMod := exec.Command("go", "mod", "init", bp.ProjectName)
	cmdGoMod.Dir = bp.ProjectName
	commands.ExecuteCommand(cmdGoMod)

	if bp.OpenEditor {
		cmdCode := exec.Command("code", ".")
		cmdCode.Dir = bp.ProjectName
		commands.ExecuteCommand(cmdCode)
	}
}

func (bp *BaseProject) createFiles(){
	filesArr := []files.File{
		{
			Path:    []string{bp.ProjectName, "main.go"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/main.go.txt")),
		},
		{
			Path:    []string{bp.ProjectName, "README.md"},
			Content: []byte(files.GetEmbeddedContent(content, "structure/README.md.txt")),
		},
	}

	for _, fileStr := range filesArr {
		files.CreateFiles(bp.ProjectName, string(fileStr.Content), fileStr.Path)
	}
}