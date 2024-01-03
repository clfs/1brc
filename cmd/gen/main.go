package main

import (
	"flag"
	"log"
	"os"

	"github.com/clfs/obrc"
)

func main() {
	nFlag := flag.Int("n", 0, "number of measurements to generate")
	flag.Parse()

	if *nFlag == 0 {
		flag.Usage()
		return
	}

	if err := obrc.GenerateCSV(os.Stdout, *nFlag); err != nil {
		log.Fatal(err)
	}
}
