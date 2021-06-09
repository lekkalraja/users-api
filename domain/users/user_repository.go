package users

import (
	"fmt"

	"github.com/lekkalraja/users-api/utils"
)

var (
	userDB = make(map[int64]*User)
)

func (u *User) Save() *utils.RestErr {
	_, ok := userDB[u.Id]

	if ok {
		return utils.NewBadRequest(fmt.Sprintf("User %d Already Exist", u.Id))
	}
	userDB[u.Id] = u
	return nil
}

func GetUsers() []*User {
	users := make([]*User, 0, len(userDB))

	for _, value := range userDB {
		users = append(users, value)
	}

	return users
}

func FindUser(id int64) (*User, *utils.RestErr) {
	user, ok := userDB[id]
	if !ok {
		return nil, utils.UserNotFound(id)
	}
	return user, nil
}
