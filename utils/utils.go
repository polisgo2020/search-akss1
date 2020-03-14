package utils

import (
	"os"
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
