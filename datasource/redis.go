package datasource

import (
	"blog-for-go/util"
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var Redis *redis.Client
var Err error
var Ctx context.Context

const (
	REDIS_HOST        = "172.29.145.2"
	REDIS_PORT        = "10007"
	REDIS_REMOTE_HOST = "172.17.0.2"
	REDIS_REMOTE_PORT = "6379"
	REDIS_PASS_WORD   = ""
	REDIS_DATABASE    = 0
)

func init() {
	host := REDIS_HOST
	port := REDIS_PORT
	if util.Mode == "remote" {
		host = REDIS_REMOTE_HOST
		port = REDIS_REMOTE_PORT
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: REDIS_PASS_WORD,
		DB:       REDIS_DATABASE,
	})

	Ctx = context.Background()
	pong, err := Redis.Ping(Ctx).Result()
	log.Println(pong, err)
}
