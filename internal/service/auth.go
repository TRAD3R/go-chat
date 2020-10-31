package service

import (
	"github/trad3r/go_temp.git/internal/models"
	"github/trad3r/go_temp.git/internal/repository"
)

type AuthService struct {
	repository repository.Authorization
}

func NewAuthService(repository repository.Authorization) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s AuthService) CreateUser(user *models.User) error {
	user.HashPassword()
	return s.repository.CreateUser(user)
}
