package http

type Comment struct {
    Message string `json:"message"`
    CreateTime string `json:"createTime"`
}

type Guestbook struct {
    CommentCount int `json:"count"`
    Comments []Comment `json:"comments"`
}
