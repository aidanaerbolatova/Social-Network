package repository

import (
	"Forum/models"
	"context"
	"time"
)

func (r *AuthSQL) AddToken(token models.Token) (models.Token, error) {
	query, err := r.db.Prepare("INSERT INTO authorization_token (userId, auth_token, expires_at) values ($1, $2, $3)")
	if err != nil {
		return models.Token{}, err
	}
	_, err = query.Exec(token.UserId, token.AuthToken, token.ExpiresAT)
	if err != nil {
		return models.Token{}, err
	}
	return token, nil
}

func (r *AuthSQL) GetToken(token string) (models.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30*time.Second))
	defer cancel()
	var userToken models.Token
	rows := r.db.QueryRowContext(ctx, "SELECT id, userId, auth_token, expires_at FROM authorization_token WHERE auth_token=$1", token)
	if err := rows.Scan(&userToken.Id, &userToken.UserId, &userToken.AuthToken, &userToken.ExpiresAT); err != nil {
		return userToken, err
	}
	return userToken, nil
}

func (r *AuthSQL) GetUserIdByToken(token string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var userId int
	if err := r.db.QueryRowContext(ctx, "SELECT userId FROM authorization_token WHERE auth_token=$1", token).Scan(&userId); err != nil {
		return userId, err
	}
	return userId, nil
}

func (r *AuthSQL) GetUserByToken(userId int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	var user models.User
	if err := r.db.QueryRowContext(ctx, "SELECT id, email, username, password, auth_method FROM users WHERE id=$1", userId).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Method); err != nil {
		return user, err
	}
	return user, nil
}

func (r *AuthSQL) DeleteToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15*time.Second))
	defer cancel()
	_, err := r.db.ExecContext(ctx, "DELETE FROM authorization_token WHERE auth_token=$1", token)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthSQL) DeleteTokenByUserID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15*time.Second))
	defer cancel()
	_, err := r.db.ExecContext(ctx, "DELETE FROM authorization_token WHERE userId=$1", id)
	if err != nil {
		return err
	}
	return nil
}
