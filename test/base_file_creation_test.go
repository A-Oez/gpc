package test

import (
	"testing"

	"github.com/A-Oez/GoProjectCreator/internal"
	bp "github.com/A-Oez/GoProjectCreator/internal/config/base"
)


func TestBPFileCreation(t *testing.T) {
	projectNameFlag := "BP_TEST"
	
	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: true,
	}
	
	internal.RunProject(bp, "")
}
