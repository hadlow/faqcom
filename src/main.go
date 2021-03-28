package main

// Local imports
import (
	"fmt"
	"log"
	"flag"
	"net/http"
)

// External imports
import (
	bolt "go.etcd.io/bbolt"
)

// Flags
var (
	pDBLocation = flag.String("database", "", "Database path")
	pHTTPAddress = flag.String("address", "localhost", "Host HTTP address")
	pHTTPPort = flag.String("port", "6969", "Host HTTP port")
)

func validateFlags() {
	flag.Parse()

	if *pDBLocation == "" {
		log.Fatal("Missing database location")
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get")
}

func set(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "set")
}

func main() {
	validateFlags()

	db, err := bolt.Open(*pDBLocation, 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	http.HandleFunc("/g", get)
	http.HandleFunc("/s", set)

	fmt.Println("Starting server on: " + *pHTTPAddress + ":" + *pHTTPPort)
	log.Fatal(http.ListenAndServe(*pHTTPAddress + ":" + *pHTTPPort, nil))
}
