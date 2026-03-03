package todo

import (
	"encoding/json"
	"errors"
	"fmt"

	"example.com/ternotion/fileops"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	filePath := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	fileops.WriteToFile(json, filePath)

	return err
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid Input.")
	}

	return Todo{
		Text: content,
	}, nil
}
