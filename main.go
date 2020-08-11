package main

import (
	"log"
	"net/http"
	h "guestbook/http"
)

func main() {
	http.HandleFunc("/guestbook", h.ViewHandler)
	http.HandleFunc("/guestbook/new", h.NewHandler)
	http.HandleFunc("/guestbook/create", h.CreateHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
