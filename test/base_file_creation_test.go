package test

import (
	"testing"

	bp "github.com/A-Oez/GoProjectCreator/internal/structures/base"
)


func TestBPFileCreation(t *testing.T) {
	projectNameFlag := "BP_TEST"
	
	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: false,
	}
	bp.CreateMainDirectory()
	bp.CreateDirectories()
	bp.CreateFiles()
}
