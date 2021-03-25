package main

// Local imports
import (
	"fmt"
	"io/ioutil"
	"log"
)

// External imports
import (
	"github.com/shenwei356/bwt"
)

func main() {
	content, error := ioutil.ReadFile("genomes/test.fasta")

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(bwt)
}
