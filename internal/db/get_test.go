package db_test

import (
	"path"
	"testing"

	"github.com/kytnacode/simple-crud/internal/db"
)

func TestGet_initShouldNotReturnAnErrorWithAMemoryDatabase(t *testing.T) {
	t.Parallel()

	// Arrange
	connectionString := ":memory:"

	// Act
	err := db.Init(connectionString)

	// Assert
	if err != nil {
		t.Error(err)
	}
}

func TestGet_initShouldNotReturnanErrorWithAFileDatabase(t *testing.T) {
	t.Parallel()

	// Arrange
	tmpDir := t.TempDir()
	connectionString := path.Join(tmpDir, "foo.db")

	// Act
	err := db.Init(connectionString)

	// Assert
	if err != nil {
		t.Error(err)
	}
}
