package services

import (
	"blog-for-go/datamodels"
	"blog-for-go/repositories"
)

type UserService interface {
	Query(index int, size int) []datamodels.User
	QueryBy(name string, password string) []datamodels.User
	Update(post datamodels.User) bool
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userService{repo: repo}
}

type userService struct {
	repo repositories.UserRepo
}

func (s *userService) Query(index int, size int) []datamodels.User {
	return s.repo.Query(index, size)
}

func (s *userService) QueryBy(name string, password string) []datamodels.User {
	return s.repo.QueryBy(name, password)
}

func (s *userService) Update(post datamodels.User) bool {
	if post.Id <= 0 {
		// insert
		return s.repo.Insert(post)
	}

	return s.repo.Update(post)
}
