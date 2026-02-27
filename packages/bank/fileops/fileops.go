package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadToFile(fileName string) (float64, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	valueText := string(content)
	data, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance.")
	}

	return data, nil
}

func WriteToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
