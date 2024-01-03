package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/clfs/obrc"
)

func main() {
	fileFlag := flag.String("f", "", "input file")
	flag.Parse()

	if *fileFlag == "" {
		flag.Usage()
		return
	}

	f, err := os.Open(*fileFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m, err := obrc.ComputeStats(f)
	if err != nil {
		log.Fatal(err)
	}

	// Sorting the results doesn't affect the final time much; skipped it.
	for k, v := range m {
		fmt.Printf("%s: %.1f/%.1f/%.1f\n", k, v.Min, v.Max, v.Mean())
	}
}
