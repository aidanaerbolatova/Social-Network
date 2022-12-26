package repository

import (
	"Forum"
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

func (r *AuthSQL) CreateUser(user Forum.User) error {
	records := fmt.Sprint("INSERT INTO users (Email, Username, Firstname, Lastname, Password) values ($1, $2, $3, $4, $5) ")
	query, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(user.Email, user.Username, user.FirstName, user.LastName, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthSQL) GetUser(username string) (Forum.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var user Forum.User
	rows := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE Username = $1 ", username)
	if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}

func (r *AuthSQL) CheckInvalid(username, email string) (Forum.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var user Forum.User
	if err := r.db.QueryRowContext(ctx, "SELECT id, email, username FROM users WHERE Email = $1 AND Username=$2 ", email, username).Scan(&user.Id, &user.Email, &user.Username); err != nil {
		return user, err
	}
	return user, nil
}
