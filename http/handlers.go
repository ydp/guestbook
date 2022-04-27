package http

import (
    "context"
    "fmt"
    "encoding/json"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "html/template"
    "log"
    "net/http"
    "time"
)

var Client *mongo.Client

func ViewHandler(w http.ResponseWriter, r *http.Request) {
    findOptions := options.Find()
    findOptions.SetSort(bson.D{{"createtime", -1}})
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

func PostHandler(w http.ResponseWriter, r *http.Request) {
    comment := Comment{Message: r.FormValue("message"), CreateTime: time.Now().Format("01-02-2006 15:04:05")}
    collection := Client.Database("guestbook").Collection("comments")
    insertResult, err := collection.InsertOne(context.TODO(), comment)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Inserted signature: ", insertResult.InsertedID)
    http.Redirect(w, r, "/guestbook", http.StatusFound)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
    var comment Comment
    err := json.NewDecoder(r.Body).Decode(&comment)
    comment.CreateTime = time.Now().Format("01-02-2006 15:04:05")
    check(err)
    collection := Client.Database("guestbook").Collection("comments")
    insertResult, err := collection.InsertOne(context.TODO(), comment)
    check(err)
    fmt.Println("Inserted signature: ", insertResult.InsertedID)
    fmt.Fprintf(w, "success")
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
    findOptions := options.Find()
    findOptions.SetSort(bson.D{{"createtime", -1}})
    collection := Client.Database("guestbook").Collection("comments")
    cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
    check(err)
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
    fmt.Println("List comments count:", len(comments))
    guestbook := Guestbook{Comments: comments, CommentCount: len(comments)}
    json, _ := json.Marshal(guestbook)
    fmt.Fprint(w, string(json))
}

