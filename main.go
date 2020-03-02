package main

import (
	"io/ioutil"
	"log"
	"os"
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

	content := map[string][]byte{}
	for _, f := range files {
		data, err := ioutil.ReadFile(path + "\\" + f.Name())
		if err != nil {
			log.Fatal("File reading error", err)
		}
		content[f.Name()] = data
	}

}
