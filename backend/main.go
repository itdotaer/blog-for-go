package main

import (
	"blog-for-go/datasource"
	"blog-for-go/repositories"
	"blog-for-go/services"
	"blog-for-go/web/controllers"
	"blog-for-go/web/middlewares"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})

	mvc.Configure(app.Party("/posts", crs), posts)
	mvc.Configure(app.Party("/users", crs), users)

	app.Run(
		//开启web服务
		iris.Addr(":8080"),
		// 按下CTRL / CMD + C时跳过错误的服务器：
		iris.WithoutServerError(iris.ErrServerClosed),
		//实现更快的json序列化和更多优化：
		iris.WithOptimizations,
	)
}

//注意mvc.Application，它不是iris.Application。
func posts(app *mvc.Application) {
	app.Router.Use(middlewares.Auth)

	repo := repositories.NewPostRepo(datasource.MysqlDb)
	postService := services.NewPostService(repo)

	app.Register(postService)
	app.Handle(new(controllers.PostController))
}

func users(app *mvc.Application) {
	repo := repositories.NewUserRepo(datasource.MysqlDb)
	userService := services.NewUserService(repo)

	app.Register(userService)
	app.Handle(new(controllers.UserController))
}
