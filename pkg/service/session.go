package service

import (
	"Forum/models"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

var (
	ErrorNoUser        = errors.New("user not found")
	ErrorEmail         = errors.New("email is empty")
	ErrorWrongPassword = errors.New("user password is not incorrect")
	ErrCheckInvalid    = errors.New("user already exists")
)

func (a *AuthService) GenerateToken(users models.User, oauth bool) (models.Token, error) {
	var err error
	var user models.User
	if len(users.Email) == 0 {
		user, err = a.repo.GetUserByUsername(users)
		if err != nil {
			return models.Token{}, err
		}
	} else {
		user, err = a.repo.GetUserByEmail(users)
		if err != nil {
			return models.Token{}, err
		}
	}
	if !oauth {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(users.Password)); err != nil {
			return models.Token{}, err
		}
	}

	var token models.Token
	token = models.Token{
		UserId:    user.Id,
		AuthToken: uuid.NewString(),
		ExpiresAT: time.Now().Add(12 * time.Hour),
	}
	if err = a.repo.DeleteTokenByUserID(token.UserId); err != nil {
		return models.Token{}, err
	}
	token2, err := a.repo.AddToken(token)
	if err != nil {
		return models.Token{}, err
	}

	return token2, nil
}

func (a *AuthService) GetToken(token string) (models.Token, error) {
	tokenStruct, err := a.repo.GetToken(token)
	if err != nil {
		return tokenStruct, err
	}
	return tokenStruct, nil
}

func (a *AuthService) GetUserByToken(token string) (models.User, error) {
	userId, err := a.repo.GetUserIdByToken(token)
	if err != nil {
		return models.User{}, err
	}
	tokenStruct, err := a.repo.GetUserByToken(userId)
	if err != nil {
		return models.User{}, err
	}
	return tokenStruct, nil
}

func (a *AuthService) DeleteToken(token string) error {
	return a.repo.DeleteToken(token)
}
