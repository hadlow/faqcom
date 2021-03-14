package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

import "github.com/shenwei356/bwt"

func main() {
	content, error := ioutil.ReadFile("genomes/test.fasta")

	if error != nil {
		log.Fatal(error)
	}

	// BWT
	bwt, error := bwt(content)

	// Artithmetic coding

	fmt.Println(bwt)
}
