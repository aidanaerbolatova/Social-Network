package repository

import (
	"Forum/models"
	"context"
	"database/sql"
	"time"
)

type FollowSQL struct {
	db *sql.DB
}

func NewFollowSQL(db *sql.DB) *FollowSQL {
	return &FollowSQL{db: db}
}

func (r *FollowSQL) CreateFollow(follow models.Follow) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	row, err := r.db.PrepareContext(ctx, "INSERT INTO follow (userId, followerId) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	_, err = row.Exec(follow.UserId, follow.AuthorId)
	if err != nil {
		return err
	}
	return nil
}

func (r *FollowSQL) DeleteFollow(follow models.Follow) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if _, err := r.db.ExecContext(ctx, "DELETE FROM follow WHERE userId=$1 AND followerId=$2", follow.UserId, follow.AuthorId); err != nil {
		return err
	}
	return nil
}

func (r *FollowSQL) CheckFollow(follow models.Follow) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var id int
	if err := r.db.QueryRowContext(ctx, "SELECT id FROM follow WHERE userId=$1 AND followerId=$2", follow.UserId, follow.AuthorId).Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *FollowSQL) MyFollowers(userId int) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var myfollowers string
	var followers []string
	row, err := r.db.QueryContext(ctx, "SELECT userId FROM follow WHERE followerId=$1", userId)
	if err != nil {
		return []string{}, err
	}
	for row.Next() {
		if err := row.Scan(&myfollowers); err != nil {
			return []string{}, err
		}
		followers = append(followers, myfollowers)
	}
	return followers, nil
}

func (r *FollowSQL) Following(userId int) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var following string
	var followers []string
	row, err := r.db.QueryContext(ctx, "SELECT followerId FROM follow WHERE userId=$1", userId)
	if err != nil {
		return []string{}, err
	}
	for row.Next() {
		if err := row.Scan(&following); err != nil {
			return []string{}, err
		}
		followers = append(followers, following)
	}
	return followers, nil
}
