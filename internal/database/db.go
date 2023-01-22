package database

import (
	"database/sql"
	"log"
)

func Init(driver string, dbName string) *sql.DB {
	db, err := sql.Open(driver, "database"+string(dbName))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	return db
}
