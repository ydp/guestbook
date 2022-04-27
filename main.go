package main

import (
    "log"
    "net/http"
    h "github.com/ydp/guestbook/http"
    db "github.com/ydp/guestbook/database"
)

func main() {
    h.Client = db.Connect()
    http.HandleFunc("/guestbook", h.ViewHandler)
    http.HandleFunc("/guestbook/view/new", h.NewHandler)
    http.HandleFunc("/guestbook/post", h.PostHandler)
    http.HandleFunc("/guestbook/create", h.CreateHandler)
    http.HandleFunc("/guestbook/list", h.ListHandler)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal(err)
    }
}
