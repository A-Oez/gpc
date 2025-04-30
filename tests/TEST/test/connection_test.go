package test

import (
	dbFactory "TEST/db"
	"testing"

	"github.com/joho/godotenv"
)

func TestPostgresConnection(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatalf("Error loading environment variables %v", err)
	}

	dbType := "postgres"
	db, err := dbFactory.GetDatabaseContextFactory(dbType).GetConnection()

	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("Error pinging database: %v", err)
	}

	defer db.Close()
}