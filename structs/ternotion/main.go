package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/ternotion/note"
	"example.com/ternotion/todo"
)

func main() {
	title, content := getUserData()

	todoContent := getTodoData()

	newTodo, err := todo.New(todoContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = newTodo.Save()

	if err != nil {
		fmt.Println("Saving the todo failed!")
		return
	}

	newTodo.Display()

	fmt.Println("Saving the todo success!")

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
