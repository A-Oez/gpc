package db

// simulate enumeration
type DatabaseType string

const (
	Postgres DatabaseType = "Postgres"
)

var SupportedDatabaseTypes = []DatabaseType{
	Postgres,
}

type DatabaseServices interface {
	CreateDirectories()
	CreateFiles()
}

func ParseDatabaseType(dbType string) (DatabaseType, bool) {
	for _, entry := range SupportedDatabaseTypes {
		if string(entry) == dbType {
			return entry, true
		}
	}
	return "", false
}

func DatabaseServiceFactory(projectName string, dbType DatabaseType) DatabaseServices {
	switch dbType {
	case Postgres:
		return &DBPostgres{
			ProjectName: projectName,
		}
	default:
		return nil
	}
}