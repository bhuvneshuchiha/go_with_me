package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Note struct {
    Content string `json:"content"`
}

func (note Note) Display() {
    fmt.Printf("Content : \n%v", note.Content)
}

func (note Note) Save() error{
    fileName := strings.ReplaceAll(note.Content, " ", "_")
    fileName = strings.ToLower(fileName) + ".json"

    json, err := json.Marshal(note)

    if err != nil {
        return err
    }
    return os.WriteFile(fileName, json, 0644)
}

func New(content string) (*Note, error) {

    if content == "" {
        return &Note{}, errors.New("Invalid input")
    }

    return &Note{
        Content: content,
    }, nil

}
