package main

// Local imports
import (
	"fmt"
	"log"
	"flag"
	"net/http"
)

import (
	"github.com/hadlow/genomdb/src/database"
)

// Flags
var (
	pDBPath = flag.String("database", "", "Database path")
	pHTTPAddress = flag.String("address", "localhost", "Host HTTP address")
	pHTTPPort = flag.String("port", "6969", "Host HTTP port")
)

type HTTPHandler struct {
	DB *database.Database
}

func (d *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := r.Form.Get("key")

	value, err := d.DB.Get(key)

	if err != nil {
		log.Fatal("error getting value")
	}

	fmt.Println(value)
}

func set(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "set")
}

func validateFlags() {
	flag.Parse()

	if *pDBPath == "" {
		log.Fatal("Missing database location")
	}
}

func main() {
	validateFlags()

	database, close, err := database.New(*pDBPath)

	if err != nil {
		log.Fatal("Error opening database")
	}

	http.Handle("/g", &HTTPHandler{DB: database})
	http.HandleFunc("/s", set)

	fmt.Println("Starting server on: " + *pHTTPAddress + ":" + *pHTTPPort)
	log.Fatal(http.ListenAndServe(*pHTTPAddress + ":" + *pHTTPPort, nil))

	defer close()
}
