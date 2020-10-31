package models

import (
	"crypto/sha512"
	"fmt"
)

const salt = "KDFJnkseNLKEKgj9*9JKE#R"

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}

func (u *User) HashPassword() {
	hash := sha512.New()
	hash.Write([]byte(u.Password))
	u.Password = fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
