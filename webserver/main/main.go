package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RequestBody struct {
    Name string `json:"name"`
}

func SayHello(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var t RequestBody
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    var response =make(map[string]string)
    response["message"] = "Hello, " + t.Name
    json.NewEncoder(w).Encode(response)
}

func main() {

    http.HandleFunc("/hello/", SayHello)
    log.Fatal(http.ListenAndServe(":8080", nil))
}













