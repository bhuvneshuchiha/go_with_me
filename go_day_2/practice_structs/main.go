package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

func main() {

    title, content, err := getNoteData()
    todoText := getUserInput("Todo text: ")

    todo, err := todo.New(todoText)
    if err != nil{
        fmt.Print(err)
        return
    }

    todo.Display()


    userNote, err := note.New(title, content)

    if err != nil {
        fmt.Print(err)
        return
    }
    userNote.Display()
    err = userNote.Save()

    if err != nil {
        fmt.Println("Error encountered")
        return
    }
    fmt.Println("\nSaving the note successfully!!")
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

