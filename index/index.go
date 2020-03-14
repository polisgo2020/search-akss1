package index

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"searchera/utils"
	"strings"
)

type Freq map[string]int
type ReverseIdx map[string][]Freq

func Read(idxFile string) (ReverseIdx, error) {
	idx := ReverseIdx{}

	b, err := ioutil.ReadFile(idxFile)
	if err != nil {
		return idx, err
	}

	err = json.Unmarshal(b, &idx)
	return idx, err
}

func SearchIn(idx ReverseIdx, substr string) Freq {
	found := Freq{}
	tkns := GetTokensFromStr(substr)

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

func ReadAndSearch(idxPath string, substr string) error {
	idx, err := Read(idxPath)
	if err != nil {
		return err
	}

	found := SearchIn(idx, substr)
	if len(found) == 0 {
		log.Println("Input string's tokens not found in reverse index")
		return nil
	}

	for k, v := range found {
		log.Printf("%s; %d matches", k, v)
	}
	return nil
}

func Make(dirPath string) (ReverseIdx, error) {
	revIdx := ReverseIdx{}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return revIdx, err
	}

	for _, f := range files {
		if f.IsDir() {
			log.Printf("File %s is a directory. Skipping...", err)
			continue
		}

		data, err := ioutil.ReadFile(filepath.Join(dirPath, f.Name()))
		if err != nil {
			log.Println("File reading error", err)
			continue
		}

		freq := Freq{} // tokens frequency in current file

		tokens := GetTokensFromStr(string(data))
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

func MakeAndWrite(dirPath, outPath string) error {
	revIdx, err := Make(dirPath)
	if err != nil {
		return err
	}

	b, err := json.Marshal(revIdx)
	if err != nil {
		return err
	}

	err = utils.WriteToFile(b, outPath)
	return err
}

func GetTokensFromStr(str string) []string {
	re := regexp.MustCompile(`[\p{L}\d]+`) // find also unicode characters
	tokens := re.FindAllString(str, -1)
	return tokens
}
