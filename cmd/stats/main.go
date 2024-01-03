package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/clfs/obrc"
)

func main() {
	var (
		inputFlag      = flag.String("input", "", "input file")
		cpuprofileFlag = flag.String("cpuprofile", "", "write cpu profile to file")
	)
	flag.Parse()

	if *inputFlag == "" {
		flag.Usage()
		return
	}

	if *cpuprofileFlag != "" {
		f, err := os.Create(*cpuprofileFlag)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	f, err := os.Open(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m, err := obrc.TakeRecordings(obrc.NewReader(f))
	if err != nil {
		log.Fatal(err)
	}

	// Sorting the results doesn't affect the final time much; skipped it.
	for k, v := range m {
		fmt.Printf("%s: %.1f/%.1f/%.1f\n", k, v.Min(), v.Mean(), v.Max())
	}
}
