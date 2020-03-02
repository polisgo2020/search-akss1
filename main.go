package main

import (
	"fmt"
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

	content := map[string][]byte{}
	revIdx := map[string][]string{}
	for _, f := range files {
		data, err := ioutil.ReadFile(path + "\\" + f.Name())
		if err != nil {
			log.Fatal("File reading error", err)
		}

		content[f.Name()] = data

		words := strings.Split(string(data), " ")
		for _, w := range words {
			revIdx[w] = append(revIdx[w], f.Name())
		}
	}
	fmt.Println(revIdx)
}
