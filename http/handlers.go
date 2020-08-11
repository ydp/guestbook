package http

import (
	"fmt"
	"html/template"
	"net/http"
)

func ViewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := guestbook{
		signatureCount: len(signatures),
		signatures: signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

