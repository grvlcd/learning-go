package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/ternotion/note"
)

const filePath = "learn_go.json"

func main() {
	title, content := getUserData()

	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.Display()
}

func getUserData() (string, string) {
	title := getUserInput("Title: ")

	content := getUserInput("Content: ")

	return title, content
}

func getUserInput(promt string) string {
	fmt.Print(promt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
