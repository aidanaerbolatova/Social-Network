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

func (r *AuthSQL) GetUser(username, email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var user models.User
	request := fmt.Sprintf("SELECT * FROM users WHERE username = $1 or email=$2")
	rows := r.db.QueryRowContext(ctx, request, username, email)
	if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Method); err != nil {
		return user, err
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
