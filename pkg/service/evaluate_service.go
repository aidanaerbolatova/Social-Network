package service

import (
	"Forum/models"
	"Forum/pkg/repository"
)

type EvaluateService struct {
	repo repository.Evaluate
}

func NewEvaluateService(repo repository.Evaluate) *EvaluateService {
	return &EvaluateService{repo: repo}
}

func (s *EvaluateService) CreateEvaluates(postLike models.Evaluate) error {
	return s.repo.CreateEvaluates(postLike)
}

func (s *EvaluateService) CheckUserPost(userId, postId int) (models.Evaluate, error) {
	return s.repo.CheckUserPost(userId, postId)
}

func (s *EvaluateService) UpdateVote(userId, postId, newVote int) error {
	return s.repo.UpdateVote(userId, postId, newVote)
}

func (s *EvaluateService) EvaluateCount(postId int) (models.Vote, error) {
	return s.repo.EvaluateCount(postId)
}

func (s *EvaluateService) CheckUserComment(userId, commentId int) (models.EvaluateComment, error) {
	return s.repo.CheckUserComment(userId, commentId)
}

func (s *EvaluateService) CreateEvaluateComment(commentVote models.EvaluateComment) error {
	return s.repo.CreateEvaluateComment(commentVote)
}

func (s *EvaluateService) UpdateCommentVote(userId, commentId, newVote int) error {
	return s.repo.UpdateCommentVote(userId, commentId, newVote)
}

func (s *EvaluateService) EvaluateCommentCount(commentId int) (models.Vote, error) {
	return s.repo.EvaluateCommentCount(commentId)
}

func (s *EvaluateService) CheckVote(userId, postId, vote int) error {
	return s.repo.CheckVote(userId, postId, vote)
}

func (s *EvaluateService) CheckCommentVote(userId, postId, vote int) error {
	return s.repo.CheckCommentVote(userId, postId, vote)
}
