package main

// Local imports
import (
	"fmt"
	"log"
	"flag"
	"strconv"
)

import (
	"github.com/hadlow/genomdb/src/database"
	"github.com/hadlow/genomdb/src/endpoints"
)

// Flags
var (
	pConfigPath = flag.String("config", "./config.yml", "Path to the config YAML file")
)

func validateFlags() {
	flag.Parse()
}

func main() {
	validateFlags()

	config, err := loadConfig(*pConfigPath)

	if err != nil {
		log.Fatal(err)
	}

	database, close, err := database.NewDatabase(config.Database)
	database.SetBucket("main")

	if err != nil {
		log.Fatal("Error opening database")
	}

	ep := endpoints.New(database)

	fmt.Println("Starting server on: " + config.Host + ":" + strconv.Itoa(config.Port))
	log.Fatal(ep.Serve(config.Host, config.Port))

	defer close()
}
