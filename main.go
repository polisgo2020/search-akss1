package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Need path with files")
	}

	path := os.Args[1]

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	content := map[string][]byte{} // for the future
	revIdx := map[string][]string{}
	for _, f := range files {
		data, err := ioutil.ReadFile(path + "\\" + f.Name())
		if err != nil {
			log.Println("File reading error", err)
			continue
		}

		content[f.Name()] = data

		words := strings.Fields(string(data))
		for _, w := range words {
			revIdx[w] = append(revIdx[w], f.Name())
		}
	}

	b, err := json.Marshal(revIdx)
	if err != nil {
		log.Fatal(err)
	}

	err = WriteFile(b, "revIdx.json")
	if err != nil {
		log.Fatal(err)
	}
}
