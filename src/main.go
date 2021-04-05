package main

// Local imports
import (
	"fmt"
	"log"
	"flag"
)

import (
	"github.com/hadlow/genomdb/src/database"
	"github.com/hadlow/genomdb/src/endpoints"
)

// Flags
var (
	pDBPath = flag.String("database", "", "Database path")
	pHTTPAddress = flag.String("address", "localhost", "Host HTTP address")
	pHTTPPort = flag.String("port", "6969", "Host HTTP port")
)

func validateFlags() {
	flag.Parse()

	if *pDBPath == "" {
		log.Fatal("Missing database location")
	}
}

func main() {
	validateFlags()

	database, close, err := database.NewDatabase(*pDBPath)
	database.SetBucket("main")

	if err != nil {
		log.Fatal("Error opening database")
	}

	config, err := loadConfig()

	if err != nil {
		log.Fatal(err)
	}

	ep := endpoints.New(database)

	fmt.Println("Starting server on: " + *pHTTPAddress + ":" + *pHTTPPort)
	log.Fatal(ep.Serve(*pHTTPAddress, *pHTTPPort))

	defer close()
}
