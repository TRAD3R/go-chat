package models

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github/trad3r/go_temp.git/internal/helpers"
	"time"
)

type User struct {
	Id    int64
	Name  string `json:"name"`
	Image string `json:"image"`
}

func NewUser() *User {
	avatar, err := helpers.GetRandomAvatar()
	if err != nil {
		logrus.Info(err)
		avatar = "https://pickaface.net/gallery/avatar/unr_funny_170108_2338_7hs7qcl1.png"
	}
	return &User{
		Id:    time.Now().Unix(),
		Name:  helpers.GetRandomName(),
		Image: avatar,
	}
}

func (u User) String() string {
	return fmt.Sprintf("Id: %d, Name: %s", u.Id, u.Name)
}
