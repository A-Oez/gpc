package internal

import (
	pg "github.com/A-Oez/GoProjectCreator/internal/structures/postgres"
)

// simulate enumeration
type DatabaseType string

type DatabaseStructure interface {
	CreateDirectories()
	CreateFiles()
	UseCommand()
}

const (
	Postgres DatabaseType = "Postgres"
)

var SupportedDatabaseTypes = []DatabaseType{
	Postgres,
}

func ParseDatabaseType(dbType string) (DatabaseType, bool) {
	for _, entry := range SupportedDatabaseTypes {
		if string(entry) == dbType {
			return entry, true
		}
	}
	return "", false
}

func DatabaseServiceFactory(projectName string, dbType DatabaseType) DatabaseStructure {
	switch dbType {
	case Postgres:
		return &pg.DBPostgres{
			ProjectName: projectName,
		}
	default:
		return nil
	}
}