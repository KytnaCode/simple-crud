package db_test

import (
	"path"
	"testing"

	"github.com/kytnacode/simple-crud/internal/db"
	_ "github.com/mattn/go-sqlite3" // this registers sqlite3 driver
)

func TestCreate_memoryDatabaseShouldNotReturnAnError(t *testing.T) {
	t.Parallel()

	// Arrange
	connectionString := ":memory:"

	// Act
	_, err := db.CreateDb(connectionString)

	// Assert
	if err != nil {
		t.Error(err)
	}
}

func TestCreate_fileDbShouldNotReturnAnError(t *testing.T) {
	t.Parallel()

	// Arrange
	tmpDir := t.TempDir()
	connectionString := path.Join(tmpDir, "foo.db")

	// Act
	_, err := db.CreateDb(connectionString)

	// Assert
	if err != nil {
		t.Error(err)
	}
}
