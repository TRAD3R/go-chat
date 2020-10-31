package repository

import (
	"github.com/jmoiron/sqlx"
	"github/trad3r/go_temp.git/internal/models"
)

const (
	tableUser     = "users"
	tableList     = "lists"
	tableItem     = "items"
	tableUserList = "user_list"
	tableListItem = "list_item"
)

type Authorization interface {
	CreateUser(user *models.User) error
}

type List interface {
}

type Item interface {
}

type Repository struct {
	Authorization
	List
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
