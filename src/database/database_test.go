package database

import (
	"testing"
	"os"
)

func TestNewDatabase(t *testing.T) {
	databasePath := "test.db"
	database, close, err := NewDatabase(databasePath)

	if err != nil {
		t.Error(err)
	}

	database.SetBucket("main")

	defer close()

	if fileExists(databasePath) {
		dErr := os.Remove(databasePath)

		if dErr != nil {
			t.Error("Database can't be deleted")
		}
	} else {
		t.Error("Database file not set")
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
