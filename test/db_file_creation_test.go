package test

import (
	"testing"

	"github.com/A-Oez/GoProjectCreator/internal"
	bp "github.com/A-Oez/GoProjectCreator/internal/config/base"
)

func TestDBFileCreation(t *testing.T) {
	projectNameFlag := "DB_TEST"
	
	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: true,
	}

	internal.ExecuteCreation(bp, "asd")
}
