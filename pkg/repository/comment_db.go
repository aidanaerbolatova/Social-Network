package repository

import (
	"Forum/models"
	"context"
	"database/sql"
	"time"
)

type CommentSQL struct {
	db *sql.DB
}

func NewCommentSQL(db *sql.DB) *CommentSQL {
	return &CommentSQL{db: db}
}

func (r *CommentSQL) AddComments(comments models.Comments) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	query, err := r.db.PrepareContext(ctx, "INSERT INTO comment (userId, postId, comment, createdAt, author, like_vote, dislike) values ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	_, err = query.Exec(comments.UserId, comments.PostId, comments.Comment, comments.CreatedAt, comments.Author, "0", "0")
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentSQL) GetCommentByPost(postId int) ([]models.Comments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var result []models.Comments
	var comment models.Comments
	row, err := r.db.QueryContext(ctx, "SELECT * FROM comment WHERE postId=$1", postId)
	if err != nil {
		return result, err
	}
	for row.Next() {
		if err := row.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Comment, &comment.CreatedAt, &comment.Author, &comment.Like, &comment.Dislike); err != nil {
			return result, err
		}
		evaluates, err := NewEvaluateSQL(r.db).EvaluateCommentCount(comment.Id)
		if err != nil {
			return nil, err
		}
		comment.Like = evaluates.Like
		comment.Dislike = evaluates.Dislike
		result = append(result, comment)
	}
	return result, nil
}

func (r *CommentSQL) UpdateComment(like string, dislike string, commentId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	_, err := r.db.ExecContext(ctx, "UPDATE comment SET like_vote=$1, dislike=$2 WHERE id=$3 ", like, dislike, commentId)
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentSQL) DeleteComment(commentId, userId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if _, err := r.db.ExecContext(ctx, "DELETE FROM comment WHERE id=$1 AND userId=$2", commentId, userId); err != nil {
		return err
	}
	return nil
}

func (r *CommentSQL) GetCommentById(commentID int) (models.Comments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var comment models.Comments
	row := r.db.QueryRowContext(ctx, "SELECT * FROM comment WHERE id=$1", commentID)
	if err := row.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Comment, &comment.CreatedAt, &comment.Author, &comment.Like, &comment.Dislike); err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *CommentSQL) EditComment(commentID int, comment string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if _, err := r.db.ExecContext(ctx, "UPDATE comment SET comment=$1 WHERE id=$2", comment, commentID); err != nil {
		return err
	}
	return nil
}
