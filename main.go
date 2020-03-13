package main

import (
	"flag"
	"log"
	"searchera/index"
)

func main() {
	actPtr := flag.String("act", "search", "Action: search or make")
	inPathPtr := flag.String("in", "", "Input path")
	outPathPtr := flag.String("out", "", "Output path")
	flag.Parse()

	switch *actPtr {
	case "search":
		// TODO: Add search
	case "make":
		err := index.MakeAndWrite(*inPathPtr, *outPathPtr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
