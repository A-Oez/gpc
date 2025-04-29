package test

import (
	"testing"

	bp "github.com/A-Oez/gpc/internal/base"
)


func TestBPFileCreation(t *testing.T) {
	projectNameFlag := "BP_TEST"
	
	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: false,
	}
	
	bp.Execute()
}
