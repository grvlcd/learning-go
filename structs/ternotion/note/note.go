package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"example.com/ternotion/fileops"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	filePath := strings.ReplaceAll(note.Title, " ", "_")
	filePath = strings.ToLower(filePath)
	filePath = filePath + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	fileops.WriteToFile(json, filePath)

	return err
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input.")
	}

	return Note{
		title,
		content,
		time.Now(),
	}, nil
}
