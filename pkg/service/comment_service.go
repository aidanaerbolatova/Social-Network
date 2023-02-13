package service

import (
	"Forum/models"
	"Forum/pkg/repository"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) AddComment(comment models.Comments) error {
	return s.repo.AddComments(comment)
}

func (s *CommentService) GetCommentByPost(postId int) ([]models.Comments, error) {
	return s.repo.GetCommentByPost(postId)
}

func (s *CommentService) UpdateComment(like string, dislike string, commentId int) error {
	return s.repo.UpdateComment(like, dislike, commentId)
}
func (s *CommentService) DeleteComment(commentId, userId int) error {
	return s.repo.DeleteComment(commentId, userId)
}

func (s *CommentService) GetCommentById(commentID int) (models.Comments, error) {
	return s.repo.GetCommentById(commentID)
}
func (s *CommentService) EditComment(commentID int, comment string) error {
	return s.repo.EditComment(commentID, comment)
}
