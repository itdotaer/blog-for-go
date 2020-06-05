package middlewares

import (
	"blog-for-go/cache"
	"blog-for-go/web/models"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
)

const (
	AUTH_KEY = "auth"
	EXP      = 60 * 60 * 24
)

func Auth(ctx iris.Context) {
	auth := ctx.GetCookie(AUTH_KEY)
	// 看看redis里有没有
	name := cache.Get(auth)

	if name == "" {
		// 没有权限
		ctx.StatusCode(401)
		jsonBytes, _ := json.Marshal(models.Resp{Success: false, Code: "not_auth", Msg: "not login", Data: nil})
		ctx.WriteString(string(jsonBytes))
		ctx.StopExecution()
	}

	ctx.Next()
}

func Login(ctx iris.Context, name string) {
	data := []byte(name)
	hash := md5.Sum(data)

	hashStr := fmt.Sprintf("%x", hash)
	ctx.SetCookieKV(AUTH_KEY, hashStr, iris.CookieExpires(time.Hour))
	cache.Set(hashStr, name, EXP)
}
