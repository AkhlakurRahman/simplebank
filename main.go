package main

import (
	"database/sql"
	"log"

	"github.com/akhlakurrahman/simplebank/api"
	db "github.com/akhlakurrahman/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatalf("Cannot connect to the database: %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
