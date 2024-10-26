package test

import (
	"testing"

	dbFactory "github.com/A-Oez/GoProjectCreator/internal"
	bp "github.com/A-Oez/GoProjectCreator/internal/structures/base"
)

func TestDBFileCreation(t *testing.T) {
	databaseFlag := "Postgres"
	projectNameFlag := "DB_TEST"
	
	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: false,
	}
	bp.CreateMainDirectory()

	if dbType, isValid := dbFactory.ParseDatabaseType(databaseFlag); isValid{  
		db := dbFactory.DatabaseServiceFactory(projectNameFlag, dbType)
		db.CreateDirectories()
		db.CreateFiles()
		db.UseCommand()
	} 
}
