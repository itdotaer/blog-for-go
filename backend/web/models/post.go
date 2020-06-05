package models

type Post struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	PostUser    string `json:"postUser"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
