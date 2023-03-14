package repository

import (
	"context"
	"database/sql"
	"time"

	"Forum/models"
)

type PostSQL struct {
	db *sql.DB
}

func NewPostSQL(db *sql.DB) *PostSQL {
	return &PostSQL{db: db}
}

func (r *PostSQL) CreatePosts(post models.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	query, err := r.db.PrepareContext(ctx, "INSERT INTO post (userId, title, text, category, createdAt, author, like_vote, dislike, image) values ($1, $2, $3, $4, $5, $6,$7,$8, $9)")
	if err != nil {
		return err
	}
	if _, err = query.Exec(post.UserId, post.Title, post.Text, post.Categories, post.CreatedAt, post.Author, "0", "0", post.Image); err != nil {
		return err
	}
	return nil
}

func (r *PostSQL) GetPostByUserID(id int) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var result []models.Post
	var post models.Post
	row, err := r.db.QueryContext(ctx, "SELECT * FROM post WHERE userID=$1 ", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Text, &post.Categories, &post.CreatedAt, &post.Author, &post.Like, &post.Dislike, &post.Image); err != nil {
			return result, err
		}
		evaluates, err := NewEvaluateSQL(r.db).EvaluateCount(post.Id)
		if err != nil {
			return nil, err
		}
		post.Like = evaluates.Like
		post.Dislike = evaluates.Dislike
		result = append(result, post)
	}
	return result, nil
}

func (r *PostSQL) GetPost() ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var result []models.Post
	var post models.Post
	row, err := r.db.QueryContext(ctx, "SELECT *FROM post")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Text, &post.Categories, &post.CreatedAt, &post.Author, &post.Like, &post.Dislike, &post.Image); err != nil {
			return nil, err
		}
		evaluates, err := NewEvaluateSQL(r.db).EvaluateCount(post.Id)
		if err != nil {
			return nil, err
		}
		post.Like = evaluates.Like
		post.Dislike = evaluates.Dislike
		result = append(result, post)
	}

	return result, nil
}

func (r *PostSQL) GetPostByTag(tags string) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var result []models.Post
	var post models.Post
	row, err := r.db.QueryContext(ctx, "SELECT * FROM post WHERE category like "+"'"+"%"+tags+"%'")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Text, &post.Categories, &post.CreatedAt, &post.Author, &post.Like, &post.Dislike, &post.Image); err != nil {
			return nil, err
		}
		evaluates, err := NewEvaluateSQL(r.db).EvaluateCount(post.Id)
		if err != nil {
			return nil, err
		}
		post.Like = evaluates.Like
		post.Dislike = evaluates.Dislike
		result = append(result, post)
	}
	return result, nil
}

func (r *PostSQL) GetPostByPostID(postId int) (models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var post models.Post
	row := r.db.QueryRowContext(ctx, "SELECT * FROM post WHERE id=$1", postId)
	if err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Text, &post.Categories, &post.CreatedAt, &post.Author, &post.Like, &post.Dislike, &post.Image); err != nil {
		return post, err
	}
	return post, nil
}

func (r *PostSQL) UpdatePost(like string, dislike string, postId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	_, err := r.db.ExecContext(ctx, "UPDATE post SET like_vote=$1, dislike=$2 WHERE id=$3 ", like, dislike, postId)
	if err != nil {
		return err
	}
	return nil
}

// ----------------delete post

func (r *PostSQL) DeletePost(postId, userId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if _, err := r.db.ExecContext(ctx, "DELETE FROM post WHERE id=$1 AND userID=$2", postId, userId); err != nil {
		return err
	}
	return nil
}

//-----------------edit post

func (r *PostSQL) EditPost(postId int, title, text string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	_, err := r.db.ExecContext(ctx, "UPDATE post SET title=$1, text=$2 WHERE id=$3", title, text, postId)
	if err != nil {
		return err
	}
	return nil
}

//----------------notification

func (r *PostSQL) CreateNotification(author, username, action string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	query, err := r.db.PrepareContext(ctx, "INSERT INTO notification(from_user, to_user, action_user) values ($1,$2,$3)")
	if err != nil {
		return err
	}
	if _, err := query.Exec(author, username, action); err != nil {
		return err
	}
	return nil
}

func (r *PostSQL) GetNotification(username string) ([]models.Notifications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var notification models.Notifications
	var notifications []models.Notifications
	row, err := r.db.QueryContext(ctx, "SELECT from_user, to_user, action_user FROM notification WHERE to_user=$1", username)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&notification.From, &notification.To, &notification.Action); err != nil {
			return nil, err
		}

		notifications = append(notifications, notification)
	}
	return notifications, nil
}
