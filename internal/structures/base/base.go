package base

import (
	"os/exec"

	"path/filepath"

	utils "github.com/A-Oez/GoProjectCreator/internal/utils"
)

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

	subDir := []string{cmdPath, internalPath}
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

func (bp *BaseProject) CreateBaseFiles(){
	readmePath := []string{bp.ProjectName, "README.md"}
	readmeConfig := []string{"gocreate_config", "README.md.txt"}

	files := []utils.File{
		{
			FilePath: readmePath,
			ConfigPath: readmeConfig,
		},
	}

	for _, file := range files{
		utils.CreateFiles(file.FilePath, file.ConfigPath)
	} 
}

func (bp *BaseProject) CreateMainFile(){
	mainPath := []string{bp.ProjectName, "main.go"}
	mainConfig := []string{"gocreate_config", "main.go.txt"}

	files := []utils.File{
		{
			FilePath: mainPath,
			ConfigPath: mainConfig,
		},
	}

	for _, file := range files{
		utils.CreateFiles(file.FilePath, file.ConfigPath)
	} 
}