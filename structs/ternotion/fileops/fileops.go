package fileops

import (
	"errors"
	"os"
)

func ReadToFile(path string) (string, error) {
	content, err := os.ReadFile(path)

	contentValue := string(content)

	if err != nil {
		return "", errors.New("Failed to parse data")
	}

	return contentValue, nil
}

func WriteToFile(content []byte, fileName string) {
	os.WriteFile(fileName, content, 0644)
}
