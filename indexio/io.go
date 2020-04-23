/*
Package indexio contains functions for index input and output.
*/
package indexio

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
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
		log.Info().Msgf(`String "%s" tokens not found in reverse index`, substr)
		return nil
	}

	for k, v := range found {
		log.Info().Int(k, v)
	}
	return nil
}

func MakeAndWriteIndex(dirPath string, outPath string) error {
	revIdx, err := index.MakeIndex(dirPath)
	if err != nil {
		return err
	}

	b, err := json.Marshal(revIdx.M)
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

	err = json.Unmarshal(b, &idx.M)
	return idx, err
}
