package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/ternotion/note"
)

func main() {
	title, content := getUserData()

	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.Display()

	err = newNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed!")
		return
	}

	fmt.Println("Saving the note success!")
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
