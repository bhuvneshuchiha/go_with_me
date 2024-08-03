package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/bhuvneshuchiha/todo_app/routes"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    r := mux.NewRouter()
    routes.RegisterTodoRoutes(r)
    http.Handle("/", r)
    //This port should be the port where you want to listen for the requests
    log.Fatal(http.ListenAndServe("localhost:9010", r))
}
