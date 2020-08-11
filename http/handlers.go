package http

import (
	"fmt"
	"html/template"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	signatures := getStrings("signatures.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(w, guestbook)
	check(err)
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	signature := r.FormValue("signature")
	fmt.Println(signature)
	http.Redirect(w, r, "/guestbook", http.StatusFound)
}