package util

import (
	"flag"
	"log"
)

var Mode string

func init() {
	// 解析启动模式
	mode := flag.String("mode", "local", "get service start mode(exp: local, remote")
	flag.Parse()
	log.Println("mode", *mode)

	Mode = *mode
}
