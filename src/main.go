package main

// Local imports
import (
	"fmt"
	"io/ioutil"
	"log"
	"flag"
)

// External imports
import (
	bolt "go.etcd.io/bbolt"
)

// Flags
var (
	pFilename = flag.String("file", "", "File path")
	pDBLocation = flag.String("db", "", "Database path")
)

func validateFlags() {
	flag.Parse()

	if *pFilename == "" {
		log.Fatal("Missing filename")
	}

	if *pDBLocation == "" {
		log.Fatal("Missing database location")
	}
}

func main() {
	validateFlags()

	content, error := ioutil.ReadFile(*pFilename)

	if error != nil {
		log.Fatal(error)
	}

	db, err := bolt.Open(*pDBLocation, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println(content)
}
