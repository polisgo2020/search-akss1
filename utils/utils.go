package utils

import (
	"os"
	"regexp"
)

func WriteToFile(b []byte, out string) error {
	file, err := os.Create(out)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(b)
	return err
}

func GetTokensFromStr(str string) []string {
	re := regexp.MustCompile(`[\p{L}\d]+`) // find also unicode characters
	tokens := re.FindAllString(str, -1)
	return tokens
}