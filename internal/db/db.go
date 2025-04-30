package db

import (
	pg "github.com/A-Oez/gpc/internal/db/postgres"
)

//db enumeration
type DatabaseType string

type DatabaseStructure interface {
	CreateFiles()
	UseCommand()
}

const (
	Postgres DatabaseType = "Postgres"
)

var SupportedDatabaseTypes = []DatabaseType{
	Postgres,
}

func (dbType DatabaseType) Execute(projectName string){
	db := dbFactory(projectName, dbType)
	db.CreateFiles()
	db.UseCommand()
}

func GetDatabaseType(dbType string) (DatabaseType, bool) {
	for _, entry := range SupportedDatabaseTypes {
		if string(entry) == dbType {
			return entry, true
		}
	}
	return "", false
}

func dbFactory(projectName string, dbType DatabaseType) DatabaseStructure {
	switch dbType {
	case Postgres:
		return &pg.DBPostgres{
			ProjectName: projectName,
		}
	default:
		return nil
	}
}