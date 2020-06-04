package datamodels

import "time"

type Post struct {
	Id          int64
	Title       string
	PostUser    string
	Description string
	Content     string
	CreateUser  string
	UpdateUser  string
	CreateTime  time.Time
	UpdateTime  time.Time
}
