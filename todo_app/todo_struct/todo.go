package todo_struct

import (
    "github.com/jinzhu/gorm"
    "github.com/bhuvneshuchiha/todo_app/db"
)

var database *gorm.DB

type Todo struct {
    gorm.Model
    Title string `gorm:""json:"name"`
    Content string `json:"content"`
}

func init() {
    db.Connect()
    database := db.GetDB()
    database.AutoMigrate(&Todo{})
}

func (todo *Todo) CreateTodo() *Todo {
    database.NewRecord(todo)
    database.Create(&todo)
    return todo
}

func GetAllTodos() []Todo {
    var Todo []Todo
    database.Find(&Todo)
    return Todo
}

func GetTodoById(ID int64) (*Todo, *gorm.DB) {
    var getTodo Todo
    db := database.Where("ID = ?", ID).Find(&getTodo)
    return &getTodo, db
}

func DeleteTodo(ID int64) Todo {
    var todo Todo
    database.Where("ID=?", ID).Delete(todo)
    return todo
}






















