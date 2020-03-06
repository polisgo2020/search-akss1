package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Freq map[string]int

func main() {
	revIdx := map[string][]Freq{}

	re := regexp.MustCompile(`[\p{L}\d]+`) // find also unicode characters

	if len(os.Args) < 3 {
		log.Fatal("Need two params: files_path output_file")
	}

	path := os.Args[1]
	out := os.Args[2]

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			log.Printf("File %s is directory. Skipping...", err)
			continue
		}

		data, err := ioutil.ReadFile(path + "/" + f.Name())
		if err != nil {
			log.Println("File reading error", err)
			continue
		}

		freq := Freq{} // tokens frequency in current file

		tokens := re.FindAllString(string(data), -1)

		// calculate tokens freq for file
		for _, t := range tokens {
			if len(t) == 1 {
				continue
			}

			w := strings.ToLower(t)
			freq[w]++
		}

		for w, num := range freq {
			revIdx[w] = append(revIdx[w], Freq{f.Name(): num})
		}
	}

	b, err := json.Marshal(revIdx)
	if err != nil {
		log.Fatal(err)
	}

	err = WriteToFile(b, out)
	if err != nil {
		log.Fatal(err)
	}
}
