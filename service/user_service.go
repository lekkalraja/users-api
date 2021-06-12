package service

import (
	"github.com/lekkalraja/users-api/domain/users"
	"github.com/lekkalraja/users-api/utils"
)

func CreateUser(user users.User) (*users.User, *utils.RestErr) {
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id int64, user users.User) (int64, *utils.RestErr) {
	affectedRows, err := user.Update(id)
	if err != nil {
		return 0, err
	}
	return affectedRows, nil
}

func GetUsers() ([]*users.User, *utils.RestErr) {
	return users.GetUsers()
}

func FindUser(id int64) (*users.User, *utils.RestErr) {
	return users.FindUser(id)
}

func DeleteUser(id int64) (int64, *utils.RestErr) {
	return users.Delete(id)
}
