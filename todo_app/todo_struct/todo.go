package todo_struct

import (
	"fmt"
	"log"

	fileOps "github.com/bhuvneshuchiha/todo_app/file_ops"
	"github.com/bhuvneshuchiha/todo_app/models"
)

var i uint8 = 1
var todo_list []models.Todo

// Add todo
func CreateTodo(title string, con string) string {
	if title == "" || con == "" {
		log.Fatal("Params cannot be empty")
		return "Could not create todo"
	}
	var t models.Todo
    t.Id = i
	t.Title = title
	t.Content = con
	t.Status = "false"

    todo_list = append(todo_list, t)
	i++

    fileOps.WriteData(t)
	return "Successfully created Todo"
}

// Remove Todo
func RemoveTodo(id uint8) string {
    if id > uint8(len(todo_list)){
        log.Fatal("Check ID")
        return "Couldnt remove"
    }
    todo_list = append(todo_list[:id], todo_list[id + 1:]...)

	return "Todo item deleted Successfully"
}

// Set the status
func UpdateStatus(id uint8) string {
	if id > uint8(len(todo_list)) {
		log.Fatal("Index is out of bounds")
		return "Wrong index"
	}
	fmt.Print(todo_list)
	return "Successfully has been marked complete"
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





















