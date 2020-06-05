package datamodels

import "time"

type User struct {
	Id          int64
	Name        string
	Password    string
	Email       string
	Phone       string
	Description string
	CreateUser  string
	UpdateUser  string
	CreateTime  time.Time
	UpdateTime  time.Time
}
