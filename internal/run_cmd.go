package internal

import (
	"log"

	bp "github.com/A-Oez/GoProjectCreator/internal/config/base"
	"github.com/A-Oez/GoProjectCreator/internal/config/db"
)

func ExecuteCreation(bp bp.BaseProject, dbStr string) {
	bp.Execute()
	executeDB(bp.ProjectName, dbStr)
}

func executeDB(projectName string, dbStr string){
	if dbStr != ""{ 
		dbType, ok := db.GetDatabaseType(dbStr)
		if !ok {
			log.Fatalf("Given db input %s doesn't exist, no db structure created", dbStr)
		}
		dbType.Execute(projectName)
	}
}