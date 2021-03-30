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

type GetHandler struct {
	DB *database.Database
}

func (d *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := r.Form.Get("key")

	value, err := d.DB.Get(key)

	if err != nil {
		log.Fatal("Error getting value")
	}

	fmt.Println(value)
}

type SetHandler struct {
	DB *database.Database
}

func (d *SetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := r.Form.Get("key")
	value := r.Form.Get("data")

	err := d.DB.Set(key, []byte(value))

	if err != nil {
		log.Fatal("Error setting value")
	}

	fmt.Println("Key value set")
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
	database.SetBucket("main")

	if err != nil {
		log.Fatal("Error opening database")
	}

	http.Handle("/g", &GetHandler{DB: database})
	http.Handle("/s", &SetHandler{DB: database})

	fmt.Println("Starting server on: " + *pHTTPAddress + ":" + *pHTTPPort)
	log.Fatal(http.ListenAndServe(*pHTTPAddress + ":" + *pHTTPPort, nil))

	defer close()
}
