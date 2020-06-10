package services

import (
	"blog-for-go/datamodels"
	"blog-for-go/repositories"
)

type PostService interface {
	Query(index int, size int) []datamodels.Post
	QueryById(id int64) datamodels.Post
	Update(post datamodels.Post, name string) bool
}

func NewPostService(repo repositories.PostRepo) PostService {
	return &postService{repo: repo}
}

type postService struct {
	repo repositories.PostRepo
}

func (s *postService) Query(index int, size int) []datamodels.Post {
	return s.repo.Query(index, size)
}

func (s *postService) QueryById(id int64) datamodels.Post {
	return s.repo.QueryById(id)
}

func (s *postService) Update(post datamodels.Post, name string) bool {
	if name == "" {
		return false
	}

	if post.Id <= 0 {
		// insert
		post.CreateUser = name
		post.UpdateUser = name
		return s.repo.Insert(post)
	}

	post.UpdateUser = name
	return s.repo.Update(post)
}
