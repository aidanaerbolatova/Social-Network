package repository

import (
	"Forum/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type AuthSQL struct {
	db *sql.DB
}

func NewAuthSQL(db *sql.DB) *AuthSQL {
	return &AuthSQL{db: db}
}

func (r *AuthSQL) CreateUser(user models.User) error {
	records := fmt.Sprintf("INSERT INTO users (email, username, password, auth_method) values ($1, $2, $3, $4)")
	query, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(user.Email, user.Username, user.Password, user.Method)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthSQL) GetUserByUsername(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var getUser models.User
	if err := r.db.QueryRowContext(ctx, "SELECT id, email, username, password FROM users WHERE username = $1 AND auth_method=$2 ", user.Username, user.Method).Scan(&getUser.Id, &getUser.Email, &getUser.Username, &getUser.Password); err != nil {
		return models.User{}, err
	}
	return getUser, nil
}

func (r *AuthSQL) GetUser(username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var user models.User
	if err := r.db.QueryRowContext(ctx, "SELECT id, email, username, password FROM users WHERE username = $1", username).Scan(&user.Id, &user.Email, &user.Username, &user.Password); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *AuthSQL) GetUserByEmail(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var getUser models.User
	if err := r.db.QueryRowContext(ctx, "SELECT id, email, username, password FROM users WHERE email = $1 AND auth_method=$2 ", user.Email, user.Method).Scan(&getUser.Id, &getUser.Email, &getUser.Username, &getUser.Password); err != nil {
		return models.User{}, err
	}
	return getUser, nil
}

func (r *AuthSQL) GetUserByUserId(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var user models.User
	if err := r.db.QueryRowContext(ctx, "SELECT email, username FROM users WHERE id=$1", id).Scan(&user.Email, &user.Username); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *AuthSQL) CheckInvalid(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var getUser models.User
	if err := r.db.QueryRowContext(ctx, "SELECT id, email, username FROM users WHERE email = $1 AND auth_method=$2 ", user.Email, user.Method).Scan(&getUser.Id, &getUser.Email, &getUser.Username); err != nil {
		return models.User{}, err
	}
	return getUser, nil
}
