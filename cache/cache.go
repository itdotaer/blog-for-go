package cache

import (
	"blog-for-go/datasource"
	"time"
)

func Get(key string) (value string) {
	cmd := datasource.Redis.Get(datasource.Ctx, key)
	value = cmd.Val()

	return
}

func Set(key string, value string, exp int) (rs bool) {
	duration := time.Second * time.Duration(exp)

	cmd := datasource.Redis.Set(datasource.Ctx, key, value, duration)
	rs = cmd.Err() == nil

	return
}
