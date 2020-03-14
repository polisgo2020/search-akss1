package main

import (
	"flag"
	"log"
	"searchera/index"
)

func main() {
	methodPtr := flag.String("m", "", "Method: search or make")
	subStrPtr := flag.String("str", "", "Str for search")
	inPathPtr := flag.String("in", "", "Input path")
	outPathPtr := flag.String("out", "", "Output path")
	flag.Parse()

	switch *methodPtr {
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
