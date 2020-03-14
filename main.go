package main

import (
	"flag"
	"log"
	"searchera/index"
)

func main() {
	actPtr := flag.String("act", "search", "Action: search or make")
	subStrPtr := flag.String("str", "", "Str for search")
	inPathPtr := flag.String("in", "", "Input path")
	outPathPtr := flag.String("out", "", "Output path")
	flag.Parse()

	switch *actPtr {
	case "search":
		err := index.ReadAndSearch(*inPathPtr, *subStrPtr)
		if err != nil {
			log.Fatal(err)
		}
	case "make":
		err := index.MakeAndWrite(*inPathPtr, *outPathPtr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
