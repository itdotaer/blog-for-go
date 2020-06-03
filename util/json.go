package util

import (
	"encoding/json"
	"log"
)

func PrettyJSON(obj interface{}) string {

	str, err := json.Marshal(obj)
	if err != nil {
		log.Println("parse error")
		return ""
	}

	return string(str)

}
