package service

import (
	"Forum/models"
	"Forum/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) error
	GenerateToken(user models.User, oauth bool) (models.Token, error)
	GetToken(token string) (models.Token, error)
	GetUser(username string) (models.User, error)
	GetUserByToken(token string) (models.User, error)
	DeleteToken(token string) error
	GetUserByUserId(id int) (models.User, error)
}

type Post interface {
	CreatePosts(post models.Post) error
	GetPostByUserID(id int) ([]models.Post, error)
	GetPost() ([]models.Post, error)
	GetPostByTag(tags string) ([]models.Post, error)
	GetPostByPostID(postId int) (models.Post, error)
	UpdatePost(like string, dislike string, postId int) error
	LikedPosts(userId int) ([]models.Post, error)
	DislikedPosts(userId int) ([]models.Post, error)
	CommentedPosts(userId int) ([]models.Post, error)
	DeletePost(postId, userId int) error
	EditPost(postId int, title, text string) error
	CreateNotification(author, username, action string) error
	GetNotification(username string) ([]models.Notifications, error)
}

type Comment interface {
	AddComment(comment models.Comments) error
	GetCommentByPost(postId int) ([]models.Comments, error)
	UpdateComment(like string, dislike string, commentId int) error
	DeleteComment(commentId, userId int) error
	GetCommentById(commentID int) (models.Comments, error)
	EditComment(commentID int, comment string) error
}

type Evaluate interface {
	CreateEvaluates(postLike models.Evaluate) error
	CheckUserPost(userId, postId int) (models.Evaluate, error)
	UpdateVote(userId, postId, newVote int) error
	EvaluateCount(postId int) (models.Vote, error)
	CheckUserComment(userId, commentId int) (models.EvaluateComment, error)
	CreateEvaluateComment(commentVote models.EvaluateComment) error
	UpdateCommentVote(userId, commentId, newVote int) error
	EvaluateCommentCount(commentId int) (models.Vote, error)
	CheckVote(userId, postId, vote int) error
	CheckCommentVote(userId, postId, vote int) error
}

type Follow interface {
	CreateFollow(follow models.Follow) error
	DeleteFollow(follow models.Follow) error
	CheckFollow(follow models.Follow) error
	MyFollowers(userId int) ([]string, error)
	Following(userId int) ([]string, error)
}

type Service struct {
	Authorization
	Post
	Comment
	Evaluate
	Follow
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Post:          NewPostService(repo.Post),
		Comment:       NewCommentService(repo.Comment),
		Evaluate:      NewEvaluateService(repo.Evaluate),
		Follow:        NewFollowService(repo.Follow),
	}
}
