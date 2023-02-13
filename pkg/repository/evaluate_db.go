package repository

import (
	"Forum/models"
	"context"
	"database/sql"
	"time"
)

type EvaluateSQL struct {
	db *sql.DB
}

func NewEvaluateSQL(db *sql.DB) *EvaluateSQL {
	return &EvaluateSQL{db: db}
}

func (r *EvaluateSQL) CreateEvaluates(postLike models.Evaluate) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	row, err := r.db.PrepareContext(ctx, "INSERT INTO evaluate (userId, postId, vote) values ($1, $2, $3)")
	if err != nil {
		return err
	}
	_, err = row.Exec(postLike.UserId, postLike.PostId, postLike.Vote)
	if err != nil {
		return err
	}
	return nil
}

func (r *EvaluateSQL) CreateEvaluateComment(commentVote models.EvaluateComment) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	row, err := r.db.PrepareContext(ctx, "INSERT INTO evaluateComment (userId, commentId, vote) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	_, err = row.Exec(commentVote.UserId, commentVote.CommentId, commentVote.Vote)
	if err != nil {
		return err
	}
	return nil
}

func (r *EvaluateSQL) CheckUserPost(userId, postId int) (models.Evaluate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var evaluate models.Evaluate
	row, err := r.db.QueryContext(ctx, "SELECT userId, postId, vote FROM evaluate WHERE userId=$1 and postId=$2", userId, postId)
	if err != nil {
		return models.Evaluate{}, err
	}
	for row.Next() {
		if err := row.Scan(&evaluate.UserId, &evaluate.PostId, &evaluate.Vote); err != nil {
			return evaluate, err
		}
	}
	return evaluate, nil
}

func (r *EvaluateSQL) CheckUserComment(userId, commentId int) (models.EvaluateComment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var evaluate models.EvaluateComment
	row, err := r.db.QueryContext(ctx, "SELECT userId, commentId, vote FROM evaluateComment WHERE userId=$1 and commentId=$2", userId, commentId)
	if err != nil {
		return evaluate, err
	}
	for row.Next() {
		if err := row.Scan(&evaluate.UserId, &evaluate.CommentId, &evaluate.Vote); err != nil {
			return evaluate, err
		}
	}
	return evaluate, nil
}

func (r *EvaluateSQL) UpdateVote(userId, postId, newVote int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	_, err := r.db.ExecContext(ctx, "UPDATE evaluate SET vote=$1 WHERE userId=$2 AND postId=$3", newVote, userId, postId)
	if err != nil {
		return err
	}
	return nil
}

func (r *EvaluateSQL) UpdateCommentVote(userId, commentId, newVote int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	_, err := r.db.ExecContext(ctx, "UPDATE evaluateComment SET vote=$1 WHERE userId=$2 AND commentId=$3", newVote, userId, commentId)
	if err != nil {
		return err
	}
	return nil
}

func (r *EvaluateSQL) EvaluateCount(postId int) (models.Vote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var like, dislike string
	query, err := r.db.PrepareContext(ctx, "SELECT COUNT(vote) FROM evaluate WHERE postId=$2 and vote=$3")
	if err != nil {
		return models.Vote{}, err
	}
	defer query.Close()
	err = query.QueryRowContext(ctx, postId, 1).Scan(&like)
	if err != nil {
		return models.Vote{}, err
	}
	err = query.QueryRowContext(ctx, postId, -1).Scan(&dislike)
	if err != nil {
		return models.Vote{}, err
	}
	result := models.Vote{
		Like:    like,
		Dislike: dislike,
	}
	return result, nil
}

func (r *EvaluateSQL) EvaluateCommentCount(commentId int) (models.Vote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var like, dislike string
	query, err := r.db.PrepareContext(ctx, "SELECT COUNT(vote) FROM evaluateComment WHERE commentId=$2 and vote=$3")
	if err != nil {
		return models.Vote{}, err
	}
	defer query.Close()
	err = query.QueryRowContext(ctx, commentId, 1).Scan(&like)
	if err != nil {
		return models.Vote{}, err
	}
	err = query.QueryRowContext(ctx, commentId, -1).Scan(&dislike)
	if err != nil {
		return models.Vote{}, err
	}
	result := models.Vote{
		Like:    like,
		Dislike: dislike,
	}
	return result, nil
}

func (r *EvaluateSQL) CheckVote(userId, postId, vote int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if _, err := r.db.ExecContext(ctx, "UPDATE evaluate SET vote=$1 WHERE userId=$2 AND postId=$3 AND vote=$4", 0, userId, postId, vote); err != nil {
		return err
	}
	return nil
}

func (r *EvaluateSQL) CheckCommentVote(userId, postId, vote int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if _, err := r.db.ExecContext(ctx, "UPDATE evaluateComment SET vote=$1 WHERE userId=$2 AND commentId=$3 AND vote=$4", 0, userId, postId, vote); err != nil {
		return err
	}
	return nil
}
