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
)

func main() {
	pfilename := flag.String("file", "", "File path")
	flag.Parse()

	content, error := ioutil.ReadFile(*pfilename)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(content)
}
