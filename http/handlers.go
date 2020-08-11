package http

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"log"
	"net/http"
)

var Client *mongo.Client

type Signature struct {
	Signature string
}

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
	signature := Signature{ Signature: r.FormValue("signature") }
	collection := Client.Database("guestbook").Collection("signatures")
	insertResult, err := collection.InsertOne(context.TODO(), signature)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted signature: ", insertResult.InsertedID)
	http.Redirect(w, r, "/guestbook", http.StatusFound)
}