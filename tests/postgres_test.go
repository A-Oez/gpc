package test

import (
	"log"
	"testing"

	"github.com/A-Oez/gpc/internal/db"
)

func TestDBFileCreation(t *testing.T) {	
	dbStr := "Postgres"
	dbType, ok := db.GetDatabaseType(dbStr)
	if !ok {
		log.Fatalf("Given db input %s doesn't exist, no db structure created", dbStr)
	}
	dbType.Execute(project_name)
}
