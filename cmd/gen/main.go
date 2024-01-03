package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/clfs/obrc"
)

func main() {
	var (
		nFlag          = flag.Int("n", 0, "number of measurements to generate")
		cpuprofileFlag = flag.String("cpuprofile", "", "write cpu profile to file")
	)
	flag.Parse()

	if *nFlag == 0 {
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

	if err := obrc.GenerateCSV(os.Stdout, *nFlag); err != nil {
		log.Fatal(err)
	}
}
