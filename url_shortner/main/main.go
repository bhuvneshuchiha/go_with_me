package main

import (
	"log"
	"net/http"

	"github.com/bhuvneshuchiha/url_shortner/controllers"
)

func main() {
	http.HandleFunc("/", controllers.ReceiveUrl)
	http.HandleFunc("/redirect", controllers.RedirectToMain)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal("Server went down")
	}
}
