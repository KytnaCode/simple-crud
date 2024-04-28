package db_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/kytnacode/simple-crud/internal/db"
)

func getANewDatabase(t *testing.T, connectionString string) *sql.DB {
	t.Helper()

	err := db.Init(connectionString) // init the database singleton
	if err != nil {
		t.Fatal(err)
	}

	return db.Get() // get the new created database
}

func TestPopulateDb_shouldCreatePostsTable(t *testing.T) {
	t.Parallel()

	// Arrange
	connectionString := ":memory:" // use an in-memory database
	tableName := "posts"

	database := getANewDatabase(t, connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Act
	err := db.PopulateDb(ctx, database)
	if err != nil {
		t.Fatal(err)
	}

	res, err := database.QueryContext(
		ctx,
		"SELECT name FROM sqlite_master WHERE type='table' AND name=?;",
		tableName,
	)

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	if !res.Next() { // if `res.Next()` is false, is because none tables were found
		t.Errorf("there is not a table named %v", tableName)
	}
}

func TestPopulateDb_shouldBePostsInPostsTable(t *testing.T) {
	t.Parallel()

	// Arrange
	connectionString := ":memory:"
	tableName := "posts"

	database := getANewDatabase(t, connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Act
	err := db.PopulateDb(ctx, database)
	if err != nil {
		t.Fatal(err)
	}

	res, err := database.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s;", tableName))

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	if !res.Next() {
		t.Error("there are not posts in the database")
	}
}
