package main

import (
	"log"
	"net/http"
	h "guestbook/http"
)

func main() {
	http.HandleFunc("/guestbook", h.ViewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
