package service

import (
	"github/trad3r/go_temp.git/internal/models"
	"github/trad3r/go_temp.git/internal/repository"
)

type Authorization interface {
	CreateUser(user *models.User) error
	GenerateToken(user *models.User) (string, error)
}

type List interface {
}

type Item interface {
}

type Service struct {
	Authorization
	List
	Item
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
	}
}
