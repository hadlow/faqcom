package main

import (
	"testing"
	"os"
	"reflect"
	"io/ioutil"
	
	"github.com/hadlow/genomdb/src/types"
)

var contents = `database: "main.db"
host: "127.0.0.1"
port: 6969

shards:
    - id: 0
      name: "shard"
      host: "127.0.0.1"
      port: 6969

    - id: 1
      name: "shard"
      host: "127.0.0.1"
      port: 6868
`

func TestLoadConfig(t *testing.T) {
	// Create temporary YAML file
	file, err := ioutil.TempFile(os.TempDir(), "config.yml")

	if err != nil {
		t.Fatal("Error creating temporary file")
	}

	defer file.Close()
	defer os.Remove(file.Name())

	_, err = file.WriteString(contents)

	if err != nil {
		t.Fatal("Error writing to file")
	}

	// Load the YAML config
	config, err := loadConfig(file.Name())

	if err != nil {
		t.Fatalf("Error loading config: %v", err)
	}

	expected := Config{
		Database: "main.db",
		Host: "127.0.0.1",
		Port: 6969,
		Shards: []types.Shard{
			{
				Id: 0,
				Name: "shard",
				Host: "127.0.0.1",
				Port: 6969,
			},
			{
				Id: 1,
				Name: "shard",
				Host: "127.0.0.1",
				Port: 6868,
			},
		},
	}

	// Test that we get the expected value back
	if !reflect.DeepEqual(config, expected) {
		t.Fatal("Config not expected")
	}
}