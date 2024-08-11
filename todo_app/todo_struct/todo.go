package todo_struct

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	fileOps "github.com/bhuvneshuchiha/todo_app/file_ops"
	"github.com/bhuvneshuchiha/todo_app/models"
)

var i int = 1
var fileName = fileOps.GetFileName()
var todo_list []models.Todo

// Add todo
func CreateTodo(title string, con string) string {
	if title == "" || con == "" {
		log.Fatal("Params cannot be empty")
		return "Could not create todo"
	}

	var t models.Todo
    t.Id = GetId()
	t.Title = title
	t.Content = con
	t.Status = "false"

    data, err := os.ReadFile("out.txt")
    if err != nil {
        fmt.Println(err)
    }

    if len(data) > 0 {
        err = json.Unmarshal(data, &todo_list)
        if err != nil {
            log.Fatal(err)
        }
    }
    todo_list = append(todo_list, t)

    updatedData, err := json.MarshalIndent(todo_list, "", " ")
    if err != nil {
        fmt.Println(err)
    }

    err = os.WriteFile("out.txt", updatedData, 0644)
    if err != nil {
        fmt.Println(err)
    }
    SaveId(i)
	return "Successfully created Todo"
}

// Remove Todo
func RemoveTodo(id int) string {
    if id > len(todo_list){
        log.Fatal("Check ID")
        return "Couldnt remove"
    }
    todo_list = append(todo_list[:id - 1], todo_list[id:]...)
    res, err := fileOps.ReadFile()
    if err != nil {
        fmt.Printf("Error occured while some shit %v", err)
    }
    fmt.Println(res)

	return "Todo item deleted Successfully"
}

// @@Set the status
func UpdateStatus(id int, title string, content string) {
	// Read the file
    data, err := fileOps.ReadFile()
    if err != nil {
        fmt.Println(err)
        return
    }

    if id < 0 || id >= len(data){
        log.Fatal("Invalid id ", id)
        return
    }
    data[id].Title = title
    data[id].Content = content
    fmt.Println("Data updated successfully")
}


// Print Todos
func PrintTodos() {
    for id, todo := range todo_list {
        fmt.Printf(
            "ID: %d\nTitle: %s\nContent: %s\nStatus: %v\n\n",
            id + 1,
            todo.Title,
            todo.Content,
            todo.Status,
        )
    }
}

// @@ Save the id so that the next time the program runs the ID can be
    // continued
func SaveId(id int) {
    idData := strconv.Itoa(id)
    err := os.WriteFile("id.txt", []byte(idData), 0644)
    if err != nil {
        fmt.Print(err)
    }
    fmt.Println("ID saved successfully")
}

// @@ Function to get id which is saved in the file
func GetId() int {
    data, err := os.ReadFile("id.txt")
    if err != nil {
        fmt.Println("Error while trying to read the file")
    }
    id, err := strconv.Atoi(string(data))
    if err != nil {
        fmt.Println("Caught an error while converting the id from str to int")
    }
    return id + 1
}





















