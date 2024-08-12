package controllers

import (
	"fmt"
	"net/http"
)

func ReceiveUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Request received bitches")
}

func RedirectToMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Send the request to the long url")
}
