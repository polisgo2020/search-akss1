package index

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"searchera/utils"
	"strings"
)

type Freq map[string]int
type ReverseIdx map[string][]Freq

func (idx ReverseIdx) Search(substr string) Freq {
	found := Freq{}
	tkns := utils.GetTokensFromStr(substr)
	for _, t := range tkns {
		t := strings.ToLower(t)

		freq, ok := idx[t]
		if !ok {
			log.Printf("Token '%s' not found in reverse index", t)
			continue
		}

		for _, f := range freq {
			for k, v := range f {
				found[k] += v
			}
		}
	}

	return found
}

func MakeIndex(dirPath string) (ReverseIdx, error) {
	revIdx := ReverseIdx{}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return revIdx, err
	}

	for _, f := range files {
		if f.IsDir() {
			log.Printf("File '%s' is a directory. Skipping...", f.Name())
			continue
		}

		data, err := ioutil.ReadFile(filepath.Join(dirPath, f.Name()))
		if err != nil {
			log.Println("File reading error", err)
			continue
		}

		freq := Freq{} // tokens frequency in current file

		tokens := utils.GetTokensFromStr(string(data))
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

	log.Println("Reverse index successfully make")

	return revIdx, nil
}