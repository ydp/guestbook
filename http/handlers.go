package http

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"log"
	"net/http"
)

var Client *mongo.Client

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	findOptions := options.Find()
	collection := Client.Database("guestbook").Collection("comments")
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	var comments []Comment
	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem Comment
		err := cur.Decode(&elem)
		check(err)

		comments = append(comments, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{Comments: comments, CommentCount: len(comments)}
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
	comment := Comment{Name: r.FormValue("name"), Message: r.FormValue("message")}
	collection := Client.Database("guestbook").Collection("comments")
	insertResult, err := collection.InsertOne(context.TODO(), comment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted signature: ", insertResult.InsertedID)
	http.Redirect(w, r, "/guestbook", http.StatusFound)
}