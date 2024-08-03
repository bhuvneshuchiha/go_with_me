package routes

import (
    "github.com/gorilla/mux"
    "github.com/bhuvneshuchiha/todo_app/views"
)

var RegisterTodoRoutes = func(router *mux.Router){
    router.HandleFunc("/todo", views.GetTodo).Methods("GET")
    router.HandleFunc("/todo/", views.CreateTodo).Methods("POST")
    router.HandleFunc("/todo/{todo_id}", views.UpdateTodo).Methods("PUT")
    router.HandleFunc("/todo/{todo_id}", views.DeleteTodo).Methods("DELETE")

}

