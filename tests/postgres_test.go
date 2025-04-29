package test

import (
	"log"
	"testing"

	bp "github.com/A-Oez/gpc/internal/base"
	"github.com/A-Oez/gpc/internal/db"
)

func TestDBFileCreation(t *testing.T) {
	projectNameFlag := "DB_TEST"
	
	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: false,
	}

	bp.Execute()

	dbStr := "Postgres"
	dbType, ok := db.GetDatabaseType(dbStr)
	if !ok {
		log.Fatalf("Given db input %s doesn't exist, no db structure created", dbStr)
	}
	dbType.Execute(projectNameFlag)
}
