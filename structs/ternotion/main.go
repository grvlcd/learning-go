package main

import (
	"errors"
	"fmt"
	"strings"
)

const filePath = "learn_go.json"

func main() {
	title, content, err := getUserData()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getUserData() (string, string, error) {
	title, err := getUserInput("Title: ")

	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Content: ")

	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(promt string) (string, error) {
	fmt.Print(promt)

	var value string
	fmt.Scan(&value)

	if value == "" {
		var builder strings.Builder
		builder.Write([]byte(promt))
		builder.Write([]byte(" "))
		builder.Write([]byte("is required"))
		result := builder.String()
		return "", errors.New(result)
	}

	return value, nil
}
