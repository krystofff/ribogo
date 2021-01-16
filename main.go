package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fmt.Println(" *** Running ribogo *** ")

	rnaStrand := transcribeDNAtoRNA()
	fmt.Println("RNA STRAND: ", rnaStrand)
}

func transcribeDNAtoRNA() (rna string) {
	dna, err := ioutil.ReadFile("sample_dna.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DNA STRAND: ", string(dna))

	rna = strings.Replace(string(dna), "T", "U", -1)

	return rna
}
