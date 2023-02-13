package service

import (
	"Forum/models"
	"Forum/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePosts(post models.Post) error {
	return s.repo.CreatePosts(post)
}

func (s *PostService) GetPostByUserID(id int) ([]models.Post, error) {
	return s.repo.GetPostByUserID(id)
}

func (s *PostService) GetPost() (**[]models.Post, error) {
	result, _ := s.repo.GetPost()
	return &result, nil
}

func (s *PostService) GetPostByTag(tags string) (**[]models.Post, error) {
	result, _ := s.repo.GetPostByTag(tags)
	return &result, nil
}

func (s *PostService) GetPostByPostID(postId int) (models.Post, error) {
	return s.repo.GetPostByPostID(postId)
}

func (s *PostService) UpdatePost(like string, dislike string, postId int) error {
	return s.repo.UpdatePost(like, dislike, postId)
}

func (s *PostService) LikedPosts(userId int) ([]models.Post, error) {
	return s.repo.LikedPosts(userId)
}

func (s *PostService) DislikedPosts(userId int) ([]models.Post, error) {
	return s.repo.DislikedPosts(userId)
}

func (s *PostService) CommentedPosts(userId int) ([]models.Post, error) {
	return s.repo.CommentedPosts(userId)
}

func (s *PostService) DeletePost(postId, userId int) error {
	return s.repo.DeletePost(postId, userId)
}

func (s *PostService) EditPost(postId int, title, text string) error {
	return s.repo.EditPost(postId, title, text)
}

func (s *PostService) CreateNotification(author, username, action string) error {
	return s.repo.CreateNotification(author, username, action)
}

func (s *PostService) GetNotification(username string) ([]models.Notifications, error) {
	return s.repo.GetNotification(username)
}
