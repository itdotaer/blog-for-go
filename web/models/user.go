package models

type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Email       string `json:"email`
	Phone       string `json:phone`
	Description string `json:description`
}
