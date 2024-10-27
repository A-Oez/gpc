package test

import (
	"testing"

	"github.com/A-Oez/gpc/internal"
	bp "github.com/A-Oez/gpc/internal/config/base"
)

func TestDBFileCreation(t *testing.T) {
	projectNameFlag := "DB_TEST"
	
	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: true,
	}

	internal.ExecuteCreation(bp, "asd")
}
