package db

import (
	"database/sql"
)

// A unexported global variable should be fine for a simple project like this.
var database *sql.DB

func Init(connectionString string) error {
	newDb, err := CreateDb(connectionString)

	database = newDb

	return err
}

func Get() *sql.DB {
	return database
}
