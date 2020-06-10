package middlewares

import (
	"blog-for-go/cache"
	"blog-for-go/web/models"
	"crypto/md5"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
)

const (
	AUTH_KEY = "auth"
	EXP      = 60 * 60 * 24
)

func Auth(ctx iris.Context) {
	auth := ctx.GetCookie(AUTH_KEY)

	name := ""
	if auth != "" {
		// 看看redis里有没有
		name = cache.Get(auth)
	}

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
	b := make([]byte, 10)
	rand.Read(b)

	hash := md5.Sum([]byte(string(data) + string(b)))

	hashStr := fmt.Sprintf("%x", hash)
	ctx.SetCookieKV(AUTH_KEY, hashStr, iris.CookieExpires(time.Hour))
	cache.Set(hashStr, name, EXP)
}

func GetUserName(ctx iris.Context) (name string, err error) {
	auth := ctx.GetCookie(AUTH_KEY)
	name = ""
	if auth != "" {
		// 看看redis里有没有
		name = cache.Get(auth)
	}

	if name == "" {
		err = errors.New("user not login")
	}

	return
}
