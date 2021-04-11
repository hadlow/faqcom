package database

import (
	"testing"
	"os"
	"bytes"
	"io/ioutil"
)

func TestGetSet(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "test.db")

	if err != nil {
		t.Fatal("Error creating database file")
	}

	defer file.Close()
	defer os.Remove(file.Name())

	database, close, err := NewDatabase(file.Name())
	database.SetBucket("main")

	if err != nil {
		t.Fatal("Error opening database")
	}

	defer close()

	err = database.Set("test", []byte("value"))

	if err != nil {
		t.Fatal("Error setting database key")
	}

	// Test get
	value, err := database.Get("test")

	if err != nil {
		t.Fatal("Could not get key")
	}

	if !bytes.Equal(value, []byte("value")) {
		t.Fatalf("Wrong key from database")
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}