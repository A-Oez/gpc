package internal

import (
	"log"

	bp "github.com/A-Oez/gpc/internal/base"
	"github.com/A-Oez/gpc/internal/db"
	"github.com/A-Oez/gpc/internal/webservice"
)

func ExecuteCreation(bp bp.BaseProject, dbStr string, webserviceFlag bool) {
	bp.Execute() //creates base project structure, mandatory for all type of structures
	executeDB(bp.ProjectName, dbStr)

	if(webserviceFlag) {
		ws := webservice.Webservice{bp.ProjectName}
		ws.Execute()
	}
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