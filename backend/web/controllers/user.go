package controllers

import (
	"blog-for-go/services"
	"blog-for-go/web/middlewares"
	"blog-for-go/web/models"
	"github.com/kataras/iris/v12"
	"log"
)

type UserController struct {
	Service services.UserService
}

func (c *UserController) PostLogin(ctx iris.Context) models.Resp {
	viewUser := &models.User{}

	err := ctx.ReadJSON(viewUser)
	if err != nil {
		log.Println("post data wrong")
		return models.Resp{Success: false, Code: "false", Msg: "user data wrong", Data: nil}
	}

	// validate name and password
	if viewUser.Name == "" || viewUser.Password == "" {
		return models.Resp{Success: false, Code: "login_failed", Msg: "user data wrong", Data: nil}
	}

	// query user by name and password
	users := c.Service.QueryBy(viewUser.Name, viewUser.Password)
	if len(users) == 0 || len(users) > 1 {
		// err login
		return models.Resp{Success: false, Code: "login_failed", Msg: "login failed"}
	}

	// set login cookie
	middlewares.Login(ctx, viewUser.Name)

	return models.Resp{Success: true, Code: "true", Msg: "ok", Data: nil}
}
