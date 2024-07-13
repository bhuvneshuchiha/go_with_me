package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
)

func main() {

    title, content, err := getNoteData()
    userNote, err := note.New(title, content)

    if err != nil {
        fmt.Print(err)
        return
    }
    userNote.Display()
}

func getNoteData() (string, string, error){

    title := getUserInput("Note title : ")
    content:= getUserInput("Note content : ")

    return title, content, nil
}

func getUserInput(prompt string) (string){

    fmt.Print(prompt)
    //Important way to read input of more than a word
    reader := bufio.NewReader(os.Stdin)
    text, err := reader.ReadString('\n')

    if err != nil {
        return ""
    }

    text = strings.TrimSuffix(text, "\n")
    text = strings.TrimSuffix(text, "\r")

   return text
}

