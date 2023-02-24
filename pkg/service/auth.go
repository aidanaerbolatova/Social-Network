package service

import (
	"Forum/models"
	"errors"

	"Forum/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user models.User) error {
	_, err := a.repo.CheckInvalid(user)
	if err == nil {
		return ErrCheckInvalid
	}
	if err := checkEmail(user.Email); err != nil {
		return err
	}
	if err := checkUsername(user.Username); err != nil {
		return err
	}
	if user.Method == "" {
		ok := checkPassword(user.Password)
		if !ok {
			return errors.New("password is not valid")
		}
		user.Password, err = generatePasswordHash(user.Password)
		if err != nil {
			return err
		}
	}

	return a.repo.CreateUser(user)
}

func generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return string(bytes), err
	}
	return string(bytes), nil
}
