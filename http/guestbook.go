package http

type Comment struct {
	Name string
	Message string
}

type Guestbook struct {
	CommentCount int
	Comments []Comment
}