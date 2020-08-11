package http

import (
	"html/template"
	"net/http"
)

func ViewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

