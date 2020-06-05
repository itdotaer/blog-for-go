package controllers

import (
	"blog-for-go/cache"
	"blog-for-go/datamodels"
	"blog-for-go/services"
	"blog-for-go/util"
	"blog-for-go/web/models"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"log"
)

const (
	SIZE = 10
)

type PostController struct {
	Service services.PostService
}

func (c *PostController) Get(ctx iris.Context) models.Resp {
	index, _ := ctx.URLParamInt("index")
	size, _ := ctx.URLParamInt("size")

	if size == 0 {
		size = SIZE
	}

	key := fmt.Sprintf("posts_%d_%d", index, size)
	postStr := cache.Get(key)

	var posts []datamodels.Post
	json.Unmarshal([]byte(postStr), &posts)

	if len(posts) == 0 {
		posts = c.Service.Query(index, size)

		if len(posts) > 0 {
			postsBytes, _ := json.Marshal(posts)
			cache.Set(key, string(postsBytes), 100)
		}
	}
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
