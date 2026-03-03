package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/ternotion/note"
	"example.com/ternotion/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getUserData()

	todoContent := getTodoData()

	newTodo, err := todo.New(todoContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = saveData(newTodo)

	if err != nil {
		return
	}

	newTodo.Display()

	err = saveData(newNote)

	if err != nil {
		return
	}

	newNote.Display()
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed!")
		return err
	}

	fmt.Println("Saving the note success!")
	return nil
}

func getTodoData() string {
	content := getUserInput("Todo: ")
	return content
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
