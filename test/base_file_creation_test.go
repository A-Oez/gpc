package test

import (
	"testing"

	"github.com/A-Oez/gpc/internal"
	bp "github.com/A-Oez/gpc/internal/config/base"
)


func TestBPFileCreation(t *testing.T) {
	projectNameFlag := "BP_TEST"
	
	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: true,
	}
	
	internal.ExecuteCreation(bp, "")
}
