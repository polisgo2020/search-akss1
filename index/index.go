package index

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"searchera/utils"
	"strings"
)

type Freq map[string]int
type ReverseIdx map[string][]Freq

func Make(dirPath string) (ReverseIdx, error) {
	revIdx := ReverseIdx{}

	re := regexp.MustCompile(`[\p{L}\d]+`) // find also unicode characters

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return revIdx, err
	}

	for _, f := range files {
		if f.IsDir() {
			log.Printf("File %s is a directory. Skipping...", err)
			continue
		}

		data, err := ioutil.ReadFile(path.Join(dirPath, f.Name()))
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
