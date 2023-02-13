package repository

import (
	"Forum/models"
	"context"
	"time"
)

func (r *PostSQL) LikedPosts(userId int) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var evaluate models.Evaluate
	var post models.Post
	var posts []int
	var result []models.Post
	row, err := r.db.QueryContext(ctx, "SELECT postId FROM evaluate WHERE userId=$1 AND vote=$2 ", userId, "1")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&evaluate.PostId); err != nil {
			return nil, err
		}
		posts = append(posts, evaluate.PostId)
	}
	for _, id := range posts {
		if err := r.db.QueryRowContext(ctx, "SELECT * FROM post WHERE id=$1", id).
			Scan(&post.Id, &post.UserId, &post.Title, &post.Text, &post.Categories, &post.CreatedAt, &post.Author, &post.Like, &post.Dislike, &post.Image); err != nil {
			return nil, err
		}
		result = append(result, post)
	}
	return result, nil
}

func (r *PostSQL) DislikedPosts(userId int) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var evaluate models.Evaluate
	var post models.Post
	var posts []int
	var result []models.Post
	row, err := r.db.QueryContext(ctx, "SELECT postId FROM evaluate WHERE userId=$1 AND vote=$2 ", userId, "-1")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&evaluate.PostId); err != nil {
			return nil, err
		}
		posts = append(posts, evaluate.PostId)
	}

	for _, id := range posts {
		if err := r.db.QueryRowContext(ctx, "SELECT * FROM post WHERE id=$1", id).
			Scan(&post.Id, &post.UserId, &post.Title, &post.Text, &post.Categories, &post.CreatedAt, &post.Author, &post.Like, &post.Dislike, &post.Image); err != nil {
			return nil, err
		}
		result = append(result, post)
	}
	return result, nil
}

func (r *PostSQL) CommentedPosts(userId int) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var comment models.Comments
	var posts []int
	var post models.Post
	var result []models.Post
	row, err := r.db.QueryContext(ctx, "SELECT postId FROM comment WHERE userId=$1", userId)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&comment.PostId); err != nil {
			return nil, err
		}
		if !contains(posts, comment.PostId) {
			posts = append(posts, comment.PostId)
		}
	}
	for _, id := range posts {
		if err := r.db.QueryRowContext(ctx, "SELECT * FROM post WHERE id=$1", id).
			Scan(&post.Id, &post.UserId, &post.Title, &post.Text, &post.Categories, &post.CreatedAt, &post.Author, &post.Like, &post.Dislike, &post.Image); err != nil {
			return nil, err
		}
		result = append(result, post)
	}
	return result, nil
}

func contains(a []int, num int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == num {
			return true
		}
	}
	return false
}
