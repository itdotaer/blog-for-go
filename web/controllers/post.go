package controllers

import (
	"blog-for-go/services"
	"blog-for-go/util"
	"blog-for-go/web/models"
	"github.com/kataras/iris/v12"
	"log"
)

type PostController struct {
	Service services.PostService
}

func (c *PostController) Get() models.Resp {
	posts := c.Service.Query(0, 100)
	var data []models.Post
	for _, post := range posts {
		data = append(data, util.ConvertToVM(post))
	}

	return models.Resp{Success: true, Code: "true", Data: data}
}

func (c *PostController) Post(ctx iris.Context) models.Resp {
	viewPost := &models.Post{}

	err := ctx.ReadJSON(viewPost)
	if err != nil {
		log.Println("post data wrong")
		return models.Resp{Success: false, Code: "false", Msg: "post data wrong", Data: nil}
	}

	// insert or update
	rs := c.Service.Update(util.ConvertToDM(*viewPost))

	return models.Resp{Success: rs, Code: "true", Msg: "ok", Data: nil}
}
