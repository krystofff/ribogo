package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// RNA codon table
var rnaCodon = map[string]string{
	"UUU": "F", "CUU": "L", "AUU": "I", "GUU": "V",
	"UUC": "F", "CUC": "L", "AUC": "I", "GUC": "V",
	"UUA": "L", "CUA": "L", "AUA": "I", "GUA": "V",
	"UUG": "L", "CUG": "L", "AUG": "M", "GUG": "V",
	"UCU": "S", "CCU": "P", "ACU": "T", "GCU": "A",
	"UCC": "S", "CCC": "P", "ACC": "T", "GCC": "A",
	"UCA": "S", "CCA": "P", "ACA": "T", "GCA": "A",
	"UCG": "S", "CCG": "P", "ACG": "T", "GCG": "A",
	"UAU": "Y", "CAU": "H", "AAU": "N", "GAU": "D",
	"UAC": "Y", "CAC": "H", "AAC": "N", "GAC": "D",
	"UAA": "STOP", "CAA": "Q", "AAA": "K", "GAA": "E",
	"UAG": "STOP", "CAG": "Q", "AAG": "K", "GAG": "E",
	"UGU": "C", "CGU": "R", "AGU": "S", "GGU": "G",
	"UGC": "C", "CGC": "R", "AGC": "S", "GGC": "G",
	"UGA": "STOP", "CGA": "R", "AGA": "R", "GGA": "G",
	"UGG": "W", "CGG": "R", "AGG": "R", "GGG": "G",
}

func main() {
	fmt.Println(" *************************** Running ribogo *************************** ")

	rnaStrand := transcribeDNAtoRNA()
	fmt.Println("  RNA STRAND: ", rnaStrand)

	transcribeRNAtoProtein(rnaStrand)
	fmt.Println(" ******************************* ribogo ******************************* ")
}

func transcribeDNAtoRNA() (rna string) {
	dna, err := ioutil.ReadFile("sample_dna.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("  DNA STRAND: ", string(dna))

	rna = strings.Replace(string(dna), "T", "U", -1)

	return rna
}

func transcribeRNAtoProtein(rna string) {
	proteinString := ""
	// Create protein strand
	for i := 0; i <= len(rna)-(3+len(rna)%3); i = i + 3 {

		if rnaCodon[rna[i:i+3]] == "STOP" {
			break
		}
		proteinString += rnaCodon[rna[i:i+3]]
	}

	fmt.Println("  PROTEIN STRAND: ", proteinString)
}
