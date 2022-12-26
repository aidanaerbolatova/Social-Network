package service

import (
	"Forum"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorNoUser        = errors.New("user not found")
	ErrorWrongPassword = errors.New("user password is not incorrect")
)

func (a *AuthService) GenerateToken(username, password string) (Forum.Token, error) {
	// get user from db
	user, err := a.repo.GetUser(username)
	if err != nil {
		return Forum.Token{}, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return Forum.Token{}, err
	}
	var token Forum.Token
	token = Forum.Token{
		UserId:    user.Id,
		AuthToken: uuid.NewString(),
		ExpiresAT: time.Now().Add(12 * time.Hour),
	}
	token2, err := a.repo.AddToken(token)
	if err != nil {
		return Forum.Token{}, err
	}

	return token2, nil
}

func (a *AuthService) GetToken(token string) (Forum.Token, error) {
	tokenStruct, err := a.repo.GetToken(token)
	if err != nil {
		return tokenStruct, err
	}
	return tokenStruct, nil
}

func (a *AuthService) GetUserByToken(token string) (Forum.User, error) {
	tokenStruct, err := a.repo.GetUserByToken(token)
	if err != nil {
		return Forum.User{}, err
	}
	return tokenStruct, nil
}

func (a *AuthService) DeleteToken(token string) error {
	return a.repo.DeleteToken(token)
}
