package base_project

import (
	"fmt"
	"os"
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

func (bp *BaseProject) CreateFiles(){
	//file creation path != file config path
	mainGoFile := []string{bp.ProjectName, "main.go"}
	path := filepath.Join(mainGoFile...)

	file, err := os.Create(path)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	content, err := utils.GetConfigFileContent(mainGoFile...)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	utils.WriteFileContent(file, content)
	defer file.Close()
}