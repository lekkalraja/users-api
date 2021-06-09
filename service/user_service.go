package service

import (
	"github.com/lekkalraja/users-api/domain"
	"github.com/lekkalraja/users-api/utils"
)

var users []domain.User

func CreateUser(user domain.User) (*domain.User, *utils.RestErr) {
	users = append(users, user)
	return &user, nil
}

func GetUsers() []domain.User {
	return users
}

func FindUser(id int64) (*domain.User, *utils.RestErr) {
	for _, user := range users {
		if id == user.Id {
			return &user, nil
		}
	}
	return nil, utils.UserNotFound(id)
}
