package indexio

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"searchera/index"
	"searchera/utils"
)

func ReadAndSearchIndex(idxPath string, substr string) error {
	idx, err := ReadIndex(idxPath)
	if err != nil {
		return err
	}

	found := idx.Search(substr)
	if len(found) == 0 {
		log.Println("Input string's tokens not found in reverse index")
		return nil
	}

	for k, v := range found {
		log.Printf("%s; %d matches", k, v)
	}
	return nil
}

func MakeAndWriteIndex(dirPath string, outPath string) error {
	revIdx, err := index.MakeIndex(dirPath)
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

func ReadIndex(idxFile string) (index.ReverseIdx, error) {
	idx := index.ReverseIdx{}

	b, err := ioutil.ReadFile(idxFile)
	if err != nil {
		return idx, err
	}

	err = json.Unmarshal(b, &idx)
	return idx, err
}