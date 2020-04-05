package index

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"searchera/utils"
	"strings"
	"sync"

	"github.com/zoomio/stopwords"
)

type Freq map[string]int

type ReverseIdx struct {
	mx sync.Mutex
	M  map[string][]Freq
}

func (idx *ReverseIdx) Init() {
	idx.M = make(map[string][]Freq)
}

func (idx ReverseIdx) Search(substr string) Freq {
	found := Freq{}
	tokens := utils.GetWordsFromStr(substr)
	for _, t := range tokens {
		t := strings.ToLower(t)

		freq, ok := idx.M[t]
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

func GetTokensFromText(data string) Freq {
	freq := Freq{} // tokens frequency in current file

	words := utils.GetWordsFromStr(data)

	// calculate tokens freq for file
	for _, w := range words {
		token := strings.ToLower(w)

		if stopwords.IsStopWord(token) {
			continue
		}

		freq[token]++
	}

	return freq
}

func MakeIndex(dirPath string) (ReverseIdx, error) {
	revIdx := ReverseIdx{}
	revIdx.Init()

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return revIdx, err
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(files))
	for _, f := range files {
		go func(file os.FileInfo) {
			defer wg.Done()

			if file.IsDir() {
				log.Printf("File '%s' is a directory. Skipping...", file.Name())
				return
			}

			data, err := ioutil.ReadFile(filepath.Join(dirPath, file.Name()))
			if err != nil {
				log.Println("File reading error", err)
				return
			}

			freq := GetTokensFromText(string(data)) // tokens frequency in current file

			for tok, num := range freq {
				revIdx.mx.Lock()
				revIdx.M[tok] = append(revIdx.M[tok], Freq{file.Name(): num})
				revIdx.mx.Unlock()
			}
		}(f)
	}

	wg.Wait()
	log.Println("Reverse index successfully make")

	return revIdx, nil
}
