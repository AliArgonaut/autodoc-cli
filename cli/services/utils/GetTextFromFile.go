package utils

import (
	"errors"
	"os"
)

func GetTextFromFile(path string) (string, error) {

	contents, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("couldnt read file")
	}

	return string(contents), nil
}
