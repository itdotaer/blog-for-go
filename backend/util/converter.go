package util

import (
	"blog-for-go/datamodels"
	"blog-for-go/web/models"
)

func ConvertToVM(post datamodels.Post) models.Post {
	viewPost := models.Post{}

	viewPost.Id = post.Id
	viewPost.Title = post.Title
	viewPost.Description = post.Description
	viewPost.PostUser = post.PostUser
	viewPost.Content = post.Content
	viewPost.CreateTime = post.CreateTime
	viewPost.UpdateTime = post.UpdateTime

	return viewPost
}

func ConvertToDM(post models.Post) datamodels.Post {
	dataPost := datamodels.Post{}

	dataPost.Id = post.Id
	dataPost.Title = post.Title
	dataPost.Description = post.Description
	dataPost.PostUser = post.PostUser
	dataPost.Content = post.Content

	return dataPost
}
