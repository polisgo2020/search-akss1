package main

import (
	"flag"
	"log"
	"searchera/indexio"
)

func main() {
	methodPtr := flag.String("m", "", "Method: make or search")
	subStrPtr := flag.String("str", "", "Str for search")
	inPathPtr := flag.String("in", "", "Input path")
	outPathPtr := flag.String("out", "", "Output path")
	flag.Parse()

	switch *methodPtr {
	case "search":
		err := indexio.ReadAndSearchIndex(*inPathPtr, *subStrPtr)
		if err != nil {
			log.Fatal(err)
		}
	case "make":
		err := indexio.MakeAndWriteIndex(*inPathPtr, *outPathPtr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
