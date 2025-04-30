package test

import (
	"testing"

	bp "github.com/A-Oez/gpc/internal/base"
)

var(
	project_name = "TEST"
)


func TestBPFileCreation(t *testing.T) {	
	bp := bp.BaseProject{
		ProjectName: project_name, 
		OpenEditor: false,
	}
	
	bp.Execute()
}
