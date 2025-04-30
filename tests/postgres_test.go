package test

import (
	"log"
	"testing"

	bp "github.com/A-Oez/gpc/internal/base"
	"github.com/A-Oez/gpc/internal/db"
)

func TestDBFileCreation(t *testing.T) {	
	bp := bp.BaseProject{
		ProjectName: project_name, 
		OpenEditor: false,
	}

	bp.Execute()

	dbStr := "Postgres"
	dbType, ok := db.GetDatabaseType(dbStr)
	if !ok {
		log.Fatalf("Given db input %s doesn't exist, no db structure created", dbStr)
	}
	dbType.Execute(project_name)
}
