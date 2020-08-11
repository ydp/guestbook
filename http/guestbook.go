package http

type Comment struct {
	Name string
	Comment string
}

type Guestbook struct {
	CommentCount int
	Comments []Comment
}