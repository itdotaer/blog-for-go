package services

import (
	"blog-for-go/datamodels"
	"blog-for-go/repositories"
)

type PostService interface {
	Query(index int, size int) []datamodels.Post
	Update(post datamodels.Post) bool
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

func (s *postService) Update(post datamodels.Post) bool {
	if post.Id <= 0 {
		// insert
		return s.repo.Insert(post)
	}

	return s.repo.Update(post)
}
