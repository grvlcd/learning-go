package main

import (
	"errors"
	"fmt"
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

func WriteToFile(content string, fileName string) {
	contentText := fmt.Sprint(content)
	os.WriteFile(fileName, []byte(contentText), 0644)
}
