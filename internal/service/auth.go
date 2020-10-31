package service

import (
	"github.com/dgrijalva/jwt-go"
	"github/trad3r/go_temp.git/internal/models"
	"github/trad3r/go_temp.git/internal/repository"
	"time"
)

const (
	tokenTTL = time.Hour * 12
	salt     = "OIJNJNmij#(m(H#9397mJIih"
	userCtx  = "user_id"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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

func (s AuthService) GenerateToken(user *models.User) (string, error) {
	user.HashPassword()
	if err := s.repository.GetUser(user); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(salt))
}
