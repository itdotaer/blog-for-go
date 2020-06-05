package models

type Resp struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
