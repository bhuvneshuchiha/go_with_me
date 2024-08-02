package views

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/bhuvneshuchiha/todo_app/todo_struct"
    "github.com/bhuvneshuchiha/todo_app/utils"
    "github.com/gorilla/mux"
)

var NewTodo todo_struct.Todo

func GetTodo(w http.ResponseWriter, r *http.Request) {
    newTodos := todo_struct.GetAllTodos()
    res, _ := json.Marshal(newTodos)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}


func GetTodoById(w http.ResponseWriter, r * http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    ID, err := strconv.ParseInt(todoId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    todoDetails, _ := todo_struct.GetTodoById(ID)
    res,_ := json.Marshal(todoDetails)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}


func CreateTodo(w http.ResponseWriter, r *http.Request){
    CreateTodo := &todo_struct.Todo{}
    utils.ParseBody(r, CreateTodo)
    b := CreateTodo.CreateTodo()
    res, _ := json.Marshal(b)
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    ID, err := strconv.ParseInt(todoId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    todo := todo_struct.DeleteTodo(ID)
    res, _ := json.Marshal(todo)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}




















