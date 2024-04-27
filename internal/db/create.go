package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // this register sqlite3 drivers
)

// CreateDb creates a new connection to sqlite3.
// connectionString is a standard sqlite3 connectionString string (ex. ":memory:", "./db.sqlite3", "./foo.db", etc).
func CreateDb(connectionString string) (*sql.DB, error) {
	const driverName = "sqlite3"

	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
