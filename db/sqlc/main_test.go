package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatalf("Cannot connect to the database: %v", err)
	}
	if err := testDB.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
