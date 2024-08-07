package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type Saver interface {
    Save() error
}

type Displayer interface {
    Display()
}

type Outputtable interface {
    Saver
    Displayer
}

func main() {

    title, content, err := getNoteData()
    todoText := getUserInput("Todo text: ")

    todo, err := todo.New(todoText)
    if err != nil{
        fmt.Print(err)
        return
    }

    todo.Display()

    err = todo.Save()
    if err != nil {
        fmt.Println("Saving the todo failed")
        return
    }
    fmt.Println("Todo was saved successfully")

    userNote, err := note.New(title, content)

    if err != nil {
        fmt.Print(err)
        return
    }
    userNote.Display()
    err = saveData(userNote)

    if err != nil {
        return
    }
}

func outputData(data Outputtable)error {
    data.Display()
    return saveData(data)
}

func saveData(data Saver) error{

    err := data.Save()

    if err != nil {
        fmt.Println("Error encountered")
        return err
    }
    fmt.Println("\nSaving the note successfully!!")
    return nil
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

