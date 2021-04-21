package main

// Local imports
import (
	"fmt"
	"log"
	"flag"
	"strconv"
	
	"github.com/hadlow/genomdb/src/database"
	"github.com/hadlow/genomdb/src/endpoints"
)

// Flags
var (
	pConfigPath = flag.String("config", "./config.yml", "Path to the config YAML file")
	pShard = flag.Int("shard", -1, "Shard number")
)

func validateFlags() {
	flag.Parse()

	if *pShard == -1 {
		log.Fatal("No shard number used")
	}
}

func main() {
	// Validate the command line flags
	validateFlags()

	// Load the config file YAML
	config, err := loadConfig(*pConfigPath)

	if err != nil {
		log.Fatal(err)
	}

	// Get the database from whatever file is in the config
	database, close, err := database.NewDatabase(config.Database)

	// Set the bucket (TODO - user specify bucket name)
	database.SetBucket("main")

	if err != nil {
		log.Fatal("Error opening database")
	}

	// Load out endpoints
	ep := endpoints.New(database, *pShard, config.Shards)

	// Start the server
	fmt.Println("Starting server on: " + config.Host + ":" + strconv.Itoa(config.Port))
	log.Fatal(ep.Serve(config.Host, config.Port))

	defer close()
}