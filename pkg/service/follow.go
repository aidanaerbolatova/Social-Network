package service

import (
	"Forum/models"
	"Forum/pkg/repository"
)

type FollowService struct {
	repo repository.Follow
}

func NewFollowService(repo repository.Follow) *FollowService {
	return &FollowService{repo: repo}
}

func (s *FollowService) CreateFollow(follow models.Follow) error {
	return s.repo.CreateFollow(follow)
}

func (s *FollowService) DeleteFollow(follow models.Follow) error {
	return s.repo.DeleteFollow(follow)
}

func (s *FollowService) CheckFollow(follow models.Follow) error {
	return s.repo.CheckFollow(follow)
}

func (s *FollowService) MyFollowers(userId int) ([]string, error) {
	return s.repo.MyFollowers(userId)
}
func (s *FollowService) Following(userId int) ([]string, error) {
	return s.repo.Following(userId)
}
