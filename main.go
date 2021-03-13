package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func bwt(sequence []byte) string {
	return ""
}

func main() {
	content, error := ioutil.ReadFile("genomes/test.fasta")

	if error != nil {
		log.Fatal(error)
	}

	// BWT
	var bwt string = bwt(content)

	// Artithmetic coding

	fmt.Println(bwt)
}
