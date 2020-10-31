package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github/trad3r/go_temp.git/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r AuthPostgres) CreateUser(user *models.User) error {
	query := fmt.Sprintf("INSERT INTO %s(name, username, password_hash) VALUES($1, $2, $3) RETURNING id", tableUser)
	return r.db.QueryRow(query, user.Name, user.Username, user.Password).Scan(&user.Id)
}

func (r AuthPostgres) GetUser(user *models.User) error {
	query := fmt.Sprintf("SELECT u.* FROM  %s u WHERE u.username = $1 AND u.password_hash = $2", tableUser)
	return r.db.Get(user, query, user.Username, user.Password)
}
